package test

import (
	"log"
	"testing"
)
import "../../tjson"

func TestJSON(t *testing.T) {
	data := "{\"name\":\"123h_____ehiu12123aduhausd$asdjiqwejsd\",\"age\":26}"
	T := new(tnjson.JSON)
	out := T.Decode(data)
	log.Println(out)
}
