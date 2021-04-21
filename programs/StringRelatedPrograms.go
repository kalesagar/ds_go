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
	for i=97; i<122; i++{
	     runeMap[i]=false
	}
	for _, runeValue := range inputString{
		runeMap[runeValue] = true
	}
	for i=97; i<122; i++{
		value, ok := runeMap[i]
		if ok && !value{
			return false
		}
	}
	return true
}

//MissingCharacterToPangramString to print characters to make string pangram
func MissingCharacterToPangramString(inputString string){
	inputString = strings.ToUpper(inputString)
	runeMap := map[rune]bool{}
	var i int32
	for i=65; i<=90;i++{
		runeMap[i]=false
	}
	for _, runeValue := range inputString{
		runeMap[runeValue] = true
	}
	var missingRunes []rune
	for i=65; i<=90;i++{
		value, _ := runeMap[i]
		if !value{
			missingRunes = append(missingRunes, i)
		}
	}
	if len(missingRunes) > 0{
		fmt.Print("[ ")
		for _, value := range missingRunes{
			fmt.Print(string(value) + " ")
		}
		fmt.Print("]")
	}else{
		fmt.Println("String is a pangram.. no character missing")
	}
}
