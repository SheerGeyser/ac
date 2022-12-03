package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
)

//A-65 ROCK
//B-66 PAPER
//C-67 Scissors
//Me----------
//X-88 ROCK
//Y-89 PAPER
//Z-90 Scissors

const win = 6
const draw = 3
const loose = 0

const rock = 1
const paper = 2
const scissors = 3

func main() {

	fmt.Println(first())
	fmt.Println(second())

}

func first() int {
	counter := 0
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()
		b := []byte(text)
		enemyHand := int(big.NewInt(0).SetBytes(b[:1]).Uint64())
		myHand := int(big.NewInt(0).SetBytes(b[len(b)-1:]).Uint64())

		//--------------WITH ROCK
		if enemyHand == 65 && myHand == 88 {
			counter += draw + rock
		} else if enemyHand == 65 && myHand == 89 {
			counter += win + paper
		} else if enemyHand == 65 && myHand == 90 {
			counter += loose + scissors
		}

		//--------------WITH PAPER
		if enemyHand == 66 && myHand == 88 {
			counter += loose + rock
		} else if enemyHand == 66 && myHand == 89 {
			counter += draw + paper
		} else if enemyHand == 66 && myHand == 90 {
			counter += win + scissors
		}

		//--------------WITH scissors
		if enemyHand == 67 && myHand == 88 {
			counter += win + rock
		} else if enemyHand == 67 && myHand == 89 {
			counter += loose + paper
		} else if enemyHand == 67 && myHand == 90 {
			counter += draw + scissors
		}

		//fmt.Println(enemyHand, myHand)
	}

	return counter
}

func second() int {
	counter := 0
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()
		b := []byte(text)
		enemyHand := int(big.NewInt(0).SetBytes(b[:1]).Uint64())
		myHand := int(big.NewInt(0).SetBytes(b[len(b)-1:]).Uint64())

		//--------------WITH ROCK
		if enemyHand == 65 && myHand == 88 {
			counter += loose + scissors
		} else if enemyHand == 65 && myHand == 89 {
			counter += draw + rock
		} else if enemyHand == 65 && myHand == 90 {
			counter += win + paper
		}

		//--------------WITH PAPER
		if enemyHand == 66 && myHand == 88 {
			counter += loose + rock
		} else if enemyHand == 66 && myHand == 89 {
			counter += draw + paper
		} else if enemyHand == 66 && myHand == 90 {
			counter += win + scissors
		}

		//--------------WITH scissors
		if enemyHand == 67 && myHand == 88 {
			counter += loose + paper
		} else if enemyHand == 67 && myHand == 89 {
			counter += draw + scissors
		} else if enemyHand == 67 && myHand == 90 {
			counter += win + rock
		}

		//fmt.Println(enemyHand, myHand)
	}

	return counter
}
