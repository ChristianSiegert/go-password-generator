package main

import (
	"bytes"
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

type Generator struct {
	characters     string   // Characters a password is allowed to contain
	characterRunes []rune   // Characters cached as runes
	upperRandLimit *big.Int // Upper limit for random number generator
}

const (
	lowerCaseCharacters = "abcdefghijklmnopqrstuvwxyz"
	numbers             = "0123456789"
	specialCharacters   = "!§$%&=?,.-;:_"
)

func NewGenerator(useLowerCaseCharacters, useUpperCaseCharacters, useNumbers, useSpecialCharacters bool, customCharacters string, passwordLength int) (generator *Generator, error error) {
	if passwordLength <= 0 {
		error = errors.New("Passwords must be at least one character long.")
		return
	}

	var characterBuffer bytes.Buffer

	if useLowerCaseCharacters {
		characterBuffer.WriteString(lowerCaseCharacters)
	}

	if useUpperCaseCharacters {
		characterBuffer.WriteString(strings.ToUpper(lowerCaseCharacters))
	}

	if useNumbers {
		characterBuffer.WriteString(numbers)
	}

	if useSpecialCharacters {
		characterBuffer.WriteString(specialCharacters)
	}

	characterBuffer.WriteString(customCharacters)

	if characterBuffer.Len() <= 0 {
		error = errors.New("Cannot generate password without any allowed characters.")
		return
	}

	generator = &Generator{characters: characterBuffer.String()}
	generator.characterRunes = []rune(generator.characters)
	generator.upperRandLimit = big.NewInt(int64(len(generator.characterRunes)))

	return
}

func (generator *Generator) Password(length int) (string, error) {
	var passwordBuffer bytes.Buffer

	for i := 0; i < length; i++ {
		randomIndex, error := rand.Int(rand.Reader, generator.upperRandLimit)

		if error != nil {
			return "", error
		}

		passwordBuffer.WriteRune(generator.characterRunes[randomIndex.Int64()])
	}

	return passwordBuffer.String(), nil
}
