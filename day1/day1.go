package day1

import (
	"AdventOfCode2021/util"
	"fmt"
	"strconv"
)

func GetPart1FileContents() ([]string, error) {
	return util.GetFileContentsSplit("input1.txt")
}

func StringArrayToIntArray(stringArray []string) ([]int64, error) {
	result := make([]int64, 0)
	for _, valueStr := range stringArray {
		value, err := strconv.ParseInt(valueStr, 10, 0)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		result = append(result, value)
	}

	return result, nil
}

func Day1Part1() {
	fileSplit, err := GetPart1FileContents()
	if err != nil {
		fmt.Println("Err", err)
		return
	}
	numbers, err := StringArrayToIntArray(fileSplit)
	if err != nil {
		fmt.Println("Err", err)
		return
	}
	var oldVal int64 = 0
	increments := 0
	for index, value := range numbers {
		if index > 0 && oldVal < value {
			increments++
		}
		oldVal = value
	}

	fmt.Printf("Increments %v", increments)
}

func Day1Part2() {
	fileSplit, err := GetPart1FileContents()
	if err != nil {
		fmt.Println("Err", err)
		return
	}
	numbers, err := StringArrayToIntArray(fileSplit)
	if err != nil {
		fmt.Println("Err", err)
		return
	}

	increments := 0
	for index, value := range numbers {
		if index >= 3 {
			currentSum := value + numbers[index-1] + numbers[index-2]
			oldSum := numbers[index-1] + numbers[index-2] + numbers[index-3]
			if currentSum > oldSum {
				increments++
			}
		}
	}

	fmt.Printf("Increments %v", increments)
}
