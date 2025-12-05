package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

type FreshRange struct {
	low  int
	high int
}

func day_five() {
	content, err := os.ReadFile("input/day5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(content), "\n")
	fresh_ranges := mapset.NewSet[FreshRange]()
	ingredients := []int{}
	for _, line := range lines {
		if strings.Contains(line, "-") {
			fresh_range := strings.Split(line, "-")
			low, _ := strconv.Atoi(fresh_range[0])
			high, _ := strconv.Atoi(fresh_range[1])
			fresh_ranges.Add(FreshRange{low, high})
		} else {
			if len(line) != 0 {
				ingredient, _ := strconv.Atoi(line)
				ingredients = append(ingredients, ingredient)
			}

		}
	}
	fresh_list := mapset.NewSet[int]()
	for _, ingredient := range ingredients {
		for fresh_range := range fresh_ranges.Iterator().C {
			if ingredient >= fresh_range.low && ingredient <= fresh_range.high {
				fresh_list.Add(ingredient)
			}
		}
	}

	fmt.Printf("Day 5 Part 1: %d\n", fresh_list.Cardinality())

	found := true
	iterations := 0

	// Reduce the list of ranges down until there is nothing left to reduce
	for found {
		iterations++
		actual_fresh_ranges := mapset.NewSet[FreshRange]()
		found = false
		for fresh_range_a := range fresh_ranges.Iterator().C {
			real_low := fresh_range_a.low
			real_high := fresh_range_a.high

			for fresh_range_b := range fresh_ranges.Iterator().C {

				if fresh_range_a == fresh_range_b {
					continue
				}

				if real_low >= fresh_range_b.low && real_low <= fresh_range_b.high {
					real_low = fresh_range_b.low
					found = true
				}
				if real_high <= fresh_range_b.high && real_high >= fresh_range_b.low {
					real_high = fresh_range_b.high
					found = true
				}
			}
			actual_fresh_ranges.Add(FreshRange{real_low, real_high})
		}
		fresh_ranges = actual_fresh_ranges
	}

	total := 0
	for fresh_range := range fresh_ranges.Iterator().C {
		total += fresh_range.high - fresh_range.low + 1
	}
	fmt.Printf("Day 5 Part 2: %d (%d iterations) \n", total, iterations)

}
