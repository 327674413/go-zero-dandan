package arrd

func InArray[T comparable](findTarget T, fromArr []T) bool {
	for _, v := range fromArr {
		if v == findTarget {
			return true
		}
	}
	return false
}

func InArrayWithIndex[T comparable](findTarget T, fromArr []T) int {
	for k, v := range fromArr {
		if v == findTarget {
			return k
		}
	}
	return -1
}

func Reverse[T comparable](targetArrPt *[]T) {
	arr := *targetArrPt
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - 1 - i
		arr[i], arr[j] = arr[j], arr[i]
	}
}
