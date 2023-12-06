package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("C:/Users/pelis/Documents/Mines2A/Programminglang/adventofcode2023/day6/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []string{}
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	times := convert_int(strings.Split(strings.Split(res[0], ":        ")[1], "     "))
	distances := convert_int(strings.Split(strings.Split(res[1], ":   ")[1], "   "))

	fmt.Println(result(times, distances))

	single_time := ""
	for _, val := range strings.Split(strings.Split(res[0], ":        ")[1], "     ") {
		single_time += val
	}
	s_time, err := strconv.Atoi(single_time)
	single_distance := ""
	for _, val := range strings.Split(strings.Split(res[1], ":   ")[1], "   ") {
		single_distance += val
	}
	s_distance, err := strconv.Atoi(single_distance)

	fmt.Println(count_waystowin(s_time, s_distance))
}

func convert_int(tab []string) []int {
	tabint := []int{}
	for _, val := range tab {
		value, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Erreur de conversion")
			return tabint
		}
		tabint = append(tabint, value)
	}
	return tabint
}

// distance if we held the button for "time_heldbutton" in a race that lasts "time"
func result_distance(time int, time_heldbutton int) int {
	return time_heldbutton * (time - time_heldbutton)
}

func count_waystowin(time int, record int) int {
	count := 0
	for i := 0; i <= time; i++ {
		time_heldbutton := i
		distance := result_distance(time, time_heldbutton)
		if distance > record {
			count++
		}
	}
	return count
}

func result(times []int, distances []int) int {
	result := 1
	for i := 0; i < 4; i++ {
		result = result * count_waystowin(times[i], distances[i])
	}
	return result
}
