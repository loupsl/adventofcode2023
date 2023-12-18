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
	file, err := os.Open("C:/Users/pelis/Documents/Mines2A/Programminglang/adventofcode2023/day18/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res := [][2]interface{}{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " (")[0]
		interf := [2]interface{}{string(line[0]), str_to_int(string(line[2:]))}
		res = append(res, interf)
	}

	path := final_tab(res)
	dig_tab := fill_tab(path)
	fmt.Println(count_lava(dig_tab))
}

func str_to_int(intstring string) int {
	value, err := strconv.Atoi(intstring)
	if err != nil {
		fmt.Println("Erreur de conversion")
		return value
	}
	return value
}

func get_direction(interf [2]interface{}) string {
	if str, ok := interf[0].(string); ok {
		return str
	}
	fmt.Println("Erreur interface")
	return ""
}

func get_size(interf [2]interface{}) int {
	if size, ok := interf[1].(int); ok {
		return size
	}
	fmt.Println("Erreur interface")
	return -1
}

func extand_tab_L(nbrcolumn int, tab []string, current_pos [2]int) ([]string, [2]int) {
	newtab := []string{}
	empty_column := ""
	marked_column := ""
	for i := 0; i < nbrcolumn; i++ {
		empty_column += "."
	}
	for i := 0; i < nbrcolumn; i++ {
		marked_column += "#"
	}
	for i := 0; i < len(tab); i++ {
		if i == current_pos[0] {
			newtab = append(newtab, marked_column+tab[i])
		} else {
			newtab = append(newtab, empty_column+tab[i])
		}
	}
	return newtab, [2]int{current_pos[0], 0}
}

func extand_tab_R(nbrcolumn int, tab []string, current_pos [2]int) ([]string, [2]int) {
	newtab := []string{}
	empty_column := ""
	marked_column := ""
	for i := 0; i < nbrcolumn; i++ {
		empty_column += "."
	}
	for i := 0; i < nbrcolumn; i++ {
		marked_column += "#"
	}
	for i := 0; i < len(tab); i++ {
		if i == current_pos[0] {
			newtab = append(newtab, tab[i]+marked_column)
		} else {
			newtab = append(newtab, tab[i]+empty_column)
		}
	}
	return newtab, [2]int{current_pos[0], current_pos[1] + nbrcolumn}
}

func extand_tab_U(nbrline int, tab []string, current_pos [2]int) ([]string, [2]int) {
	marked_line := ""
	for i := 0; i < len(tab[0]); i++ {
		if i == current_pos[1] {
			marked_line += "#"
		} else {
			marked_line += "."
		}
	}
	newtab := []string{}
	for i := 0; i < nbrline; i++ {
		newtab = append(newtab, marked_line)
	}
	for _, line := range tab {
		newtab = append(newtab, line)
	}
	return newtab, [2]int{0, current_pos[1]}
}

func extand_tab_D(nbrline int, tab []string, current_pos [2]int) ([]string, [2]int) {
	marked_line := ""
	for i := 0; i < len(tab[0]); i++ {
		if i == current_pos[1] {
			marked_line += "#"
		} else {
			marked_line += "."
		}
	}
	newtab := []string{}
	for _, line := range tab {
		newtab = append(newtab, line)
	}
	for i := 0; i < nbrline; i++ {
		newtab = append(newtab, marked_line)
	}
	return newtab, [2]int{current_pos[0] + nbrline, current_pos[1]}
}

func enough_space(size int, direction string, current_pos [2]int, tab []string) bool {
	i := current_pos[0]
	j := current_pos[1]
	if direction == "L" {
		return j-size >= 0
	}
	if direction == "R" {
		return j+size < len(tab[0])
	}
	if direction == "U" {
		return i-size >= 0
	}
	if direction == "D" {
		return i+size < len(tab)
	}
	fmt.Println("Probleme de direction")
	return false
}

func needed_space(size int, direction string, current_pos [2]int, tab []string) int {
	i := current_pos[0]
	j := current_pos[1]
	if direction == "L" {
		return size - j
	}
	if direction == "R" {
		return j + size - len(tab[0]) + 1
	}
	if direction == "U" {
		return size - i
	}
	if direction == "D" {
		return i + size - len(tab) + 1
	}
	fmt.Println("Probleme de direction")
	return 0
}

func update_tabL(size int, current_pos [2]int, tab []string) ([]string, [2]int) {
	i := current_pos[0]
	j := current_pos[1]
	new_line := ""
	for ind := 0; ind < len(tab[i]); ind++ {
		if ind >= j-size && ind <= j {
			new_line += "#"
		} else {
			new_line += string(tab[i][ind])
		}
	}
	newtab := []string{}
	for ind := 0; ind < len(tab); ind++ {
		if ind == i {
			newtab = append(newtab, new_line)
		} else {
			newtab = append(newtab, tab[ind])
		}
	}
	return newtab, [2]int{i, j - size}
}

func update_tabR(size int, current_pos [2]int, tab []string) ([]string, [2]int) {
	i := current_pos[0]
	j := current_pos[1]
	new_line := ""
	for ind := 0; ind < len(tab[i]); ind++ {
		if ind >= j && ind <= j+size {
			new_line += "#"
		} else {
			new_line += string(tab[i][ind])
		}
	}
	newtab := []string{}
	for ind := 0; ind < len(tab); ind++ {
		if ind == i {
			newtab = append(newtab, new_line)
		} else {
			newtab = append(newtab, tab[ind])
		}
	}
	return newtab, [2]int{i, j + size}
}

