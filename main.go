package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
)

var characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var runes = []rune(characters)
var upperRandomLimit = big.NewInt(int64(len(runes)))

func main() {
	for i := 0; i < 10; i++ {
		password, error := password(10)

		if error != nil {
			fmt.Println("Couldn't generate password:", error)
			continue
		}

		fmt.Println(password)
	}
}

func password(length int) (string, error) {
	var buffer bytes.Buffer

	for i := 0; i < length; i++ {
		index, error := rand.Int(rand.Reader, upperRandomLimit)

		if error != nil {
			return "", error
		}

		buffer.WriteRune(runes[index.Int64()])
	}

	return buffer.String(), nil
}
