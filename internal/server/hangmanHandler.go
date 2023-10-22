package server

import (
	"fmt"
	"unicode"
	"strings"
)

var AlreadyFoundLettersSlice []string

func stringExistsInSlice(target string, slice []string) bool { //PRIVATE
    for _, element := range slice {
        if element == target {
            return true
        }
    }
    return false
}

//GLOBAL
func HangmanHandler(inVisable_array []string, visable_array []string, word string, guess string) (bool, []string, string) { //return new blank list state, return if word found or not -> then server looked
	var letterWasFound bool = false
	var info string;
	//input validator
	for _, char := range guess {
		if !unicode.IsLetter(char) && !unicode.IsSpace(char) { 
			info = fmt.Sprintf("'%v', '%v' is not a valid character, should be Alphabetic.", inVisable_array, string(char)) //case non-letter char
			return true, inVisable_array, info
		}
	}


	//result validator
	if len(guess) <= 0 {
		info = fmt.Sprintf("%v, You didn't entered anything, should be Alphabetic", inVisable_array) //case nothing was entered (X should include space right now space alone will return 3)
		return false, inVisable_array, info
	} else if len(guess) > 1 {
		if AlreadyFoundLettersSlice == nil {
			AlreadyFoundLettersSlice = []string{}
		} 
		if stringExistsInSlice(guess, AlreadyFoundLettersSlice) { //case letter already used not hurt hm
			info = fmt.Sprintf("%v, You already tried using '%v'.", inVisable_array, guess)
			return true, inVisable_array, info
		}
		AlreadyFoundLettersSlice = append(AlreadyFoundLettersSlice, guess)
		if strings.EqualFold(guess, word) {
			copy(inVisable_array[:], visable_array[:])
			info = fmt.Sprintf("%v, You won! the word was %v", inVisable_array, word) //case won 3
			return true, inVisable_array, info
		} else {
			info = fmt.Sprintf("%v, '%v' is an incorrect word.", inVisable_array, guess) //case 66 wrong full word
			return false, inVisable_array, info
		}
	} else {
		//CASE ONE 1 LETTER
		if AlreadyFoundLettersSlice == nil {
			AlreadyFoundLettersSlice = []string{}
		} 
		if stringExistsInSlice(guess, AlreadyFoundLettersSlice) { //case letter already used not hurt hm
			info = fmt.Sprintf("%v, You already tried using '%v'.", inVisable_array, guess)
			return true, inVisable_array, info
		}
		AlreadyFoundLettersSlice = append(AlreadyFoundLettersSlice, guess)

		for index, _ := range visable_array {
			if strings.EqualFold(guess, visable_array[index]) {
				inVisable_array[index] = visable_array[index]
				letterWasFound = true
			} else if !letterWasFound { //case single letter is wrong
				info = fmt.Sprintf("%v, '%v' is an incorrect letter.", inVisable_array, guess)
				return false, inVisable_array, info
			}
		}
	}

	//case looking if all letters
	for _, element := range inVisable_array {
		if element == "_" {
			if letterWasFound { //CASE input letter is true
				info = fmt.Sprintf("%v, Correct letter '%v'.", inVisable_array, guess)
				return true, inVisable_array, info
			}
			
			info = fmt.Sprintf("%v", inVisable_array) //case found _ in sentence means not complete
			return false, inVisable_array, info
		}
	}

	info = fmt.Sprintf("%v, You won! the word was %v", inVisable_array, word) //case won 3
	return true, inVisable_array, info
}