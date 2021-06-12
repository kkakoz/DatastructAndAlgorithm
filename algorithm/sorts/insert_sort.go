package sorts

func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j - 1] {
				arr[j], arr[j - 1] = arr[j - 1], arr[j]
			}
		}
	}
}

func InsertSort2(arr []int)  {
	for i := 0; i < len(arr); i++ {
		ele := arr[i]
		var j int
		for j = i; j > 0 && ele < arr[j - 1]; j-- {
			arr[j] = arr[j -1]
		}
		arr[j] = ele
	}
}