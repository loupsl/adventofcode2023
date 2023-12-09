package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("C:/Users/pelis/Documents/Mines2A/Programminglang/adventofcode2023/day3/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []string{}
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	//fmt.Println(make_posint(res))
	//fmt.Println(make_possymbol(res))
	//fmt.Println(is_adjacent(make_posint(res)[2], make_possymbol(res), res))
	fmt.Println(total_sum(res))

}

func str_to_int(intstring string) int {
	value, err := strconv.Atoi(intstring)
	if err != nil {
		fmt.Println("Erreur de conversion")
		return value
	}
	return value
}

func len_int(number int) int {
	if number > 999 {
		return 4
	}
	if number > 99 {
		return 3
	}
	if number > 9 {
		return 2
	}
	if number >= 0 {
		return 1
	}
	fmt.Println("longueur inconnue")
	return 0
}

func ascii_to_strint(ascii byte) string {
	if ascii == '0' {
		return "0"
	}
	if ascii == '1' {
		return "1"
	}
	if ascii == '2' {
		return "2"
	}
	if ascii == '3' {
		return "3"
	}
	if ascii == '4' {
		return "4"
	}
	if ascii == '5' {
		return "5"
	}
	if ascii == '6' {
		return "6"
	}
	if ascii == '7' {
		return "7"
	}
	if ascii == '8' {
		return "8"
	}
	if ascii == '9' {
		return "9"
	}
	return "pb de code ascii"
}

func make_posint(file []string) [][3]int {
	positions_int := [][3]int{}
	for i := 0; i < len(file); i++ {
		int_pos := [3]int{0, 0, 0}
		bool := false
		str_int := ""
		for j := 0; j < len(file[i]); j++ {
			ascii := file[i][j]
			if ascii >= '0' && ascii <= '9' {
				if bool { //on est au milieu d'1 int
					str_int += ascii_to_strint(ascii)
				}
				if !bool { //nv int trouvé
					bool = true
					int_pos[1] = i
					int_pos[2] = j
					str_int += ascii_to_strint(ascii)
				}
			} else {
				if bool { //on a fini de paser l'int
					int_pos[0] = str_to_int(str_int)
					positions_int = append(positions_int, int_pos)
					int_pos = [3]int{0, 0, 0}
					bool = false
					str_int = ""
				}
			}
		}
	}
	return positions_int
}

func make_possymbol(file []string) [][2]int {
	positions_symbol := [][2]int{}
	for i := 0; i < len(file); i++ {
		for j := 0; j < len(file[i]); j++ {
			ascii := file[i][j]
			if ascii != '.' && !(ascii >= '0' && ascii <= '9') {
				positions_symbol = append(positions_symbol, [2]int{i, j})
			}
		}
	}
	return positions_symbol
}

func is_adjacent(posint [3]int, possymbol [][2]int, res []string) bool {
	lenint := len_int(posint[0])
	ligne := posint[1]
	colonne := posint[2]
	indstart := -1
	indend := lenint
	if colonne == 0 {
		indstart = 0
	}
	if colonne == len(res[0])-lenint {
		indend = colonne + lenint - 1
	}
	for _, pos := range possymbol {
		if pos[0] == ligne {
			if colonne != 0 {
				if pos[1] == colonne-1 {
					return true
				}
			}
			if colonne != len(res[0])-lenint {
				if pos[1] == colonne+lenint {
					return true
				}
			}
		} else {
			if ligne != 0 {
				if pos[0] == ligne-1 {
					for i := indstart; i <= indend; i++ {
						if pos[1] == colonne+i {
							return true
						}
					}
				}
			}
			if ligne != len(res)-1 {
				if pos[0] == ligne+1 {
					for i := indstart; i <= indend; i++ {
						if pos[1] == colonne+i {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

func total_sum(file []string) int {
	totalsum := 0
	posint := make_posint(file)
	possymbol := make_possymbol(file)
	for _, tabint := range posint {
		if is_adjacent(tabint, possymbol, file) {
			totalsum += tabint[0]
		}
	}
	return totalsum
}
