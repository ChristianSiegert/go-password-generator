package main

import (
	"flag"
	"fmt"
)

var (
	useLowerCaseCharacters = flag.Bool("lower", true, "Use lower-case characters a-z")
	useUpperCaseCharacters = flag.Bool("upper", true, "Use upper-case characters A-Z")
	useNumbers             = flag.Bool("numbers", true, "Use numbers 0-9")
	useSpecialCharacters   = flag.Bool("special", false, "Use special characters !§$%&=?,.-;:_")
	userCharacters         = flag.String("own", "", "Custom set of characters to use")
	passwordCount          = flag.Int("count", 3, "Number of passwords to generate")
	passwordLength         = flag.Int("length", 10, "Length of passwords")
	verbose                = flag.Bool("verbose", false, "Display additional information")
)

func main() {
	flag.Parse()

	generator, error := NewGenerator(*useLowerCaseCharacters, *useUpperCaseCharacters, *useNumbers, *useSpecialCharacters, *userCharacters, *passwordLength)

	if error != nil {
		fmt.Println("Error:", error)
		return
	}

	if *verbose {
		fmt.Println(additionalInformation(generator))
	}

	for i := 0; i < *passwordCount; i++ {
		password, error := generator.Password(*passwordLength)

		if error != nil {
			fmt.Println("Error:", error)
			continue
		}

		fmt.Println(password)
	}
}

func additionalInformation(generator *Generator) string {
	passwordCountInWords := ""

	if *passwordCount == 1 {
		passwordCountInWords = "1 password"
	} else {
		passwordCountInWords = fmt.Sprintf("%d passwords", *passwordCount)
	}

	return fmt.Sprintf("Generating %s with a length of %d from the following characters: %s", passwordCountInWords, *passwordLength, generator.characters)
}
