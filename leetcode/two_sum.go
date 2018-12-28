package leetcode

func twoSum(nums []int, target int) []int {
	ix := map[int]int{}
	for i, value := range nums {
		candidate := target - value
		if j, exists := ix[candidate]; exists {
			return []int{j, i}
		}
		ix[value] = i
	}

	panic("no answer")
}
