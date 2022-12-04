package main

import (
	_ "embed"
	"fmt"
	"github.com/juliangruber/go-intersect"
	"reflect"
	"regexp"
	strconv "strconv"
	"strings"
)

//go:embed "input.txt"
var input string

func main() {
	input_text := ""
	input_text = strings.TrimSuffix(input, "\n")
	lines := regexp.MustCompile("[\\,\\n]+").Split(input_text, -1)

	//Part1(lines)
	Part2(lines)
}

func SumCrossings(t interface{}) int {
	sum := 0
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)

		for i := 0; i < s.Len(); i++ {
			var arr reflect.Value = s.Index(i)
			sum += arr.Interface().(int)
		}
	}
	return sum
}

func Part1(lines []string) {
	leftRange := []int{}
	rightRange := []int{}
	count := 0

	for i := 0; i < len(lines); i += 2 {
		firstRange := strings.Split(lines[i], "-")
		Fnums1, _ := strconv.Atoi(firstRange[0])
		Fnums2, _ := strconv.Atoi(firstRange[1])
		Fsum := 0
		for i := Fnums1; i <= Fnums2; i++ {
			leftRange = append(leftRange, i)
			Fsum += i
		}

		secondRange := strings.Split(lines[i+1], "-")
		Snums1, _ := strconv.Atoi(secondRange[0])
		Snums2, _ := strconv.Atoi(secondRange[1])
		Ssum := 0
		for i := Snums1; i <= Snums2; i++ {
			rightRange = append(rightRange, i)
			Ssum += i
		}

		//fmt.Println(firstRange[0], firstRange[1], "=", Fsum, ", ", secondRange[0], secondRange[1], "= ", Ssum)
		sumA := 0
		for _, i := range leftRange {
			sumA += i
		}
		sumB := 0
		for _, i := range rightRange {
			sumB += i
		}
		if SumCrossings(intersect.Simple(leftRange, rightRange)) == sumA || SumCrossings(intersect.Simple(leftRange, rightRange)) == sumB {
			count++
		}
		leftRange, rightRange = nil, nil
	}
	fmt.Println(count)
}

func Part2(lines []string) {
	leftRange := []int{}
	rightRange := []int{}
	count := 0

	for i := 0; i < len(lines); i += 2 {
		firstRange := strings.Split(lines[i], "-")
		Fnums1, _ := strconv.Atoi(firstRange[0])
		Fnums2, _ := strconv.Atoi(firstRange[1])
		Fsum := 0
		for i := Fnums1; i <= Fnums2; i++ {
			leftRange = append(leftRange, i)
			Fsum += i
		}

		secondRange := strings.Split(lines[i+1], "-")
		Snums1, _ := strconv.Atoi(secondRange[0])
		Snums2, _ := strconv.Atoi(secondRange[1])
		Ssum := 0
		for i := Snums1; i <= Snums2; i++ {
			rightRange = append(rightRange, i)
			Ssum += i
		}

		if SumCrossings(intersect.Simple(leftRange, rightRange)) > 0 {
			count++
		}
		leftRange, rightRange = nil, nil
	}
	fmt.Println(count)
}
