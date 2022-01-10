package list_test

import (
	"DatastructAndAlgorithm/datastruct/list"
	"constraints"
	"fmt"
	"testing"
)

func print[T constraints.Ordered](l list.List[T]) {
	fmt.Println(l)
}

func TestList(t *testing.T) {
	var linkList = list.NewLinkList[int]()
	linkList.Push(1, 2, 3, 4, 5)
	print[int](linkList)
	//Convey("list", t, func() {
	//	Convey("list add", func() {
	//
	//	})
	//	Convey("list ")
	//})
}
