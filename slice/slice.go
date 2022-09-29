package slice

// 判断元素是否在切片中
func InSlice[T comparable](item T, slice []T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