func update_tabU(size int, current_pos [2]int, tab []string) ([]string, [2]int) {
	i := current_pos[0]
	j := current_pos[1]
	newtab := []string{}
	ind := 0
	for _, line := range tab {
		if ind >= i-size && ind < i {
			newline := ""
			for u := 0; u < len(line); u++ {
				if u == j {
					newline += "#"
				} else {
					newline += string(line[u])
				}
			}
			newtab = append(newtab, newline)
		} else {
			newtab = append(newtab, line)
		}
		ind++
	}
	return newtab, [2]int{i - size, j}
}

func update_tabD(size int, current_pos [2]int, tab []string) ([]string, [2]int) {
	i := current_pos[0]
	j := current_pos[1]
	newtab := []string{}
	ind := 0
	for _, line := range tab {
		if ind > i && ind <= i+size {
			newline := ""
			for u := 0; u < len(line); u++ {
				if u == j {
					newline += "#"
				} else {
					newline += string(line[u])
				}
			}
			newtab = append(newtab, newline)
		} else {
			newtab = append(newtab, line)
		}
		ind++
	}
	return newtab, [2]int{i + size, j}
}

func make_instruction(instruction [2]interface{}, current_pos [2]int, tab []string) ([]string, [2]int) {
	size := get_size(instruction)
	direction := get_direction(instruction)
	if enough_space(size, direction, current_pos, tab) {
		if direction == "L" {
			return update_tabL(size, current_pos, tab)
		}
		if direction == "R" {
			return update_tabR(size, current_pos, tab)
		}
		if direction == "U" {
			return update_tabU(size, current_pos, tab)
		}
		if direction == "D" {
			return update_tabD(size, current_pos, tab)
		}
	} else {
		neededspace := needed_space(size, direction, current_pos, tab)
		if direction == "L" {
			newtab, newpos := update_tabL(size-neededspace, current_pos, tab)
			return extand_tab_L(neededspace, newtab, newpos)
		}
		if direction == "R" {
			newtab, newpos := update_tabR(size-neededspace, current_pos, tab)
			return extand_tab_R(neededspace, newtab, newpos)
		}
		if direction == "U" {
			newtab, newpos := update_tabU(size-neededspace, current_pos, tab)
			return extand_tab_U(neededspace, newtab, newpos)
		}
		if direction == "D" {
			newtab, newpos := update_tabD(size-neededspace, current_pos, tab)
			return extand_tab_D(neededspace, newtab, newpos)
		}
	}
	fmt.Println("Probleme make instructions")
	return tab, [2]int{}
}

func final_tab(input [][2]interface{}) []string {
	tab := []string{"#"}
	pos := [2]int{0, 0}
	for _, instruction := range input {
		tab, pos = make_instruction(instruction, pos, tab)
	}
	return tab
}

func fill_rec(bytestab [][]byte, i int, j int, width int, height int) {
	if i < 0 || i >= width || j < 0 || j >= height || bytestab[j][i] != '.' {
		return
	}
	bytestab[j][i] = '#'
	fill_rec(bytestab, i+1, j, width, height)
	fill_rec(bytestab, i-1, j, width, height)
	fill_rec(bytestab, i, j+1, width, height)
	fill_rec(bytestab, i, j-1, width, height)
}

func mark_borders(bytestab [][]byte, i int, y int, width int, height int) {
	if i < 0 || i >= width || y < 0 || y >= height || bytestab[y][i] != '.' {
		return
	}
	bytestab[y][i] = '*'
	mark_borders(bytestab, i+1, y, width, height)
	mark_borders(bytestab, i-1, y, width, height)
	mark_borders(bytestab, i, y+1, width, height)
	mark_borders(bytestab, i, y-1, width, height)
}

func fill_tab(tab []string) []string {
	bytestab := [][]byte{}
	for i := 0; i < len(tab); i++ {
		bytestab = append(bytestab, []byte(tab[i]))
	}
	for i := 0; i < len(tab[0]); i++ {
		mark_borders(bytestab, i, 0, len(tab[0]), len(tab))
		mark_borders(bytestab, i, len(tab)-1, len(tab[0]), len(tab))
	}
	for j := 0; j < len(tab); j++ {
		mark_borders(bytestab, 0, j, len(tab[0]), len(tab))
		mark_borders(bytestab, len(tab[0])-1, j, len(tab[0]), len(tab))
	}

	for j := 1; j < len(tab)-1; j++ {
		for i := 1; i < len(tab[0])-1; i++ {
			fill_rec(bytestab, i, j, len(tab[0]), len(tab))
		}
	}

	filledtab := []string{}
	for i := 0; i < len(tab); i++ {
		filledtab = append(filledtab, strings.ReplaceAll(string(bytestab[i]), "*", "."))
	}
	return filledtab
}

func count_lava(tab []string) int {
	lava := 0
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab[i]); j++ {
			if tab[i][j] == '#' {
				lava++
			}
		}
	}
	return lava
}
