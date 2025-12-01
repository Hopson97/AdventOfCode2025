package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day_one() {
	// instructions, err := os.ReadFile("input/day1.txt")
	instructions, err := os.ReadFile("input/day1_example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(instructions), "\n")
	var dial = 50
	var zero_count = 0

	for index, instruction := range lines {
		instruction = strings.TrimSpace(instruction)
		dir := string(instruction[0])
		steps, _ := strconv.Atoi(instruction[1:])

		switch dir {
		case "L":
			dial -= steps
		case "R":
			dial += steps
		}

		dial = ((dial % 100) + 100) % 100

		if dial == 0 {
			zero_count++
		}
		fmt.Printf("Step: %d - Instruction -> (%s %d) - Result: %d \n\n", index, dir, steps, dial)
	}
	fmt.Printf("Password: %d ", zero_count)

}
