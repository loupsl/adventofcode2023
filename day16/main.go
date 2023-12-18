package main

import (
	"bufio"
	_ "embed"
	"os"
)

func main() {
	file, err := os.Open("C:/Users/pelis/Documents/Mines2A/Programminglang/adventofcode2023/day16/input_test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	//fmt.Println(left_path(res,))
}

/* func new_direction(ascii int, direction int) int {
	if ascii == 92 { // "\""
		if direction == 0 {
			return 3
		}
		if direction == 1 {
			return 2
		}
		if direction == 2 {
			return 1
		}
		if direction == 3 {
			return 0
		}
	}
	if ascii == 47 { // "/"
		if direction == 0 {
			return 2
		}
		if direction == 1 {
			return 3
		}
		if direction == 2 {
			return 0
		}
		if direction == 3 {
			return 1
		}
	}
	if ascii == 45 { // "-"
		if direction == 0 {
			return 0
		}
		if direction == 1 {
			return 1
		}
		if direction == 2 {
			return 4
		}
		if direction == 3 {
			return 4
		}
	}
	if ascii == 124 { // "|"
		if direction == 0 {
			return 5
		}
		if direction == 1 {
			return 5
		}
		if direction == 2 {
			return 2
		}
		if direction == 3 {
			return 3
		}
	}
	fmt.Println("Erreur direction")
	return -1
}

func is_already_marked(i int, j int, tabmarked [][2]int, res []string) bool {
	for _, mark := range tabmarked {
		if i == mark[0] && j == mark[1] {
			return true
		}
	}
	return false
}

func left_path(i int, j int, res []string, tabmarked [][2]int) ([][2]int, int, int, int) {
	ligne := i
	colonne := j + 1
	direction := 0
	newtab := [][2]int{}
	for _, mark := range tabmarked {
		newtab = append(newtab, mark)
	}
	if colonne < len(res[ligne]) {
		for res[ligne][colonne] == '.' {
			if !is_already_marked(ligne, colonne, tabmarked, res) {
				newtab = append(newtab, [2]int{ligne, colonne})
			}
			colonne++
		}
	}
	if colonne == len(res[ligne]) {
		return newtab, -1, ligne, colonne
	}
	if res[ligne][colonne] != '.' {
		direction = new_direction(int(res[ligne][colonne]), direction)
		return newtab, direction, ligne, colonne
	}
	fmt.Println("Probleme left path")
	return [][2]int{}, -1, 0, 0
}

func right_path(i int, j int, res []string, tabmarked [][2]int) ([][2]int, int, int, int) {
	ligne := i
	colonne := j - 1
	direction := 1
	newtab := [][2]int{}
	for _, mark := range tabmarked {
		newtab = append(newtab, mark)
	}
	if colonne >= 0 {
		for res[ligne][colonne] == '.' {
			if !is_already_marked(ligne, colonne, tabmarked, res) {
				newtab = append(newtab, [2]int{ligne, colonne})
			}
			colonne--
		}
	}
	if colonne == -1 {
		return newtab, -1, ligne, colonne
	}
	if res[ligne][colonne] != '.' {
		direction = new_direction(int(res[ligne][colonne]), direction)
		return newtab, direction, ligne, colonne
	}
	fmt.Println("Probleme right path")
	return [][2]int{}, -1, 0, 0
}

func up_path(i int, j int, res []string, tabmarked [][2]int) ([][2]int, int, int, int) {
	ligne := i - 1
	colonne := j
	direction := 2
	newtab := [][2]int{}
	for _, mark := range tabmarked {
		newtab = append(newtab, mark)
	}
	if ligne >= 0 {
		for res[ligne][colonne] == '.' {
			if !is_already_marked(ligne, colonne, tabmarked, res) {
				newtab = append(newtab, [2]int{ligne, colonne})
			}
			ligne--
		}
	}
	if ligne == -1 {
		return newtab, -1, ligne, colonne
	}
	if res[ligne][colonne] != '.' {
		direction = new_direction(int(res[ligne][colonne]), direction)
		return newtab, direction, ligne, colonne
	}
	fmt.Println("Probleme left path")
	return [][2]int{}, -1, 0, 0
}

func down_path(i int, j int, res []string, tabmarked [][2]int) ([][2]int, int, int, int) {
	ligne := i + 1
	colonne := j
	direction := 3
	newtab := [][2]int{}
	for _, mark := range tabmarked {
		newtab = append(newtab, mark)
	}
	if ligne < len(res) {
		for res[ligne][colonne] == '.' {
			if !is_already_marked(ligne, colonne, tabmarked, res) {
				newtab = append(newtab, [2]int{ligne, colonne})
			}
			ligne++
		}
	}
	if ligne == len(res) {
		return newtab, -1, ligne, colonne
	}
	if res[ligne][colonne] != '.' {
		direction = new_direction(int(res[ligne][colonne]), direction)
		return newtab, direction, ligne, colonne
	}
	fmt.Println("Probleme left path")
	return [][2]int{}, -1, ligne, colonne
}

func path(res []string) int {
	ligne := 0
	colonne := 0
	tabmarked := [][2]int{}
	tab, direction, i, j := left_path(ligne, colonne, res, tabmarked)
	for direction != -1 {
		if direction == 0 {
			tab, direction, i, j = left_path(i, j, res, tabmarked)
		}
		if direction == 1 {
			tab, direction, i, j = right_path(i, j, res, tab)
		}
		if direction == 2 {
			tab, direction, i, j = up_path(i, j, res, tab)
		}
		if direction == 3 {
			tab, direction, i, j = down_path(i, j, res, tab)
		}
		if direction == 4 {
			direction = 0
			for direction != -1 {
				if direction == 0 {
					tab, direction, i, j = left_path(i, j, res, tabmarked)
				}
				if direction == 1 {
					tab, direction, i, j = right_path(i, j, res, tab)
				}
				if direction == 2 {
					tab, direction, i, j = up_path(i, j, res, tab)
				}
				if direction == 3 {
					tab, direction, i, j = down_path(i, j, res, tab)
				}
			}
			direction = 1
			for direction != -1 {
				if direction == 0 {
					tab, direction, i, j = left_path(i, j, res, tabmarked)
				}
				if direction == 1 {
					tab, direction, i, j = right_path(i, j, res, tab)
				}
				if direction == 2 {
					tab, direction, i, j = up_path(i, j, res, tab)
				}
				if direction == 3 {
					tab, direction, i, j = down_path(i, j, res, tab)
				}
			}
		}
		if direction == 5 {
			direction = 2
			for direction != -1 {
				if direction == 0 {
					tab, direction, i, j = left_path(i, j, res, tabmarked)
				}
				if direction == 1 {
					tab, direction, i, j = right_path(i, j, res, tab)
				}
				if direction == 2 {
					tab, direction, i, j = up_path(i, j, res, tab)
				}
				if direction == 3 {
					tab, direction, i, j = down_path(i, j, res, tab)
				}
			}
			direction = 3
			for direction != -1 {
				if direction == 0 {
					tab, direction, i, j = left_path(i, j, res, tabmarked)
				}
				if direction == 1 {
					tab, direction, i, j = right_path(i, j, res, tab)
				}
				if direction == 2 {
					tab, direction, i, j = up_path(i, j, res, tab)
				}
				if direction == 3 {
					tab, direction, i, j = down_path(i, j, res, tab)
				}
			}
		}
	}
	return len(tab)
}
*/
