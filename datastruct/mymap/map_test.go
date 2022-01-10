package mymap

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type User struct {
	Name string
}

func TestHash(t *testing.T) {
	a := 1
	//a1, err := GobEncode(a)
	//if err != nil {
	//	t.Fatal(err)
	//}
	b := 1
	//b1, err := GobEncode(b)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Println(a1)
	//fmt.Println(b1)
	//err := binary.Write(buf, binary.BigEndian, a)
	//if err != nil {
	//	t.Fatal(err)
	//}

	u := memhash(unsafe.Pointer(&a), 1, 36)
	fmt.Println(u)
	bu := memhash(unsafe.Pointer(&b), 1, 36)
	fmt.Println(bu)
}

func GobEncode(v any) ([]byte, error) {
	if reflect.TypeOf(v).Kind() == reflect.Pointer {
		v = fmt.Sprintf("%p", v)
		fmt.Println(v)
	}
	s1 := make([]byte, 0)
	buf := bytes.NewBuffer(s1)
	err := gob.NewEncoder(buf).Encode(v)
	return buf.Bytes(), err
}
