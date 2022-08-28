func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)
	for index, value := range nums {
		if idx, isPresent := sumMap[target-value]; isPresent {
			fmt.Println(idx)
			return []int{idx, index}
		}
		sumMap[value] = index
	}
	return []int{}
}