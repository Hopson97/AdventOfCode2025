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

	pattern := `^(.+)\1$`
	re := regexp2.MustCompile(pattern, 0)

	text := "123123"
	match, _ := re.MatchString(text)

	fmt.Println(match)

	product_id_ranges := strings.Split(string(content), ",")
	var total = 0

	for index, product_id_range := range product_id_ranges {
		product_ids := strings.Split(string(product_id_range), "-")

		begin, _ := strconv.Atoi(product_ids[0])
		end, _ := strconv.Atoi(product_ids[1])

		for i := begin; i <= end; i++ {
			match, _ := re.MatchString(strconv.Itoa(i))
			if match {
				fmt.Printf("%d -> %d\n", index, i)
				total += i
			}
		}
	}
	fmt.Printf("Total: %d ", total)
}
