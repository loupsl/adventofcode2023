package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("C:/Users/pelis/Documents/Mines2A/Programminglang/adventofcode2023/day1/input1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []string{}
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	fmt.Println(sum(calibration(res)))
	fmt.Println(sum(calibration_letter(res)))
}

func calibration(file []string) []int {
	var count []int
	for _, val := range file {
		count_val := firstint(val)*10 + lastint(val)
		count = append(count, count_val)
	}
	return count
}

func firstint(mot string) int {
	for _, char := range mot {
		if unicode.IsDigit(char) {
			return int(char - '0')
		}
	}
	return 0
}

func lastint(mot string) int {
	for i := len(mot) - 1; i >= 0; i-- {
		char := rune(mot[i])
		if unicode.IsDigit(char) {
			return int(char - '0')
		}
	}

	return 0
}

func sum(list []int) int {
	sum := 0
	for i := 0; i < len(list); i++ {
		sum += list[i]
	}
	return sum
}

func calibration_letter(file []string) []int {
	var count []int
	for _, val := range file {
		count_val := firstintletter(val)*10 + lastintletter(val)
		count = append(count, count_val)
	}
	return count
}

func firstintletter(mot string) int {
	for i := 0; i < len(mot); i++ {
		char := rune(mot[i])
		if unicode.IsDigit(char) {
			return int(char - '0')
		}
		if i < len(mot)-4 {
			if mot[i:i+5] == "three" {
				return 3
			}
			if mot[i:i+5] == "seven" {
				return 7
			}
			if mot[i:i+5] == "eight" {
				return 8
			}
		}
		if i < len(mot)-3 {
			if mot[i:i+4] == "four" {
				return 4
			}
			if mot[i:i+4] == "five" {
				return 5
			}
			if mot[i:i+4] == "nine" {
				return 9
			}
		}
		if i < len(mot)-2 {
			if mot[i:i+3] == "one" {
				return 1
			}
			if mot[i:i+3] == "two" {
				return 2
			}
			if mot[i:i+3] == "six" {
				return 6
			}
		}
	}
	return 0
}

func lastintletter(mot string) int {
	for i := len(mot) - 1; i >= 0; i-- {
		char := rune(mot[i])
		if unicode.IsDigit(char) {
			return int(char - '0')
		}
		if i > 1 {
			if mot[i-2:i+1] == "one" {
				return 1
			}
			if mot[i-2:i+1] == "two" {
				return 2
			}
			if mot[i-2:i+1] == "six" {
				return 6
			}
		}
		if i > 2 {
			if mot[i-3:i+1] == "four" {
				return 4
			}
			if mot[i-3:i+1] == "five" {
				return 5
			}
			if mot[i-3:i+1] == "nine" {
				return 9
			}
		}
		if i > 3 {
			if mot[i-4:i+1] == "three" {
				return 3
			}
			if mot[i-4:i+1] == "seven" {
				return 7
			}
			if mot[i-4:i+1] == "eight" {
				return 8
			}
		}
	}
	return 0
}
