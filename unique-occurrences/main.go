package main

func main() {
	arr := []int{1, 2, 2, 1, 1, 3}
	uniqueOccurrences(arr)
}

func uniqueOccurrences(arr []int) bool {
	freq := make(map[int]int)

	for _, item := range arr {
		freq[item]++
	}

	seen := make(map[int]bool)

	for _, count := range freq {
		if seen[count] {
			return false
		}
		seen[count] = true
	}

	return true
}
