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

	Part1(lines)
}

//type Stack struct {
//	box [][]string
//}

func Part1(lines []string) {
	var Stack [8][9]string
	x := 0

	for i, line := range lines {
		y := 0
		if line == "" {
			break
		}
		fmt.Println(i, ": ", line)
		for j := 1; j < len(line); j += 4 {
			if line[j] > 64 && line[j] < 91 {
				//fmt.Print(string(line[j]), " ")
				Stack[x][y] = string(line[j])
				y++
			} else if line[j] == 32 {
				Stack[x][y] = "♦"
				y++
			}
		}
		x++
	}
	fmt.Println(Stack)
}
