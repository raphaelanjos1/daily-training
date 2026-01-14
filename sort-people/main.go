package main

func main() {
	names := []string{"Mary", "John", "Emma"}
	heights := []int{180, 165, 170}
	sortPeople(names, heights)
}

func sortPeople(names []string, heights []int) []string {
	byHeight := make(map[int]string, len(heights))
	for i := 0; i < len(heights); i++ {
		byHeight[heights[i]] = names[i]
	}

	for i := 1; i < len(heights); i++ {
		key := heights[i]
		j := i - 1
		for j >= 0 && heights[j] < key {
			heights[j+1] = heights[j]
			j--
		}
		heights[j+1] = key
	}

	res := make([]string, len(heights))
	for i := 0; i < len(heights); i++ {
		res[i] = byHeight[heights[i]]
	}

	return res
}
