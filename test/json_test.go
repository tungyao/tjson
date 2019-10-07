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
	log.Println(new(tjson.JSON).Encode(out))
}
