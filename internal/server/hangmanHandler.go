package server

import (
	"fmt"
	"unicode"
	"strings"
)

var AlreadyFoundLettersSlice []string
var life int = 0
var drawLife int = 6

//GLOBAL
func DrawHM(doNotHurtHM bool, listState []string, info string) (string, int) { 
    if !doNotHurtHM {
		life = life + 1
		drawLife = drawLife - 1
	}

	blankPartsSlice := []string{" ", " ", " ", " ", " ", " "}
    partsSlice := []string{"/", "\\", "/", "\\", "|", "o"}
	for i := 0; i < life; i++ {
		blankPartsSlice[i] = partsSlice[i]
	}

    formattedString := fmt.Sprintf(`
     %v
     |=======\     word: %v
     %v        |    life: %v
    %v%v%v       |    
    %v %v       |
              
    `, info, listState, blankPartsSlice[5], drawLife, blankPartsSlice[0], blankPartsSlice[4],
	blankPartsSlice[3], blankPartsSlice[2], blankPartsSlice[1])
	return formattedString, life
}

//PRIVATE
func stringExistsInSlice(target string, slice []string) bool { 
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
			info = fmt.Sprintf("You have '%v' in '%v' which is not a valid character, should be Alphabetic or Space.", string(char), guess) //case non-letter char
			return true, inVisable_array, info
		}
	}


	//result validator
	if len(guess) <= 0 {
		info = fmt.Sprintf("You didn't entered anything, should be Alphabetic") //case nothing was entered (X should include space right now space alone will return 3)
		return false, inVisable_array, info
	} else if len(guess) > 1 {
		if AlreadyFoundLettersSlice == nil {
			AlreadyFoundLettersSlice = []string{}
		} 
		if stringExistsInSlice(guess, AlreadyFoundLettersSlice) { //case letter already used not hurt hm
			info = fmt.Sprintf("You already tried using '%v'.", guess)
			return true, inVisable_array, info
		}
		AlreadyFoundLettersSlice = append(AlreadyFoundLettersSlice, guess)
		if strings.EqualFold(guess, word) {
			copy(inVisable_array[:], visable_array[:])
			info = fmt.Sprintf("You won! the word was %v", word) //case won 3
			return true, inVisable_array, info
		} else {
			info = fmt.Sprintf("'%v' is an incorrect word.", guess) //case 66 wrong full word
			return false, inVisable_array, info
		}
	} else {
		//CASE ONE 1 LETTER
		if guess == " " { //handle single space input case
			info = fmt.Sprintf("'space' alone is not a valid character.")
			return true, inVisable_array, info
		}
		if AlreadyFoundLettersSlice == nil {
			AlreadyFoundLettersSlice = []string{}
		} 
		if stringExistsInSlice(guess, AlreadyFoundLettersSlice) { //case letter already used not hurt hm
			info = fmt.Sprintf("You already tried using '%v'.", guess)
			return true, inVisable_array, info
		}
		AlreadyFoundLettersSlice = append(AlreadyFoundLettersSlice, guess)

		for index, _ := range visable_array {
			if strings.EqualFold(guess, visable_array[index]) {
				inVisable_array[index] = visable_array[index]
				letterWasFound = true
			}
		}
		if !letterWasFound { //case single letter is wrong
			info = fmt.Sprintf("'%v' is an incorrect letter.", guess)
			return false, inVisable_array, info
		}
	}

	//case looking if all letters
	for _, element := range inVisable_array {
		if element == "_" {
			if letterWasFound { //CASE input letter is true
				info = fmt.Sprintf("Correct letter '%v'.", guess)
				return true, inVisable_array, info
			}
			//info = fmt.Sprintf("%v REACHEDDD", inVisable_array) //case found _ in sentence means not complete
			//return false, inVisable_array, info
			//^ there's a good chance this case will never reach because the only one possibility for this is
			//single wrong letter n "_" exists, which case single letter is wrong already handling, I'll ignore it for now
		}
	}

	info = fmt.Sprintf("You won! the word was %v", word) //case won 3
	return true, inVisable_array, info
}