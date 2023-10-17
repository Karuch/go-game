package server

import (
	"fmt"
	"unicode"
	"strings"
)

func HangmanHandler(){
	
	for 0 < 1 {
	
		//initiallize blank and word list
		var word string = "He llo"
		visable_array := strings.Split(word, "")
		inVisable_array := make([]string, len(visable_array))
		copy(inVisable_array[:], visable_array[:])
		for i := 0; i < len(inVisable_array); i++ {
			inVisable_array[i] = "_"
		}
		fmt.Println(inVisable_array, visable_array)

		
		//input validator
		var guess string = "h e llo"
		for _, char := range guess {
			if !unicode.IsLetter(char) && !unicode.IsSpace(char) { 
				fmt.Printf("Non-letter character: %c\n", char)
			}
		}


		//result validator
		if len(guess) <= 0 {
			//DO SOMETHING
		} else if len(guess) < 1 {
			if strings.EqualFold(guess, word) {
				//DO SOMETHING IF WORD == GUESS i think i will wait for loop
			} else {
				//DO SOMETHING IN CASE NO OUTPUT i think i will wait for loop
			}
		} else {
			//CASE ONE 1 LETTER
			for index, _ := range visable_array {
				if strings.EqualFold(guess, visable_array[index]) {
					inVisable_array[index] = visable_array[index]
				}
			}
		}

		fmt.Println(inVisable_array)
	}
}