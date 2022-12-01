package main

import (
	"fmt"
	"os"

	"github.com/alenius/aoc22/day1"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please select the day you want to run as the first arg")
		os.Exit(1)
	}
	day := os.Args[1]

	switch day {
	case "1":
		day1.Day1()
	default:
		fmt.Println("day not implemented yet")
	}
}
