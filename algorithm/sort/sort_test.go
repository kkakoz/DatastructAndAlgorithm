package sort

import (
	"math/rand"
	"testing"
	"time"
)

var sortData []int

func initSortData(dataLen int) {
	rand.Seed(time.Now().Unix())
	dataRange := 10000
	sortData = make([]int, 0, dataLen)
	for i := 0; i < dataLen; i++ {
		sortData = append(sortData, rand.Intn(dataRange))
	}
}

func TestMain(m *testing.M) {
	initSortData(1000)
	m.Run()
	initSortData(10000)
	m.Run()
	initSortData(100000)
	m.Run()
}

func TestSelectSort(t *testing.T) {
	size := len(sortData)
	tempData := make([]int, size)
	copy(tempData, sortData)
	before := time.Now()
	SelectSort(tempData)
	useTime := time.Now().Sub(before)
	validate := validateSortData(tempData)
	if !validate {
		t.Error("sort failed")
	}
	t.Logf("size = %d, use time = %v", size, useTime)
}

func TestInsertSort(t *testing.T) {
	size := len(sortData)
	tempData := make([]int, size)
	copy(tempData, sortData)
	before := time.Now()
	InsertSort(tempData)
	useTime := time.Now().Sub(before)
	validate := validateSortData(tempData)
	if !validate {
		t.Error("sort failed")
	}
	t.Logf("size = %d, use time = %v", size, useTime)
}

func TestInsertSort2(t *testing.T) {
	size := len(sortData)
	tempData := make([]int, size)
	copy(tempData, sortData)
	before := time.Now()
	InsertSort2(tempData)
	useTime := time.Now().Sub(before)
	validate := validateSortData(tempData)
	if !validate {
		t.Error("sort failed")
	}
	t.Logf("size = %d, use time = %v", size, useTime)
}

func validateSortData(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i - 1] {
			return false
		}
	}
	return true
}
