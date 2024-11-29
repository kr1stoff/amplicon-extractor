package common

func RemoveDupplicates(slice []int) []int {
	uniqueMap := make(map[int]bool)
	uniqueSlice := []int{}
	for _, num := range slice {
		if !uniqueMap[num] {
			uniqueMap[num] = true
			uniqueSlice = append(uniqueSlice, num)
		}
	}
	return uniqueSlice
}
