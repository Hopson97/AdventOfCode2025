package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	day_three()
	elapsed := time.Since(start)
	fmt.Println()
	fmt.Println(elapsed)
}
