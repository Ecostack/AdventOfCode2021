package day5

import (
	"AdventOfCode2021/util"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"testing"
)

type LineDefinition struct {
	x1, x2, y1, y2 int64
	diagonal       bool
	angle          float64
	is45Degree     bool
}

var regexLine = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

func parseLineToStruct(line string) *LineDefinition {
	matches := regexLine.FindAllStringSubmatch(line, -1)
	for _, v := range matches {
		x1, _ := strconv.ParseInt(v[1], 10, 0)
		y1, _ := strconv.ParseInt(v[2], 10, 0)
		x2, _ := strconv.ParseInt(v[3], 10, 0)
		y2, _ := strconv.ParseInt(v[4], 10, 0)

		return &LineDefinition{
			x1: x1,
			x2: x2,
			y1: y1,
			y2: y2,
		}
	}
	return nil
}

func parseLines(lines []string, diagonal bool) ([]LineDefinition, int64, int64) {
	linesParsed := make([]LineDefinition, 0)

	xLarge := int64(0)
	yLarge := int64(0)
	for _, line := range lines {
		lineParsed := parseLineToStruct(line)
		if lineParsed != nil {
			lineParsed.diagonal = lineParsed.x1 != lineParsed.x2 && lineParsed.y1 != lineParsed.y2
			if lineParsed.diagonal {
				lineParsed.angle = math.Atan2(float64(lineParsed.y2-lineParsed.y1), float64(lineParsed.x2-lineParsed.x1)) * (180 / math.Pi)
			}
			if math.Abs(lineParsed.angle) == float64(45) || math.Abs(lineParsed.angle) == float64(135) {
				lineParsed.is45Degree = true
			}

			if !diagonal && lineParsed.diagonal {
				continue
			}

			if diagonal && lineParsed.diagonal && !lineParsed.is45Degree {
				continue
				//angle := math.Atan2(float64(lineParsed.y2-lineParsed.y1), float64(lineParsed.x2-lineParsed.x1)) * (180 / math.Pi)
				//
				//if math.Abs(angle) != float64(45) {
				//	fmt.Printf("angle: %v\n", angle)
				//	continue
				//}
			}

			linesParsed = append(linesParsed, *lineParsed)

			if xLarge < lineParsed.x1 {
				xLarge = lineParsed.x1
			}
			if xLarge < lineParsed.x2 {
				xLarge = lineParsed.x2
			}
			if yLarge < lineParsed.y1 {
				yLarge = lineParsed.y1
			}
			if yLarge < lineParsed.y2 {
				yLarge = lineParsed.y2
			}
		}
	}
	return linesParsed, xLarge + 1, yLarge + 1
}

func getCountFromFile(file string, diagonal bool, printMatrix bool) int {
	currentWD, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	linesRaw, err := util.GetFileContentsSplit(currentWD + file)
	if err != nil {
		panic(err)
	}
	linesParsed, xLarge, yLarge := parseLines(linesRaw, diagonal)
	matrix := make([][]int, yLarge)
	for i := int64(0); i < yLarge; i++ {
		matrix[i] = make([]int, xLarge)
	}

	for _, line := range linesParsed {
		if line.diagonal {
			lastRun := false
			x := line.x1
			y := line.y1
			xModify := func() {
				xSmaller := line.x1 < line.x2
				if xSmaller {
					x++
				} else {
					x--
				}

			}
			yModify := func() {
				smaller := line.y1 < line.y2
				if smaller {
					y++
				} else {
					y--
				}
			}

			for {
				matrix[y][x]++
				if lastRun {
					break
				}
				xModify()
				yModify()
				if (x == line.x1 || x == line.x2) && (y == line.y1 || y == line.y2) {
					lastRun = true
				}
			}
			continue
		}

		x := line.x1
		xEqualPass := false
		xModify := func() {
			if line.x1 == line.x2 {
				return
			}
			xSmaller := line.x1 < line.x2
			if xSmaller {
				x++
			} else {
				x--
			}
		}
		for {
			y := line.y1
			yEqualPass := false
			yModify := func() {
				if line.y1 == line.y2 {
					return
				}
				smaller := line.y1 < line.y2
				if smaller {
					y++
				} else {
					y--
				}
			}

			for {
				matrix[y][x]++
				if line.y1 == line.y2 || yEqualPass {
					break
				}

				yModify()
				if line.y1 != line.y2 && y == line.y2 {
					yEqualPass = true
					continue
				}
				if y == line.y2 {
					break
				}
			}
			if line.x1 == line.x2 || xEqualPass {
				break
			}

			xModify()
			if line.x1 != line.x2 && x == line.x2 {
				xEqualPass = true
				continue
			}
			if x == line.x2 {
				break
			}
		}
	}
	count := 0
	for _, row := range matrix {
		for _, column := range row {
			if column > 1 {
				count++
			}
		}
	}

	if printMatrix {
		fmt.Println()
		for _, row := range matrix {
			for _, column := range row {
				if column == 0 {
					fmt.Print(".")
					continue
				}
				fmt.Print(column)
			}
			fmt.Println()
		}
		fmt.Println()
	}

	fmt.Printf("count %v\n", count)
	return count
}

func TestDay5Part1Test(t *testing.T) {
	result := getCountFromFile("/day5_test.txt", false, false)
	if result != 5 {
		t.Fail()
	}
}

func TestDay5Part2Test(t *testing.T) {
	result := getCountFromFile("/day5_test.txt", true, false)
	if result != 12 {
		t.Fail()
	}
}

func TestDay5Part1(t *testing.T) {
	result := getCountFromFile("/day5.txt", false, false)
	if result != 5280 {
		t.Fail()
	}
}

func TestDay5Part2(t *testing.T) {
	result := getCountFromFile("/day5.txt", true, false)
	if result != 16716 {
		t.Fail()
	}
}
