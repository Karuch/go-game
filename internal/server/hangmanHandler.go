package server

import (
	"fmt"
	"unicode"
	"strings"
)

func HangmanHandler(){
	
	//initiallize blank and word list
	var word string = "Hello"
	visable_array := strings.Split(word, "")
	inVisable_array := make([]string, len(visable_array))
	copy(inVisable_array[:], visable_array[:])
	for i := 0; i < len(inVisable_array); i++ {
        inVisable_array[i] = "_"
    }
	fmt.Println(inVisable_array, visable_array)

	
	//input validator
	var guess string = "He.llo1"
	for _, char := range guess {
		if !unicode.IsLetter(char) { //what about space?
			fmt.Printf("Non-letter character: %c\n", char)
		}
	}


	//result validator
	if len(guess) <= 0 {
		//DO SOMETHING
	} else if len(guess) < 1 {
		if guess == word {
			//DO SOMETHING
		} else {
			//DO SOMETHING
		}
	} else {
		//CASE ONE 1 LETTER
	}

	
}