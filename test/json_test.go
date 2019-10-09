package test

import (
	"../../tjson"
	"log"
	"testing"
)

func TestJSON(t *testing.T) {
	data := `{"id":"1","sess                              ion_key":"24SlFzzCtR8qNbqGJzwgdA=="}`
	out := tjson.Decode(data)
	log.Println(out)
	log.Println(new(tjson.JSON).Encode(out))
}
