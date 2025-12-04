package main

import (
	"fmt"
	"os"
	"strings"
)

func day_four() {
	content, err := os.ReadFile("input/day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	grid := strings.Split(string(content), "\n")
	width := len(grid[0])
	height := len(grid)

	count := 0
	papers := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '@' {
				total := 0
				papers++
				for oy := -1; oy <= 1; oy++ {
					for ox := -1; ox <= 1; ox++ {
						if oy == 0 && ox == 0 {
							continue
						}
						ax := x + ox
						ay := y + oy
						if ay >= 0 && ay < height && ax >= 0 && ax < width &&
							grid[ay][ax] == '@' {
							total++
						}
					}
				}

				if total < 4 {
					count++
				}
			}

		}
	}

	fmt.Printf("Day 4 part 1 %d", count) // 1560
}
