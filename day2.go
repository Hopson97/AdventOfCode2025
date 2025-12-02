package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
)

func day_two() {
	content, err := os.ReadFile("input/day2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	re_p1 := regexp2.MustCompile(`^(.+)\1$`, 0)
	re_p2 := regexp2.MustCompile(`^(.+)\1+$`, 0)

	product_id_ranges := strings.Split(string(content), ",")
	var total_part1 = 0
	var total_part2 = 0

	for _, product_id_range := range product_id_ranges {
		product_ids := strings.Split(string(product_id_range), "-")

		begin, _ := strconv.Atoi(product_ids[0])
		end, _ := strconv.Atoi(product_ids[1])

		for i := begin; i <= end; i++ {
			match, _ := re_p1.MatchString(strconv.Itoa(i))
			if match {
				total_part1 += i
			}

			match2, _ := re_p2.MatchString(strconv.Itoa(i))
			if match2 {
				total_part2 += i
			}
		}
	}
	fmt.Printf("Part1: %d - Part2: %d ", total_part1, total_part2)
}
