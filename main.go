package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	day_five()
	elapsed := time.Since(start)
	fmt.Println()
	fmt.Println(elapsed)
}
