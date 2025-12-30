package main

import (
	"fmt"
	"utils"
)

func main() {
	file, scanner := utils.GetScanner()
	defer file.Close()

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
