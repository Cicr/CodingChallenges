package main

import (
	"fmt"
	"strings"
)

func sliceToString(strArr []string) string {
	result := ""
	for _, str := range strArr {
		result += str
	}
	return result
}

func isVowel(ch byte) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, rune(ch))
}

func MatrixChallenge(strArr []string) string {

	//fmt.Println(fmt.Sprintf("Full Array is %v with %d rows", strArr, len(strArr)))
	next_i := 0
	next_j := 0
	for i := 0; i < len(strArr); i++ {
		limit_i := len(strArr) - 1
		next_i = 1
		for j := 0; j < len(strArr[i]); j++ {
			limit_j := len(strArr[i]) - 1
			next_j = 1

			if i == limit_i {
				next_i = 0
				_ = next_i
			}

			if j == limit_j {
				next_j = 0
				_ = next_j
			}

			item := string(strArr[i][j])
			item_next := string(strArr[i+next_i][j+next_j])
			Vowelcheck := isVowel(strArr[i][j])
			Vowelcheck_next := isVowel(strArr[i+next_i][j+next_j])
			_ = Vowelcheck
			_ = Vowelcheck_next
			_ = item
			_ = item_next
			//fmt.Println(fmt.Sprintf("Array item is %s  ", item))
			//fmt.Println(fmt.Sprintf("Array item is %s and check is Vowel is %t ", item, Vowelcheck))
			//fmt.Println(fmt.Sprintf("NEXT Array item is %s and check is Vowel is %t ", item_next, Vowelcheck_next))
                        //Algoritmo facil para calcular una sub-matriz 2x2  fila y fila mas 1 y columna y columna mas 1 son igual a 2x2 items
			if isVowel(strArr[i][j]) &&
				isVowel(strArr[i][j+1]) &&
				isVowel(strArr[i+1][j]) &&
				isVowel(strArr[i+1][j+1]) {
				return fmt.Sprintf("%d-%d", i, j)
			}

		}
	}

	return "end of function"
}

func main() {

	strArr1 := []string{"abcd", "eikr", "oufj"}

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(MatrixChallenge(strArr1))

}
