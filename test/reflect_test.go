package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/JFJun/go-substrate-rpc-client/v3/scale"
	"github.com/JFJun/go-substrate-rpc-client/v3/types"
	"reflect"
	"testing"
)

type T struct {
	U8  types.U8
	U32 types.U32
}

func Test_ReflectType(t *testing.T) {
	tt := T{}
	rt := reflect.TypeOf(tt)
	fmt.Println(rt.NumField())
	fmt.Println(rt.Field(0).Type)

	rn := reflect.New(rt.Field(1).Type)
	data := []byte{1, 43, 2, 12, 3, 22, 2, 2, 34}
	decoder := scale.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(rn.Interface())
	if err != nil {
		t.Fatal(err)
	}

	s := make(map[string]interface{})
	s["U8"] = rn.Interface()
	d, _ := json.Marshal(s)
	fmt.Println(string(d))
}
