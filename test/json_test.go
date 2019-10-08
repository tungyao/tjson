package test

import (
	"../../tjson"
	"log"
	"testing"
)

func TestJSON(t *testing.T) {
	data := "{\"name\":\"123h_____ehiu12123aduhausd$asdjiqwejsd\",\"age\":26}"
	out := tjson.Decode(data)
	log.Println(out)
	out1 := map[string]interface{}{"name": "son", "age": 12, "other": map[string]interface{}{"phone": "xiaomi", "test": 14}}
	log.Println(tjson.Encode(out1))
}
