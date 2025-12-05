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
	fresh_ranges := []FreshRange{}
	ingredients := []int{}
	for _, line := range lines {
		if strings.Contains(line, "-") {
			fresh_range := strings.Split(line, "-")
			low, _ := strconv.Atoi(fresh_range[0])
			high, _ := strconv.Atoi(fresh_range[1])
			fresh_ranges = append(fresh_ranges, FreshRange{low, high})
		} else {
			if len(line) != 0 {
				ingredient, _ := strconv.Atoi(line)
				ingredients = append(ingredients, ingredient)
			}

		}
	}
	fresh_list := mapset.NewSet[int]()
	for _, ingredient := range ingredients {
		for _, fresh_range := range fresh_ranges {
			if ingredient >= fresh_range.low && ingredient <= fresh_range.high {
				fresh_list.Add(ingredient)
				//fmt.Printf("Ingredient %d is fresh (%d-%d)\n", ingredient, fresh_range.low, fresh_range.high)
			}
		}
	}

	fmt.Printf("Day 5 Part 1: %d\n", fresh_list.Cardinality())
}
