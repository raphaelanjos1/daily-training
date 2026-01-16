package main

func main() {
	list := []int{1, 2, 3, 4}
	item := 4
	binarySearch(list, item)
}

func binarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		middle := (low + high) / 2
		guess := list[middle]

		if guess == item {
			return middle
		} else if guess > item {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	return -1
}
