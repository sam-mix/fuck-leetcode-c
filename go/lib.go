package lib

func max(arr ...int) int {
	m := arr[0]
	for i := 1; i < len(arr); i++ {
		if m < arr[i] {
			m = arr[i]
		}
	}
	return m
}

func min(arr ...int) int {
	m := arr[0]
	for i := 1; i < len(arr); i++ {
		if m > arr[i] {
			m = arr[i]
		}
	}
	return m
}
