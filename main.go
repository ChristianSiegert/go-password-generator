package main

import (
	"flag"
	"fmt"
)

var (
	customCharacters       = flag.String("custom", "", "Custom set of characters to use")
	passwordCount          = flag.Int("count", 10, "Number of passwords to generate")
	passwordLength         = flag.Int("length", 20, "Length of passwords")
	useLowerCaseCharacters = flag.Bool("lower", false, "Use lower-case characters a-z")
	useNumbers             = flag.Bool("numbers", false, "Use numbers 0-9")
	useSpecialCharacters   = flag.Bool("special", false, "Use special characters !§$%&=?,.-;:_")
	useUpperCaseCharacters = flag.Bool("upper", false, "Use upper-case characters A-Z")
	verbose                = flag.Bool("verbose", false, "Display additional information")
)

func main() {
	flag.Parse()

	generator, error := NewGenerator(*useLowerCaseCharacters, *useUpperCaseCharacters, *useNumbers, *useSpecialCharacters, *customCharacters, *passwordLength)

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
