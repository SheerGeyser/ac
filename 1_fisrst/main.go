package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	arr := []int{0, 0, 0}
	var elfSum int

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if num > 0 {
			elfSum += num
		} else {
			arr = append(arr, elfSum)
			elfSum = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Ints(arr)
	sums := 0
	for _, i := range arr {
		fmt.Println(i)
		if i > 68704 {
			sums += i
		}
	}
	fmt.Println("TOP 3:   ", arr[len(arr)-3:])
	fmt.Println("SUM:   ", sums)
	fmt.Println(Max(arr))
}

func Max(array []int) int {
	var max int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}
