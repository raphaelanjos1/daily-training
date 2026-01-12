package main

func main() {
	jewels := "aA"
	stones := "aAAbbbb"
	numJewelsInStones(jewels, stones)
}

func numJewelsInStones(jewels string, stones string) int {
	jewelsMap := make(map[rune]bool)

	for _, jewel := range jewels {
		jewelsMap[jewel] = true
	}

	sum := 0

	for _, stone := range stones {
		if jewelsMap[stone] {
			sum++
		}
	}

	println(sum)
	return sum
}
