package programs

import (
	"fmt"
	"strings"
)

//IsPangram to check the input string contains all alphabets from 'a' to 'z'
func IsPangram(inputString string) bool {
	inputString = strings.ToLower(inputString)
	runeMap := map[rune]bool{}
	var i int32
	for i = 97; i < 122; i++ {
		runeMap[i] = false
	}
	for _, runeValue := range inputString {
		runeMap[runeValue] = true
	}
	for i = 97; i < 122; i++ {
		value, ok := runeMap[i]
		if ok && !value {
			return false
		}
	}
	return true
}

//MissingCharacterToPangramString to print characters to make string pangram
func MissingCharacterToPangramString(inputString string) {
	inputString = strings.ToUpper(inputString)
	runeMap := map[rune]bool{}
	var i int32
	for i = 65; i <= 90; i++ {
		runeMap[i] = false
	}
	for _, runeValue := range inputString {
		runeMap[runeValue] = true
	}
	var missingRunes []rune
	for i = 65; i <= 90; i++ {
		value, _ := runeMap[i]
		if !value {
			missingRunes = append(missingRunes, i)
		}
	}
	if len(missingRunes) > 0 {
		fmt.Print("[ ")
		for _, value := range missingRunes {
			fmt.Print(string(value) + " ")
		}
		fmt.Print("]")
	} else {
		fmt.Println("String is a pangram.. no character missing")
	}
}

// RemovePunctuationsFromString to remove punctuations from string
func RemovePunctuationsFromString(inputString string) {
	finalString := ""
	for _, runeValue := range inputString {
		if (runeValue >= 65 && runeValue <= 90) || (runeValue >= 97 && runeValue <= 122) || (runeValue >= 48 && runeValue <= 58) || runeValue == 32 {
			finalString = finalString + string(runeValue)
		}
	}
	fmt.Println(finalString)
}

func FindVowels(inputString string) string {
	inputString = strings.ToLower(inputString)
	finalString := ""
	for _, runeValue := range inputString {
		if runeValue == 'a' || runeValue == 'e' || runeValue == 'i' || runeValue == 'o' || runeValue == 'u' {
			finalString = finalString + string(runeValue)
		}
	}
	return finalString
}

func MergeSlices(a1 []string, a2 []string) []string {
	a1 = append(a1, a2...)
	return a1
}

func CheckPanlindrome(inputString string) bool {
	inputString = strings.ToLower(inputString)
	runeMap := map[rune]struct{}{}
	for _, runeValue := range inputString {
		_, ok := runeMap[runeValue]
		if ok {
			delete(runeMap, runeValue)
		} else {
			runeMap[runeValue] = struct{}{}
		}
	}
	mapLength := len(runeMap)
	if mapLength == 0 || mapLength == 1 {
		return true
	}
	return false
}

type Test interface {
	add() int
	subtract() int
}

//PrintNNumbers print n numbers sequentially without using loop
func PrintNNumbers(i, n int) {
	if i == n {
		return
	}
	i = i + 1
	fmt.Println(i)
	PrintNNumbers(i, n)
}

/* You are given a string S. Count the number of occurrences of all the digits in the string S.

Input:
First line contains string S.
e.g. 77256530

Output:
For each digit starting from 0 to 9, print the count of their occurrences in the string S. So, print  lines, each line containing 2 space separated integers. First integer i and second integer count of occurrence of i. See sample output for clarification.
o/p for e.g. 77256530
0 1
1 0
2 1
3 1
4 0
5 2
6 1
7 2
8 0
9 0

*/

func PrintDigitCountInDigitString(str string) {
	digitmap := map[rune]int{rune('0'): 0, rune('1'): 0, rune('2'): 0, rune('3'): 0, rune('4'): 0, rune('5'): 0, rune('6'): 0, rune('7'): 0, rune('8'): 0, rune('9'): 0}
	for _, runeVal := range str {
		val, ok := digitmap[runeVal]
		if !ok {
			digitmap[runeVal] = 1
		} else {
			digitmap[runeVal] = val + 1
		}
	}
	a := "0123456789"
	for _, runeval := range a {
		fmt.Println(string(runeval), digitmap[runeval])
	}
}

//input k=3, str :="DAbcefgh",
//output:="cefghDAb"

func LeftShiftCharsInString(str string, k int) string {
	k = k % len(str)
	strArr := strings.Split(str, "")
	if len(str) <= 1 {
		return str
	}
	strslice1 := strArr[k:]
	strSlice2 := strArr[0:k]

	finalStr := ""
	for _, val := range strslice1 {
		finalStr = finalStr + val
	}
	for _, val := range strSlice2 {
		finalStr = finalStr + val
	}
	fmt.Println(finalStr)
	return finalStr
}
