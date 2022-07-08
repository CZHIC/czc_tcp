package pubFuc

import (
	"czc_tcp/app/model"
	"czc_tcp/library/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/net/gudp"
	"github.com/tealeg/xlsx"
)

//json  转 map
func JSONToMap(str string) (map[string]interface{}, error) {

	var tempMap map[string]interface{}

	err := json.Unmarshal([]byte(str), &tempMap)

	if err != nil {
		return nil, err
	}
	return tempMap, nil
}

//json  转 Arr_map
func JSONToArrMap(str string) ([]map[string]interface{}, error) {

	var tempMap []map[string]interface{}

	err := json.Unmarshal([]byte(str), &tempMap)

	if err != nil {
		return nil, err
	}
	return tempMap, nil
}

//合并两个map
func MergeMap(a map[string]interface{}, b map[string]interface{}) map[string]interface{} {
	for k, v := range a {
		b[k] = v
	}
	return b
}

/**
* 结构体赋值
* binding type interface 要修改的结构体
* value type interace 有数据的结构体
 */
func StructAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}

// addslashes() 函数返回在预定义字符之前添加反斜杠的字符串。
// 预定义字符是：
// 单引号（'）
// 双引号（"）
// 反斜杠（\）
func Addslashes(str string) string {
	tmpRune := []rune{}
	strRune := []rune(str)
	for _, ch := range strRune {
		switch ch {
		case []rune{'\\'}[0], []rune{'"'}[0], []rune{'\''}[0]:
			tmpRune = append(tmpRune, []rune{'\\'}[0])
			tmpRune = append(tmpRune, ch)
		default:
			tmpRune = append(tmpRune, ch)
		}
	}
	return string(tmpRune)
}

// stripslashes() 函数删除由 addslashes() 函数添加的反斜杠。
func Stripslashes(str string) string {
	dstRune := []rune{}
	strRune := []rune(str)
	strLenth := len(strRune)
	for i := 0; i < strLenth; i++ {
		if strRune[i] == []rune{'\\'}[0] {
			i++
		}
		dstRune = append(dstRune, strRune[i])
	}
	return string(dstRune)
}

//excel解析
func ExcelParse(fileName string) ([]string, [][]string) {
	dir, _ := os.Getwd() // 获取main.go所在的路径
	filePath := dir + "/document/" + fileName
	xlFile, err := xlsx.OpenFile(filePath)

	if err != nil {
		logger.Error("ExcelParse", err)
	}
	//获取行数
	//length := len(xlFile.Sheets[0].Rows)
	//开辟除表头外的行数的数组内存
	resourceArr := [][]string{}
	title := []string{}
	//遍历sheet
	for _, sheet := range xlFile.Sheets {
		//遍历每一行
		for rowIndex, row := range sheet.Rows {
			//跳过第一行表头信息
			if rowIndex == 0 {
				for _, cell := range row.Cells {
					text := cell.String()
					title = append(title, text)
				}
				continue
			}
			//遍历每一个单元
			result := []string{}
			for cellIndex, cell := range row.Cells {
				text := cell.String()
				if text != "" {
					//如果是每一行的第一个单元格
					if cellIndex == 0 {
						logger.Print("第一列", cellIndex)
					}
				}
				result = append(result, text)

			}
			resourceArr = append(resourceArr, result)
			//logger.Print("每一行的结果：", result)
		}
	}
	return title, resourceArr
}

//截取小数精度
func FloatRound(f float64, n int) float64 {

	format := "%." + strconv.Itoa(n) + "f"

	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)

	return res

}

func GetExchangeRate(text string) []map[string]interface{} {

	//text := "100日元等于多少美元"
	time1 := gconv.String(time.Now().Unix()-200) + "000"
	time2 := gconv.String(time.Now().Unix()) + "000"
	logger.Print(time1, time2)
	url := "http://sp0.baidu.com/8aQDcjqpAAV3otqbppnN2DJv/api.php?query=" + text + "&co=&resource_id=6017&t=" + time2 + "&cardId=6017&ie=utf8&oe=utf8&cb=op_aladdin_callback&format=json&tn=baidu&cb=jQuery1102012038258240309319_" + time1 + "&_=" + time2
	cc := g.Client()
	postData := ""
	s, err := cc.Get(url, postData)
	if err != nil {
		logger.Error(err)
	}
	ret := s.ReadAllString()
	arr := strings.Split(ret, "(")
	strs := arr[1][0 : len(arr[1])-2]
	jsonMap, _ := JSONToMap(strs)
	jsonMap2, _ := JSONToArrMap(gconv.String(jsonMap["data"]))
	//endtext := jsonMap2[0]["content1Mini"]
	return jsonMap2
}

func GET(url string) {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil) //建立一个请求
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}
	//Add 头协议
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Add("Accept-Language", "ja,zh-CN;q=0.8,zh;q=0.6")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("Cookie", "设置cookie")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	response, err := client.Do(reqest) //提交
	if err != nil {
		logger.Print(err)
	}
	defer response.Body.Close()
	cookies := response.Cookies() //遍历cookies
	for _, cookie := range cookies {
		fmt.Println("cookie:", cookie)
	}

	body, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		logger.Print(err)
	}
	fmt.Println(string(body)) //网页源码

}

//data => "name=1&age=2"
func Post(url string, data string) {
	client := &http.Client{}
	reqest, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		logger.Print(err)
	}
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Add("Accept-Language", "ja,zh-CN;q=0.8,zh;q=0.6")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("Cookie", "设置cookie")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	resp, err := client.Do(reqest)
	cookies := resp.Cookies()
	for _, cookie := range cookies {
		fmt.Println("cookie:", cookie)
	}
	defer resp.Body.Close()

	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		logger.Print(err)
	}

	fmt.Println(string(body))
}

// 自定义格式发送消息包
func SendPkg(conn *gtcp.Conn, act string, data ...string) error {
	s := ""
	if len(data) > 0 {
		s = data[0]
	}
	msg, err := json.Marshal(model.Msg{
		Act:  act,
		Data: s,
	})
	if err != nil {
		panic(err)
	}
	return conn.SendPkg(msg)
}

// 自定义格式接收消息包
func RecvPkg(conn *gtcp.Conn) (msg *model.Msg, err error) {
	if data, err := conn.RecvPkg(); err != nil {
		return nil, err
	} else {
		msg = &model.Msg{}
		err = json.Unmarshal(data, msg)
		if err != nil {
			return nil, fmt.Errorf("invalid package structure: %s", err.Error())
		}
		return msg, err
	}
}

func UdpSendPkg(conn *gudp.Conn, msg *model.Msg) error {
	msgstr, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return conn.Send(msgstr)
}

// 自定义格式接收消息包
func UdpRecvPkg(conn *gudp.Conn) (msg *model.Msg, err error) {
	if data, err := conn.Recv(-1); err != nil {
		return nil, err
	} else {
		msg = &model.Msg{}
		logger.Println("收到请求---二进制----：", data)
		fmt.Println("--------收到请求---二进制------", data)
		logger.Println("收到请求---二进制----：", string(data))
		err = json.Unmarshal(data, msg)
		if err != nil {
			return nil, nil
			//return nil, fmt.Errorf("invalid package structure: %s", err.Error())
		}
		return msg, err
	}
}
