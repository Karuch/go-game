package server

import (
	"fmt"
	"unicode"
	"strings"
)

func HangmanHandler(inVisable_array []string, visable_array []string, word string, guess string) (bool, []string, string) { //return new blank list state, return if word found or not -> then server looked
	
	fmt.Println(inVisable_array, visable_array, word, guess)

	var info string;
	//input validator
	for _, char := range guess {
		if !unicode.IsLetter(char) && !unicode.IsSpace(char) { 
			info = "1" //case non-letter char
			fmt.Println(info)
			return false, inVisable_array, info
		}
	}


	//result validator
	if len(guess) <= 0 {
		info = "2" //case nothing was entered (X should include space right now space alone will return 3)
		fmt.Println(info)
		return false, inVisable_array, info
	} else if len(guess) > 1 {
		if strings.EqualFold(guess, word) {
			info = "3"
			fmt.Println(info)
			copy(inVisable_array[:], visable_array[:])
			return true, inVisable_array, info
		} else {
			return false, inVisable_array, info
		}
	} else {
		//CASE ONE 1 LETTER
		for index, _ := range visable_array {
			if strings.EqualFold(guess, visable_array[index]) {
				inVisable_array[index] = visable_array[index]
			}
		}
	}

	//case looking if all letters
	for _, element := range inVisable_array {
		fmt.Println(element)
		if element == "_" {
			info = "4" //case found _ in sentence means not complete
			return false, inVisable_array, info
		}
	}


	info = "3" //case won
	return true, inVisable_array, info
}