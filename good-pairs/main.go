package main

func main() {
	nums := []int{1, 2, 3, 1, 1, 3}
	numIdenticalPairs(nums)
}

func numIdenticalPairs(nums []int) int {
	good_pairs := 0

	for index, num := range nums {
		for i := 0; i < len(nums); i++ {
			if num == nums[i] && index < i {
				good_pairs += 1
			}
		}
	}

	return good_pairs
}
