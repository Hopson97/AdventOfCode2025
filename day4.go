package main

import (
	"fmt"
	"os"
	"strings"
)

type xy struct {
	x int
	y int
}

func accessible(grid []string) []xy {
	width := len(grid[0])
	height := len(grid)
	coords := []xy{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '@' {
				total := 0
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
					coords = append(coords, xy{x, y})
				}
			}

		}
	}
	return coords
}

func remove_accessible(grid []string, total int) int {
	positions := accessible(grid)

	if len(positions) != 0 {
		for _, position := range positions {
			row := []byte(grid[position.y])
			row[position.x] = '.'
			grid[position.y] = string(row)
		}
		total = remove_accessible(grid, total+len(positions))
	}
	return total
}

func day_four() {
	content, err := os.ReadFile("input/day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	grid := strings.Split(string(content), "\n")
	count := len(accessible(grid))

	total_removed := remove_accessible(grid, 0)

	fmt.Printf("Day 4 part 1 %d\n", count)       // 1560
	fmt.Printf("Day 4 part 2 %d", total_removed) // 1560
}
