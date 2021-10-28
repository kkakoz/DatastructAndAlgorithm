package sorts

import (
	"math/rand"
	"testing"
	"time"
)

var sortData []int

var sortSizes = []int{1000, 10000, 50000}

type SortFunc func(arr []int)

func initSortData(dataLen int) {
	rand.Seed(time.Now().Unix())
	dataRange := 10000
	sortData = make([]int, 0, dataLen)
	for i := 0; i < dataLen; i++ {
		sortData = append(sortData, rand.Intn(dataRange))
	}
}

func TestMain(m *testing.M) {
	for _, size := range sortSizes {
		initSortData(size)
		m.Run()
	}
}

func SortTest(f SortFunc, sortName string, t *testing.T)  {
	size := len(sortData)
	tempData := make([]int, size)
	copy(tempData, sortData)
	before := time.Now()
	f(tempData)
	useTime := time.Now().Sub(before)
	validate := validateSortData(tempData)
	if !validate {
		t.Error("sort failed")
	}
	t.Logf("sort = %s size = %d, use time = %v", sortName, size, useTime)
}

func validateSortData(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i - 1] {
			return false
		}
	}
	return true
}

func TestSelectSort(t *testing.T) {
	SortTest(SelectSort, "SelectSort", t)
}

func TestInsertSort(t *testing.T) {
	SortTest(InsertSort, "InsertSort", t)
}

func TestInsertSort2(t *testing.T) {
	SortTest(InsertSort2, "InsertSort2", t)
}

func TestMergeSort(t *testing.T) {
	SortTest(MergeSort, "MergeSort", t)
}

func TestMergeSort2(t *testing.T) {
	SortTest(MergeSort2, "MergeSort", t)
}

func TestQuickSort(t *testing.T) {
	SortTest(QuickSort, "QuickSort", t)
}

func TestBubbleSort(t *testing.T) {
	SortTest(BubbleSort, "BubbleSort", t)
}