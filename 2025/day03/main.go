package main

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"utils"
)

func GetBatteries() (joltages []string) {
	file, scanner := utils.GetScanner()
	defer file.Close()

	for scanner.Scan() {
		joltages = append(joltages, scanner.Text())
	}
	return joltages
}

func FindMaxJoltages(bank string, amount int) (joltages []int) {
	joltages = make([]int, amount)
	start := 0
	nextStart := 0
	for i := range amount {
		start = nextStart
		end := len(bank) - amount + i

		// move the window to the right
		for j := range bank[start : end+1] {
			idx := start + j
			v, err := strconv.Atoi(bank[idx : idx+1])
			if err != nil {
				log.Fatal("couldn't convert string to int", idx)
			}
			if v > joltages[i] {
				joltages[i] = v
				nextStart = idx + 1
			}
		}
	}
	return joltages
}

func main() {
	joltages := GetBatteries()
	amount := 2
	sum := 0
	for _, b := range joltages {
		j := FindMaxJoltages(b, amount)
		value := 0
		slices.Reverse(j)
		for i, v := range j {
			value += v * int(math.Pow(10, float64(i)))
		}
		sum += value
	}
	fmt.Println("part 1 answer:", sum)
}
