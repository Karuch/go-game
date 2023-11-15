package main

import (
	"fmt"
	"test/creator"
	"test/validator"
	"time"
)

func main(){
	var id int = 0
	for id < 5 {
		id++ 
		creator.Creator(id)
		fmt.Println("\n")
	}

	validator.Validator_accesstoken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjMiLCJmaXJzdCI6IiIsImxhc3QiOiIiLCJleHAiOjE3MDAwNDU5ODYsImlhdCI6MTcwMDA0NTA4Nn0.S8Jb5-KhPXlDkQM76699xuPJhkyiNw7RqvLra-oC0Pg")
	validator.Translator("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjMiLCJmaXJzdCI6IiIsImxhc3QiOiIiLCJleHAiOjE3MDAwNDU5ODYsImlhdCI6MTcwMDA0NTA4Nn0.S8Jb5-KhPXlDkQM76699xuPJhkyiNw7RqvLra-oC0Pg")
	validator.Time()	
	fmt.Println(time.DateTime)
}