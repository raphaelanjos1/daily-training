package main

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	intersection(nums1, nums2)
}

func intersection(nums1 []int, nums2 []int) []int {
	hashmap := make(map[int]bool)

	for _, num1 := range nums1 {
		hashmap[num1] = false
	}

	for _, num2 := range nums2 {
		if _, exists := hashmap[num2]; exists {
			hashmap[num2] = true
		}
	}

	values := []int{}

	for index, hash := range hashmap {
		if hash == true {
			values = append(values, index)
		}
	}

	return values
}
