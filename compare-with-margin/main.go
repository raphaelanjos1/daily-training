package main

func main() {
	closeCompare(8.1, 5.0, 3.0)
}

func closeCompare(a, b, margin float64) int {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}

	if diff <= margin {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}
