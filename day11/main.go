package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("C:/Users/pelis/Documents/Mines2A/Programminglang/adventofcode2023/day11/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []string{}
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	pos_galaxies, count_ligne, count_colonne := make_pos_galaxies(res)
	fmt.Println(sum_shortestpath(expand_galaxies(pos_galaxies, count_ligne, count_colonne, res)))

}

func make_pos_galaxies(file []string) ([][2]int, []int, []int) {
	positions_galaxies := [][2]int{}
	count_colonne := []int{}
	count_ligne := []int{}
	for i := 0; i < len(file); i++ {
		count_ligne = append(count_ligne, 0)
	}
	for j := 0; j < len(file[0]); j++ {
		count_colonne = append(count_colonne, 0)
	}
	for i := 0; i < len(file); i++ {
		for j := 0; j < len(file[i]); j++ {
			ascii := file[i][j]
			if ascii == '#' {
				positions_galaxies = append(positions_galaxies, [2]int{i, j})
				count_ligne[i]++
				count_colonne[j]++
			}
		}
	}
	return positions_galaxies, count_ligne, count_colonne
}

func expand_galaxies(pos_galaxies [][2]int, count_ligne []int, count_colonne []int, file []string) [][2]int {
	tab := [][2]int{}
	for i := 0; i < len(pos_galaxies); i++ {
		tab = append(tab, pos_galaxies[i])
	}
	for ligne := 0; ligne < len(file); ligne++ {
		if count_ligne[ligne] == 0 {
			ind := 0
			for _, tabgalaxie := range pos_galaxies {
				if tabgalaxie[0] > ligne {
					tab[ind][0]++
				}
				ind++
			}
		}
	}
	for colonne := 0; colonne < len(file[0]); colonne++ {
		if count_colonne[colonne] == 0 {
			ind := 0
			for _, tabgalaxie := range pos_galaxies {
				if tabgalaxie[1] > colonne {
					tab[ind][1]++
				}
				ind++
			}
		}
	}
	return tab
}

func abs(nombre int) int {
	if nombre < 0 {
		return -1 * nombre
	} else {
		return nombre
	}
}

func shortestpath(pos1 [2]int, pos2 [2]int) int {
	return abs(pos2[0]-pos1[0]) + abs(pos2[1]-pos1[1])
}

func sum_shortestpath(posgalaxies [][2]int) int {
	sum := 0
	for i := 0; i < len(posgalaxies)-1; i++ {
		for j := i + 1; j < len(posgalaxies); j++ {
			sum += shortestpath(posgalaxies[i], posgalaxies[j])
		}
	}
	return sum
}
