package util

func Int64InArray(i int64, arr []int64) bool {
	for _, item := range arr {
		if i == item {
			return true
		}
	}
	return false
}
