package list_test

import (
	"DatastructAndAlgorithm/datastruct/list"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestList(t *testing.T) {
	Convey("list", t, func() {
		Convey("list add", func() {
			linkList := list.NewLinkList()
			testData := []int{1, 5, 2, 3, 1}
			for _, data := range testData {
				linkList.Push(data)
			}
			for i, data := range testData {
				v, err := linkList.Get(i)
				So(err, ShouldBeNil)
				So(v, ShouldEqual, data)
			}
			So(len(testData), ShouldEqual, linkList.Len())
		})
		Convey("list ")
	})
}

func UnNil(value interface{}, f func()) {
	switch v := value.(type) {
	case string:
		if v == "" {
			return
		}
	case int, uint, int64, int32, int16, int8, uint64, uint32, uint16, uint8:
		if v == 0 {
			return
		}
	}
}
