package main

func main() {
	nums := []int{1, 2, 3, 1, 1, 3}
	numIdenticalPairs(nums)
}

func numIdenticalPairs(nums []int) int {
	/*
		good_pairs := 0

		for index, num := range nums {
			for i := 0; i < len(nums); i++ {
				if num == nums[i] && index < i {
					good_pairs += 1
				}
			}
		}

		return good_pairs
		Time O(n2)
		Space O(1)
	*/
	freq := make(map[int]int)
	pairs := 0

	for _, num := range nums {
		pairs += freq[num]
		freq[num]++
	}
	return pairs
	// Time O(n)
	// Space O(k)
}
