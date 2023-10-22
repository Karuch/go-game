package server

import (
	"fmt"
	"unicode"
	"strings"
)

func HangmanHandler(inVisable_array []string, visable_array []string, word string, guess string, test *&[]string) (bool, []string, string) { //return new blank list state, return if word found or not -> then server looked
	var letterWasFound bool = false
	fmt.Println(inVisable_array, visable_array, word, guess)
	append(test, "gay")
	var info string;
	//input validator
	for _, char := range guess {
		if !unicode.IsLetter(char) && !unicode.IsSpace(char) { 
			info = fmt.Sprintf("'%v', '%v' is not a valid character, should be Alphabetic.", inVisable_array, string(char)) //case non-letter char
			fmt.Println(info)
			return false, inVisable_array, info
		}
	}


	//result validator
	if len(guess) <= 0 {
		info = fmt.Sprintf("%v, You didn't entered anything, should be Alphabetic", inVisable_array) //case nothing was entered (X should include space right now space alone will return 3)
		fmt.Println(info)
		return false, inVisable_array, info
	} else if len(guess) > 1 {
		if strings.EqualFold(guess, word) {
			copy(inVisable_array[:], visable_array[:])
			info = fmt.Sprintf("%v, You won! the word was %v", inVisable_array, word) //case won 3
			fmt.Println(info)
			return true, inVisable_array, info
		} else {
			info = fmt.Sprintf("%v, '%v' is an incorrect word.", inVisable_array, guess) //case 66 wrong full word
			return false, inVisable_array, info
		}
	} else {
		//CASE ONE 1 LETTER
		for index, _ := range visable_array {
			if strings.EqualFold(guess, visable_array[index]) {
				inVisable_array[index] = visable_array[index]
				letterWasFound = true
			}
		}
	}

	//case looking if all letters
	for _, element := range inVisable_array {
		fmt.Println(element)
		if element == "_" {
			if letterWasFound { //CASE input letter is true
				info = fmt.Sprintf("%v, Correct letter '%v'.", inVisable_array, guess)
				return true, inVisable_array, info
			}
			
			info = fmt.Sprintf("%v", inVisable_array) //case found _ in sentence means not complete
			fmt.Println(info)
			return false, inVisable_array, info
		}
	}


	info = fmt.Sprintf("%v, You won! the word was %v", inVisable_array, word) //case won 3
	fmt.Println(info)
	return true, inVisable_array, info
}