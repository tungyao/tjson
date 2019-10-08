package test

import (
	"../../tjson"
	"log"
	"testing"
)

func TestJSON(t *testing.T) {
	data := "{\"name\":\"023ijCDb27x3JM0AqMDb25tJDb2ijCD3\",\"age\":26}"
	out := tjson.Decode(data)
	log.Println(out["name"].(string) + "asd")
	//out1 := map[string]interface{}{"name": "son", "age": 12, "other": map[string]interface{}{"phone": "xiaomi", "test": 14}}
	//log.Println(tjson.Encode(out1))
}
