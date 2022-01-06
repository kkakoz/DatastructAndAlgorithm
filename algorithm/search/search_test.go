package search

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

var data []int

func TestBinarySearchRecursion(t *testing.T) {
	index := rand.Intn(len(data))
	val := data[index]
	searchIndex := BinarySearchRecursion(data, val)
	if searchIndex != index {
		t.Error("search err")
	}
}

func TestBinarySearch(t *testing.T) {
	index := rand.Intn(len(data))
	val := data[index]
	searchIndex := BinarySearch(data, val)
	if searchIndex != index {
		t.Error("search err")
	}
}

func InitData(length int) {
	rand.Seed(time.Now().Unix())
	data = make([]int, 0, length)
	addVal := 1
	for i := 0; i < length; i++ {
		addVal += rand.Intn(10) + 1
		data = append(data, addVal)
	}
	sort.Ints(data)
}

func TestMain(m *testing.M) {
	InitData(100)
	m.Run()
}
