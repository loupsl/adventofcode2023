package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("C:/Users/pelis/Documents/Mines2A/Programminglang/adventofcode2023/day13/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	res := [][]string{}

	group := []string{}
	ind := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			group = append(group, line)
			ind++
			if ind == 1355 {
				res = append(res, group)
			}
		} else {
			res = append(res, group)
			group = []string{}
			ind++
		}
	}

	fmt.Println(sum_mirrors(res))

}

func ind_mirror_ver(group []string) int {
	for ind := 1; ind < len(group[0]); ind++ {
		if is_mirror_vert(ind, group) {
			return ind
		}
	}
	return 0
}

func is_mirror_vert(center int, group []string) bool {
	length := min(len(group[0])-center, center)
	for _, ligne := range group {
		for i := 1; i <= length; i++ {
			if ligne[center-i] != ligne[center-1+i] {
				return false
			}
		}
	}
	return true
}

func ind_mirror_hor(group []string) int {
	for ind := 1; ind < len(group); ind++ {
		if is_mirror_hor(ind, group) {
			return ind
		}
	}
	return 0
}

func is_mirror_hor(center int, group []string) bool {
	heigth := min(len(group)-center, center)
	for i := 1; i <= heigth; i++ {
		for ind := 0; ind < len(group[0]); ind++ {
			if group[center-i][ind] != group[center-1+i][ind] {
				return false
			}
		}
	}
	return true
}

func sum_mirrors(groups [][]string) int {
	sum := 0
	for _, group := range groups {
		sum += ind_mirror_ver(group)
		sum += 100 * ind_mirror_hor(group)
	}
	return sum
}

//PART2

func locate_smudge(group []string) [2]int  {
	newgroup := []string{}
	for _,line := range group {
		newgroup = append(newgroup, line)
	}
	for j:=0;j<len(group);j++ {
		for i:=0;i<len(line);i++{
			if group[j][i] == '#' {
				newgroup[j][i] = '.'
				if ind_mirror_hor(group) != 0 {
					if ind_mirror_hor(newgroup) != 
				}
			}
			else 
		}
	}

}

func make_newtab(groupes [][]string) {}
