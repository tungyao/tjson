package tjson

import (
	"encoding/json"
	"github.com/tungyao/yell"
	"reflect"
	"strconv"
	"strings"
)

var log = yell.New(yell.Config{
	Path:     "/var/log",
	FileName: "tjson",
}, "[TJSON]")

type JSON struct {
	json strings.Builder
	mp   bool
	jp   map[string]interface{}
}

func (j *JSON) _format(obj interface{}) *JSON {
	switch reflect.ValueOf(obj).Kind() {
	case reflect.Ptr:
		f := reflect.ValueOf(obj).Elem().Interface().(map[string]interface{})
		for k, v := range f {
			switch reflect.TypeOf(v).Kind() {
			case reflect.Map:
				j.json.WriteString("\"" + k + "\":{")
				j.mp = true
				j._format(v)
			case reflect.String:
				j.json.WriteString("\"" + k + "\":" + "\"" + reflect.ValueOf(v).String() + "\",")
			case reflect.Int:
				j.json.WriteString("\"" + k + "\":" + "\"" + strconv.Itoa(reflect.ValueOf(v).Interface().(int)) + "\",")
			}
		}
	case reflect.Map:
		f := reflect.ValueOf(obj).Interface().(map[string]interface{})
		for k, v := range f {
			switch reflect.TypeOf(v).Kind() {
			case reflect.Ptr:
				j.json.WriteString("\"" + k + "\":{")
				j.mp = false
				j._format(v)
			case reflect.String:
				j.json.WriteString("\"" + k + "\":" + "\"" + reflect.ValueOf(v).String() + "\"")
			case reflect.Int:
				j.json.WriteString("\"" + k + "\":" + "\"" + strconv.Itoa(reflect.ValueOf(v).Interface().(int)) + "\"")
			}
			if j.mp == true {
				j.json.WriteString("}")
			}
			j.json.WriteString(",")
		}
	}
	return j
}
func Encode(obj interface{}) string {
	//js := j._format(obj).json.String()
	//js = js[:len(js)-1]
	//return "{" + js + "}"
	data, err := json.Marshal(obj)
	if err != nil {
		println(err)
	}
	return string(data)
}

//func Decode(json string) map[string]interface{} {
//	jp := make(map[string]interface{})
//	json = strings.ReplaceAll(json, "\n", "")
//	json = strings.ReplaceAll(json, "\t", "")
//	json = strings.ReplaceAll(json, "\r", "")
//	json = json[1 : len(json)-1]
//	stringArr := strings.Split(json, ",")
//	for i := 0; i < len(stringArr); i++ {
//		str := strings.Split(stringArr[i], ":")
//		jp[strings.Replace(str[0], "\"", "", -1)] = strings.Replace(str[1], "\"", "", -1)
//	}
//	return jp
//}

type braces struct {
	index int
	value byte
	tp    int
}
type jsons struct {
	braces []braces `从左往右计算出现多少个大括号并记录位置 123 125`
	comma  []int    `记录逗号位置 44`
}

func Decode(buff []byte) map[string]interface{} {
	js := new(jsons)
	outMap := make(map[string]interface{})
	for k, v := range buff {
		if v == 123 {
			js.braces = append(js.braces, braces{index: k, value: v, tp: -1})
		} else if v == 125 {
			js.braces = append(js.braces, braces{index: k, value: v, tp: 1})
		} else if v == 44 {
			js.comma = append(js.comma, k)
		}
	}
	js.comma = append(js.comma, len(buff))
	if len(js.braces)%2 != 0 && len(js.braces) > 1 {
		log.Panicln("JSON format error ", string(buff))
	}
	buf := make([][]byte, 0)
	for i := 0; i < len(js.comma); i++ {
		in := make([]byte, 5)
		if i == 0 {
			in = buff[0:js.comma[i]]
		} else {
			in = buff[1*js.comma[i-1]+1 : js.comma[i]]
		}
		buf = append(buf, in)
	}
	for i := 0; i < len(buf); i++ {
		colon := make([]int, 0)
		for j := 0; j < len(buf[i]); j++ {
			if buf[i][j] == 58 {
				colon = append(colon, j)
			}
		}
		if formatByteToMap(buf[i], colon) {
			outMap[string(deleteSymbol(buf[i][:colon[0]]))] = string(deleteSymbol(buf[i][colon[0]+1:]))
		} else {
			outMap[string(deleteSymbol(buf[i][:colon[0]]))] = Decode(buf[i][colon[0]+1:])
		}

	}
	log.Println("---TJSON---", outMap)
	return outMap
}
func formatByteToMap(j []byte, n []int) bool {
	for i := 0; i < len(j); i++ {
		if i > n[0] && i < n[len(n)-1:][0] && j[i] == 123 {
			return false
		}
	}
	return true
}
func deleteSymbol(b []byte) []byte {
	for j := 0; j < len(b); j++ {
		if b[j] == 123 {
			b = append(b[:j], b[j+1:]...)
		}
	}
	for j := 0; j < len(b); j++ {
		if b[j] == 34 {
			b = append(b[:j], b[j+1:]...)
		}
	}
	for j := 0; j < len(b); j++ {
		if b[j] == 125 {
			b = append(b[:j], b[j+1:]...)
		}
	}
	return b
}
