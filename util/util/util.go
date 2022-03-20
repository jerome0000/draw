package util

// ArrayContain .
func ArrayContain[T int64](i T, arr []T) bool {
	for _, item := range arr {
		if i == item {
			return true
		}
	}
	return false
}
