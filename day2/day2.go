package day2

import (
	"AdventOfCode2021/util"
	"fmt"
	"strconv"
	"strings"
)

func getFileContents() ([]string, error) {
	return util.GetFileContentsSplit("input2.txt")
}

func Day2Part1() {
	fileSplit, err := getFileContents()
	if err != nil {
		fmt.Println("Err", err)
		return
	}
	horizontalVal := 0
	depthVal := 0
	for _, value := range fileSplit {
		rowSplit := strings.Split(value, " ")
		action := rowSplit[0]
		number, err := strconv.Atoi(rowSplit[1])
		if err != nil {
			fmt.Println("Err", err)
			return
		}

		if action == "forward" {
			horizontalVal += number
		}
		if action == "up" {
			depthVal -= number
		}
		if action == "down" {
			depthVal += number
		}
	}
	fmt.Printf("depth %v, horizontal %v, combined %v", depthVal, horizontalVal, depthVal*horizontalVal)
}

func Day2Part2() {
	fileSplit, err := getFileContents()
	if err != nil {
		fmt.Println("Err", err)
		return
	}
	horizontalVal := 0
	depthVal := 0
	aimVal := 0
	for _, value := range fileSplit {
		rowSplit := strings.Split(value, " ")
		action := rowSplit[0]
		number, err := strconv.Atoi(rowSplit[1])
		if err != nil {
			fmt.Println("Err", err)
			return
		}

		if action == "forward" {
			horizontalVal += number
			depthVal += aimVal * number
		}
		if action == "up" {
			aimVal -= number
		}
		if action == "down" {
			aimVal += number
		}
	}
	fmt.Printf("depth %v, horizontal %v, combined %v", depthVal, horizontalVal, depthVal*horizontalVal)
}
