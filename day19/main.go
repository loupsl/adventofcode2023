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
	file, err := os.Open("C:/Users/pelis/Documents/Mines2A/Programminglang/adventofcode2023/day19/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res1 := []string{}
	res2 := []string{}
	scanner := bufio.NewScanner(file)
	bool := false
	for scanner.Scan() {
		if scanner.Text() != "" && !bool {
			res1 = append(res1, scanner.Text())
		}
		if scanner.Text() == "" {
			bool = true
		}
		if scanner.Text() != "" && bool {
			res2 = append(res2, scanner.Text())
		}
	}
	workflows := make_workflows(res1)
	ratings := make_ratings(res2)
	//fmt.Println(ratings)
	//fmt.Println(workflows["bfb"])
	fmt.Println(count_acceptance(ratings, workflows))
}

func str_to_int(intstring string) int {
	value, err := strconv.Atoi(intstring)
	if err != nil {
		fmt.Println("Erreur de conversion")
		return value
	}
	return value
}

func make_ratings(res []string) []map[string]int {
	dic := []map[string]int{}
	for i := 0; i < len(res); i++ {
		dic_line := make(map[string]int)
		chain := strings.Split(res[i][1:len(res[i])-1], ",")
		for _, eq := range chain {
			split := strings.Split(eq, "=")
			char := string(split[0])
			val := str_to_int(split[1])
			dic_line[char] = val
		}
		dic = append(dic, dic_line)
	}
	return dic
}

func make_workflows(res []string) map[string][]string {
	dic := map[string][]string{}
	for i := 0; i < len(res); i++ {
		line := strings.Split(res[i][0:len(res[i])-1], "{")
		name := line[0]
		instructions := strings.Split(line[1], ",")
		dic[name] = instructions
	}
	return dic
}

func acceptance(rating map[string]int, workflows map[string][]string) string {
	curr_workflow := workflows["in"]
	bool := false
	ind := 0
	for !bool {
		instruction := curr_workflow[ind]
		int_instr, name := read_instruction(instruction, rating)
		if int_instr == 0 {
			curr_workflow = workflows[name]
			ind = 0
		}
		if int_instr == 1 {
			return "A"
		}
		if int_instr == 2 {
			return "R"
		}
		if int_instr == 3 {
			ind++
		}
	}
	fmt.Println("Problème path")
	return ""
}

func read_instruction(instruction string, rating map[string]int) (int, string) { //0 if workflow, 1 if A, 2 if R, 3 if "next instruction"
	if len(instruction) == 1 {
		if instruction == "A" {
			return 1, "A"
		}
		if instruction == "R" {
			return 2, "R"
		}
	}
	if len(instruction) == 2 || len(instruction) == 3 {
		return 0, instruction
	} else {
		tab_instruct := strings.Split(instruction, ":")
		cond := tab_instruct[0]
		if is_true(cond, rating) {
			direction := tab_instruct[1]
			if direction == "A" {
				return 1, "A"
			}
			if direction == "R" {
				return 2, "R"
			} else {
				return 0, direction
			}
		} else {
			return 3, ""
		}

	}
}

func is_true(condition string, rating map[string]int) bool {
	variable := string(condition[0])
	number := str_to_int(condition[2:])
	if condition[1] == '>' {
		return rating[variable] > number
	}
	if condition[1] == '<' {
		return rating[variable] < number
	}
	fmt.Println("Problème de condition")
	return false
}

func sum_rating(rating map[string]int) int {
	return rating["x"] + rating["s"] + rating["a"] + rating["m"]
}

func count_acceptance(ratings []map[string]int, workflows map[string][]string) int {
	sum := 0
	for _, rating := range ratings {
		if acceptance(rating, workflows) == "A" {
			sum += sum_rating(rating)
		}
	}
	return sum
}
