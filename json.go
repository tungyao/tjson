package tjson

import (
	"encoding/json"
	"log"
	"reflect"
	"strconv"
	"strings"
)

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
func Decode(json string) map[string]interface{} {
	jp := make(map[string]interface{})
	json = strings.ReplaceAll(json, "\n", "")
	json = strings.ReplaceAll(json, "\t", "")
	json = strings.ReplaceAll(json, "\r", "")
	json = json[1 : len(json)-1]
	stringArr := strings.Split(json, ",")
	for i := 0; i < len(stringArr); i++ {
		str := strings.Split(stringArr[i], ":")
		log.Println(reflect.TypeOf(str[1]))
		jp[strings.Replace(str[0], "\"", "", -1)] = strings.Replace(str[1], "\"", "", -1)
	}
	return jp
}
