package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"utils"
)

type Range struct {
	start int
	end   int
}

func GetRanges() (ranges []*Range) {
	file, scanner := utils.GetScanner()
	defer file.Close()

	for scanner.Scan() {
		rs := strings.SplitSeq(scanner.Text(), ",")
		for s := range rs {
			r := new(Range)
			values := strings.Split(s, "-")
			if len(values) > 2 {
				log.Fatal("Range has too many values", values)
			}
			var err error
			r.start, err = strconv.Atoi(strings.TrimSpace(values[0]))
			if err != nil {
				log.Fatal("cannot convert string to int", values[0])
			}
			r.end, err = strconv.Atoi(strings.TrimSpace(values[1]))
			if err != nil {
				log.Fatal("cannot convert string to int", values[0])
			}
			ranges = append(ranges, r)
		}
	}
	return ranges
}

func CheckValidity(r *Range) (invalidIDs []int) {
	for i := r.start; i <= r.end; i++ {
		s := strconv.Itoa(i)
		length := len(s)
		if length%2 != 0 {
			continue
		}
		if s[:length/2] == s[length/2:] {
			invalidIDs = append(invalidIDs, i)
		}

	}

	return invalidIDs
}

func part1(ranges []*Range) {
	allInvalidIDs := []int{}
	for _, r := range ranges {
		invalidIDs := CheckValidity(r)
		allInvalidIDs = append(allInvalidIDs, invalidIDs...)
	}
	sum := 0
	for _, i := range allInvalidIDs {
		sum += i
	}
	fmt.Println("part1 solution: ", sum)
}

func CheckValidity2(r *Range) (invalidIDs []int) {
	for i := r.start; i <= r.end; i++ {
		s := strconv.Itoa(i)
		length := len(s)

	outer:
		for j := 2; j <= length; j++ {
			if length%j == 0 {
				substring := s[:length/j]
				invalid := false
			inner:
				for k := 2; k <= j; k++ {
					if substring != s[(length/j)*(k-1):k*length/j] {
						invalid = false
						break inner
					}

					invalid = true
				}
				if invalid {
					invalidIDs = append(invalidIDs, i)
					break outer
				}
			}
		}

	}

	return invalidIDs
}

func part2(ranges []*Range) {
	allInvalidIDs := []int{}
	for _, r := range ranges {
		invalidIDs := CheckValidity2(r)
		allInvalidIDs = append(allInvalidIDs, invalidIDs...)
	}
	sum := 0
	for _, i := range allInvalidIDs {
		sum += i
	}
	fmt.Println("part2 solution: ", sum)
}

func main() {
	ranges := GetRanges()
	part1(ranges)
	part2(ranges)
}
