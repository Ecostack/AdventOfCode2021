package day4

import (
	"AdventOfCode2021/util"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func parseNumbers(line string, separator string) []int {
	numStrings := strings.Split(line, separator)
	result := make([]int, 0)
	for _, num := range numStrings {
		if len(num) > 0 {
			val, err := strconv.ParseInt(num, 10, 32)
			if err != nil {
				panic(err)
			}
			result = append(result, int(val))
		}
	}
	return result
}

type Bingo [][]int

func hasBingoWon(bingo Bingo, picked map[int]struct{}) bool {
	for _, row := range bingo {
		matches := 0
		for _, column := range row {
			if _, ok := picked[column]; ok {
				matches++
			}
		}
		if matches == len(row) {
			return true
		}
	}
	for column := 0; column < len(bingo); column++ {
		matches := 0
		for row := 0; row < len(bingo); row++ {
			if _, ok := picked[bingo[row][column]]; ok {
				matches++
			}
		}
		if matches == len(bingo) {
			return true
		}
	}
	return false
}

func calcBingoSum(bingo Bingo, picked map[int]struct{}) int {
	sum := 0
	for _, row := range bingo {
		for _, column := range row {
			if _, ok := picked[column]; !ok {
				sum = sum + column
				//matches++
			}
		}
		//if matches == len(row) {
		//	return true
		//}
	}
	return sum
}

func TestDay4Part1(t *testing.T) {
	currentWD, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	lines, err := util.GetFileContentsSplit(currentWD + "/day4.txt")
	if err != nil {
		panic(err)
	}
	pickedNumbers := parseNumbers(lines[0], ",")
	log.Println(pickedNumbers)

	bingos := make([]Bingo, 0)
	var bingo Bingo
	//bingoMatrixLineIndex := 0

	for lineIndex, line := range lines {
		if lineIndex == 0 {
			continue
		}
		if len(line) == 0 {
			if bingo != nil {
				bingos = append(bingos, bingo)
			}

			bingo = make(Bingo, 0)
			continue
		}
		numbers := parseNumbers(line, " ")
		bingo = append(bingo, numbers)
		//log.Println(numbers)
	}
	log.Println(bingos)
	currentPickedNumbers := make(map[int]struct{})
	for _, currentPicked := range pickedNumbers {
		currentPickedNumbers[currentPicked] = struct{}{}
		//found :=
		for _, bingo := range bingos {
			val := hasBingoWon(bingo, currentPickedNumbers)
			log.Println(val)
			if val {
				sum := calcBingoSum(bingo, currentPickedNumbers)
				log.Println(sum * currentPicked)
				break
			}
		}
	}
}

func TestDay4Part2(t *testing.T) {
	currentWD, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	lines, err := util.GetFileContentsSplit(currentWD + "/day4.txt")
	if err != nil {
		panic(err)
	}
	pickedNumbers := parseNumbers(lines[0], ",")
	log.Println(pickedNumbers)

	bingos := make([]Bingo, 0)
	var bingo Bingo
	//bingoMatrixLineIndex := 0

	for lineIndex, line := range lines {
		if lineIndex == 0 {
			continue
		}
		if len(line) == 0 {
			if bingo != nil {
				bingos = append(bingos, bingo)
			}

			bingo = make(Bingo, 0)
			continue
		}
		numbers := parseNumbers(line, " ")
		bingo = append(bingo, numbers)
		//log.Println(numbers)
	}
	log.Println(bingos)
	currentPickedNumbers := make(map[int]struct{})
	bingosNotYetWon := make([]Bingo, len(bingos))
	copy(bingosNotYetWon, bingos)
	lastPicked := 0
	for _, currentPicked := range pickedNumbers {
		currentPickedNumbers[currentPicked] = struct{}{}
		bingosTemp := make([]Bingo, 0)
		for _, bingo := range bingosNotYetWon {

			val := hasBingoWon(bingo, currentPickedNumbers)
			if val {
				lastPicked = currentPicked
				if len(bingosNotYetWon) == 1 {
					sum := calcBingoSum(bingo, currentPickedNumbers)
					log.Println(sum * lastPicked)
				}
			} else {
				bingosTemp = append(bingosTemp, bingo)
			}
		}
		bingosNotYetWon = bingosTemp
	}
}
