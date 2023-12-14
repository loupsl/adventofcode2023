package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("C:/Users/pelis/Documents/Mines2A/Programminglang/adventofcode2023/day14/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	res := []string{}

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	platform := [][]string{}
	for _, ligne := range res {
		platform = append(platform, convert_line(ligne))
	}

	fmt.Println(sum(platform))
}

func convert_line(line string) []string {
	convert := []string{}
	for _, char := range line {
		convert = append(convert, string(char))
	}
	return convert
}

func dist_first_rock_above(line0 int, column0 int, platform [][]string) int {
	for i := line0 - 1; i >= 0; i-- {
		if platform[i][column0] == "#" {
			return line0 - i - 1
		}
	}
	return -1
}

func nbr_rounds_before_rock(line0 int, column0 int, dist_rock int, platform [][]string) int {
	if dist_rock != -1 {
		if dist_rock == 0 {
			return 0
		}
		count := 0
		for i := 1; i <= dist_rock; i++ {
			if platform[line0-i][column0] == "O" {
				count++
			}
		}
		return count
	}
	return -1
}

func nbr_rounds_norocks(line0 int, column0 int, platform [][]string) int {
	count := 0
	for i := line0 - 1; i >= 0; i-- {
		if platform[i][column0] == "O" {
			count++
		}
	}
	return count
}

func move_0(line0 int, column0 int, platform [][]string) int {
	dist_rock := dist_first_rock_above(line0, column0, platform)
	if dist_rock != -1 {
		nbr_rounds_above := nbr_rounds_before_rock(line0, column0, dist_rock, platform)
		return line0 - dist_rock + nbr_rounds_above
	} else {
		nbr_rounds_above := nbr_rounds_norocks(line0, column0, platform)
		return nbr_rounds_above
	}
}

func sum(platform [][]string) int {
	sum := 0
	for i := 0; i < len(platform); i++ {
		for j := 0; j < len(platform[i]); j++ {
			if platform[i][j] == "O" {
				sum += len(platform) - move_0(i, j, platform)
			}
		}
	}
	return sum
}
