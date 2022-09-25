package day3

import (
	"AdventOfCode2021/util"
	"errors"
	"fmt"
	"log"
	"strconv"
)

func Day3Part1() error {
	lines, err := util.GetFileContentsSplit("day3.txt")
	if err != nil {
		fmt.Println("Err", err)
		return err
	}

	positionCounter := make([]int, len(lines[0]))

	for _, line := range lines {
		for index, lineChar := range line {
			switch {
			case lineChar == '0':
				positionCounter[index]--
			case lineChar == '1':
				positionCounter[index]++
			default:
				return errors.New("line contains neither 0 nor 1")
			}
		}
	}

	gamma := ""
	epsilon := ""

	for _, value := range positionCounter {
		if value > 0 {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}
	}

	gammaNum, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		return err
	}
	epsilonNum, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		return err
	}
	fmt.Printf("gamma %v, epsilon %v, gammaNum %v, epsilon %v, gamma*epsilon %v\n", gamma, epsilon, gammaNum, epsilonNum, gammaNum*epsilonNum)

	return nil
}

func binaryToDecimal(val string) (int64, error) {
	return strconv.ParseInt(val, 2, 64)
}

func findPart2(lines []string, moreFirst bool) string {
	columns := len(lines[0])

	colLines := make([]string, len(lines))
	copy(colLines, lines)

	for col := 0; col < columns; col++ {
		lineZeros := []string{}
		lineOnes := []string{}
		for _, line := range colLines {
			if line[col] == '0' {
				lineZeros = append(lineZeros, line)
			} else {
				lineOnes = append(lineOnes, line)
			}
		}

		if len(lineZeros) == len(lineOnes) {
			if moreFirst {
				colLines = lineOnes
			} else {
				colLines = lineZeros
			}
			break
		}

		if moreFirst {
			if len(lineZeros) > len(lineOnes) {
				colLines = lineZeros
			} else {
				colLines = lineOnes
			}
		} else {
			if len(lineZeros) > len(lineOnes) {
				colLines = lineOnes
			} else {
				colLines = lineZeros
			}
		}
	}
	if len(colLines) != 1 {
		panic("colLines should have only 1")
	}
	return colLines[0]
}

func Day3Part2() error {
	lines, err := util.GetFileContentsSplit("day3.txt")
	if err != nil {
		fmt.Println("Err", err)
		return err
	}

	val1 := findPart2(lines, true)
	val2 := findPart2(lines, false)

	val1Num, _ := binaryToDecimal(val1)
	val2Num, _ := binaryToDecimal(val2)

	log.Println(val1, val2, val1Num*val2Num)
	return nil
}
