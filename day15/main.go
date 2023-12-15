package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("C:/Users/pelis/Documents/Mines2A/Programminglang/adventofcode2023/day15/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	strings := strings.Split(res[0], ",")

	fmt.Println(total_result(strings))

}

func hash_algorithm(ascii int, currentvalue int) int {
	value := currentvalue + ascii
	value = value * 17
	value = value % 256
	return value
}

func hash_chain(chain string) int {
	value_chain := 0
	for _, char := range chain {
		val := hash_algorithm(int(char), value_chain)
		value_chain = val
	}
	return value_chain
}

func total_result(strings []string) int {
	sum := 0
	for _, chain := range strings {
		sum += hash_chain(chain)
	}
	return sum
}
