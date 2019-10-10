package test

import (
	"../../tjson"
	"testing"
)

func TestJSON(t *testing.T) {
	data := "{\"id\":\"1\",\"session_key\":\"24SlFzzCtR8qNbqGJzwgdA==\"}"
	out := tjson.Decode(data)
	t.Log(out["id"])
	//out1 := map[string]interface{}{"name": "son", "age": 12, "other": map[string]interface{}{"phone": "xiaomi", "test": 14}}
	//log.Println(tjson.Encode(out1))
}
