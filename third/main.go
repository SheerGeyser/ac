package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed "input.txt"
var input string

func main() {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	fmt.Println("First ans: ", Part1(lines))
	fmt.Println("Second ans: ", Part2(lines))
}

func Part1(lines []string) int {
	res := 0
	sum := 0
	for _, line := range lines {
		res = FindRepeatOneLine(line[:len(line)/2], line[len(line)/2:])
		sum += res
	}
	return sum
}

func FindRepeatOneLine(first string, second string) int {
	for _, i := range first {
		for _, j := range second {
			if i == j {
				if i > 96 {
					return int(i - 96)
				} else {
					return int(i - 38)
				}
			}
		}
	}
	return 0
}

func Part2(lines []string) int {
	res := 0
	sum := 0
	for i := 0; i < len(lines); i += 3 {
		res = FindRepeatThreeLine(lines[i], lines[i+1], lines[i+2])
		sum += res
	}
	return sum
}

func FindRepeatThreeLine(a string, b string, c string) int {
	for _, i := range a {
		for _, j := range b {
			for _, k := range c {
				if i == j && j == k {
					if i > 96 {
						return int(i - 96)
					} else {
						return int(i - 38)
					}
				}
			}
		}
	}
	return 0
}
