package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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

func GetMovement(s string) int {
	d := strings.ToLower(s[:1])
	v, err := strconv.Atoi(s[1:])
	if err != nil {
		log.Fatal("invalid int found in line", s)
	}
	switch d {
	case "l":
		return -v
	case "r":
		return v
	default:
		log.Fatal("Invalid direction found in line", s)
	}
	return 0
}

func Modulo(value int, div int) int {
	return (value%div + div) % div
}

func main() {
	file, scanner := GetScanner()
	defer file.Close()

	dial := 50
	answer1 := 0
	answer2 := 0

	for scanner.Scan() {
		m := GetMovement(scanner.Text())

		// move the dial
		nextPosition := dial + m
		var count int

		// check if we crossed zerp and how many times
		if m > 0 {
			count = int(math.Floor(float64(nextPosition)/100)) - int(math.Floor(float64(dial)/100))
		} else if m < 0 {
			count = int(math.Ceil(float64(dial)/100)) - int(math.Ceil(float64(nextPosition)/100))
		}
		// fmt.Println(dial, m, nextPosition, count)
		// if dial == 0 || nextPosition == 100 {
		// 	fmt.Println("===========")
		// }
		answer2 += count

		dial = Modulo(nextPosition, 100)

		if dial == 0 {
			answer1 += 1
		}
	}
	fmt.Println("answer pt1", answer1)
	fmt.Println("answer pt2", answer2)
}
