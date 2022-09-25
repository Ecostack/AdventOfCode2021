package ctf

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const BASIC_MOD_2_NUMBERS = "268 413 110 190 426 419 108 229 310 379 323 373 385 236 92 96 169 321 284 185 154 137 186"

const BASIC_MOD_2_ALPHABET = ".abcdefghijklmnopqrstuvwxyz0123456789_"

func modInverse(a int, m int) int {
	for x := 1; x < m; x++ {
		if ((a%m)*(x%m))%m == 1 {
			return x
		}
	}
	return 1
}

func TestCTFBasicMod2(t *testing.T) {
	splitted := strings.Split(BASIC_MOD_2_NUMBERS, " ")

	numbers := make([]int, 0)
	for _, s := range splitted {
		val, _ := strconv.Atoi(s)
		numbers = append(numbers, val)
	}
	result := ""
	fmt.Print("picoCTF{")
	for _, number := range numbers {

		modInv := modInverse(number%41, 41)

		result = result + string(BASIC_MOD_2_ALPHABET[modInv])

	}
	//fmt.Scanner()
	fmt.Print(result)
	fmt.Print("}")
}
