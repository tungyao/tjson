package test

import (
	"../../tjson"
	"testing"
)

func TestJSON(t *testing.T) {
	data := `{"name":"tung:yao","age":"12"}`
	//out := tjson.Decode(data)
	//t.Log(out)
	tjson.Decode([]byte(data))
	//out1 := map[string]interface{}{"name": "son", "age": 12, "other": map[string]interface{}{"phone": "xiaomi", "test": 14}}
	//log.Println(tjson.Encode(out1))
}
