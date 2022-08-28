func maxSubArray(nums []int) int {
	result := nums[0]
	sum := 0
	for _, value := range nums {
		tmp := sum + value
		result = max(result, tmp)
		if tmp > 0 {
			sum = tmp
		} else {
			sum = 0
		}
	}
	return result
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}