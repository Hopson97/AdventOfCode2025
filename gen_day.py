"""Generates AoC files for a given day."""

import os


DAY = 4
DAY_STRING = "four"

TEMPLATE = f"""
package main

import (
	"fmt"
	"os"
)

func day_{DAY_STRING}() {{
	content, err := os.ReadFile("input/day{DAY}_example.txt")
	if err != nil {{
		fmt.Println(err)
		return
	}}

	fmt.Printf("Day {DAY} input: %s", content)
}}

"""
GO_FILE = f"day{DAY}.go"
if not os.path.exists(GO_FILE):
    with open(f"input/day{DAY}_example.txt", "w", encoding="UTF-8"):
        pass

    with open(f"input/day{DAY}.txt", "w", encoding="UTF-8"):
        pass

    with open(GO_FILE, "w", encoding="UTF-8") as f:
        f.write(TEMPLATE)
else:
    print(f"Did not create, files for {DAY} already exist.")
