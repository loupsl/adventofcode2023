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
	file, err := os.Open("C:/Users/pelis/Documents/Mines2A/Programminglang/adventofcode2023/day9/input.txt")
	if err != nil {
		panic(err)
	}

	var tab [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		entiers := convert_line_int(line)
		tab = append(tab, entiers)
	}

	fmt.Println(sum_total(tab))
}

func convert_line_int(line string) []int {
	mots := strings.Fields(line)
	var entiers []int
	for _, mot := range mots {
		entier, err := strconv.Atoi(mot)
		if err != nil {
			fmt.Println("Erreur lors de la conversion de l'entier:", err)
			return nil
		}
		entiers = append(entiers, entier)
	}

	return entiers
}

func make_linediff(ligne []int) []int {
	ligne_diff := []int{}
	for i := 0; i < len(ligne)-1; i++ {
		diff_i := ligne[i+1] - ligne[i]
		ligne_diff = append(ligne_diff, diff_i)
	}
	return ligne_diff
}

func is_zero(ligne []int) bool {
	for i := 0; i < len(ligne); i++ {
		if ligne[i] != 0 {
			return false
		}
	}
	return true
}

func make_tab_ligne(ligne []int) [][]int {
	tabligne := [][]int{}
	tabligne = append(tabligne, ligne)
	current_line := ligne
	for !is_zero(current_line) {
		next_line := make_linediff(current_line)
		tabligne = append(tabligne, next_line)
		current_line = next_line
	}
	return tabligne
}

func sum_ligne(ligne []int) int {
	tabline := make_tab_ligne(ligne)
	sum := 0
	for _, line := range tabline {
		sum += line[len(line)-1]
	}
	return sum
}

func sum_total(tab [][]int) int {
	sumtot := 0
	for _, line := range tab {
		sumtot += sum_ligne(line)
	}
	return sumtot
}
