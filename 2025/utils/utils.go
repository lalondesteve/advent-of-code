// Package utils implement functions that might be reused in the different puzzles
package utils

import (
	"bufio"
	"log"
	"os"
)

func GetScanner() (*os.File, *bufio.Scanner) {
	if len(os.Args) < 2 {
		log.Fatal("missing input file")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Error opening file", err)
	}

	scanner := bufio.NewScanner(file)
	return file, scanner
}

func Modulo(value int, div int) int {
	return (value%div + div) % div
}
