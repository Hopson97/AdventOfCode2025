package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max_jolts(bank string) (int, int) {
	max_index := -1
	max_jolts := 0
	for i, ch := range bank {
		jolts := int(ch - '0')
		if jolts > max_jolts {
			max_jolts = jolts
			max_index = i
		}
	}
	return max_jolts, max_index
}

func day_three() {
	content, err := os.ReadFile("input/day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	banks := strings.Split(string(content), "\n")

	total := 0
	for _, bank := range banks {
		first, i := max_jolts(bank[0 : len(bank)-1])
		second, _ := max_jolts(bank[i+1:])

		n, _ := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(second))

		// fmt.Printf("Bank: %s -> %d %d %d\n", bank, first, second, n)
		total += n

	}
	fmt.Printf("Total: %d", total) // 17281
}
