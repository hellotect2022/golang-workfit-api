package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	input := "1234"
	hashedInput, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error hashing input:", err)
		return
	}

	fmt.Println("SHA256 Hash of", input, hashedInput, string(hashedInput))

	inputPassword := "1234"
	hashedPassword := string(hashedInput)

	err1 := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))

	fmt.Println("Password does match:", err1 == nil)

	fmt.Println("Test!!!")
}
