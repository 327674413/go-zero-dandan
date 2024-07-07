package arrd

// Contain 数组中是否包含某个元素
func Contain[T comparable](fromArr []T, findTarget T) bool {

	for _, v := range fromArr {
		if v == findTarget {
			return true
		}
	}
	return false
}

// Index 返回数组中的位置，不存在则是-1
func Index[T comparable](fromArr []T, findTarget T) int {
	for k, v := range fromArr {
		if v == findTarget {
			return k
		}
	}
	return -1
}

// Reverse 将数组内容反转
func Reverse[T comparable](targetArrPt *[]T) {
	arr := *targetArrPt
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - 1 - i
		arr[i], arr[j] = arr[j], arr[i]
	}
}
