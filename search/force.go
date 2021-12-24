package search

func Contains(array []int, target int) bool {
	for i := 0; i < len(array); i++ {
		if i == target {
			return true
		}
	}
	return false
}
