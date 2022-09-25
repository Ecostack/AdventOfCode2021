package ctf

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const numbersStr = "54 396 131 198 225 258 87 258 128 211 57 235 114 258 144 220 39 175 330 338 297 288"

const ALPHABET = "abcdefghijklmnopqrstuvwxyz"

func getAlphabet() string {
	res := strings.ToUpper(ALPHABET)
	res = res + "0123456789" + "_"

	return res
}

func TestCTFBasicMod1(t *testing.T) {
	alphabet := getAlphabet()
	splitted := strings.Split(numbersStr, " ")
	numbers := make([]int, 0)
	for _, s := range splitted {
		val, _ := strconv.Atoi(s)
		numbers = append(numbers, val)
	}
	result := ""
	fmt.Print("picoCTF{")
	for _, number := range numbers {
		val := number % 37
		result = result + string(alphabet[val])
	}
	fmt.Print(result)
	fmt.Print("}")
}
