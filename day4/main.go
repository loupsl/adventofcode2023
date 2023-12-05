package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("C:/Users/pelis/Documents/Mines2A/Programminglang/adventofcode2023/day4/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []string{}
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	new_res := []string{}
	for _, ligne := range res {
		newligne := delete_doublespace(ligne)
		new_res = append(new_res, newligne)
	}

	cards := [][2][]int{}

	for _, ligne := range new_res {
		list_card := [2][]int{}
		entire_card := strings.Split((strings.Split(ligne, ": ")[1]), " | ")
		win_card := strings.Split(entire_card[0], " ")
		my_card := strings.Split(entire_card[1], " ")
		list_card[0] = convert_int(win_card)
		list_card[1] = convert_int(my_card)
		cards = append(cards, list_card)
	}

	fmt.Println(total_points(cards))
	fmt.Println(total_scratchcards(cards))

}

func delete_doublespace(card string) string {
	new_card := ""
	for i := 0; i < len(card); i++ {
		new_card += string(card[i])
		if i < len(card)-1 {
			if card[i:i+2] == "  " {
				i++
			}
		}
	}
	return new_card
}

func convert_int(card []string) []int {
	cardint := []int{}
	for _, val := range card {
		value, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Erreur de conversion")
			return cardint
		}
		cardint = append(cardint, value)
	}
	return cardint
}

func is_int_in(number int, card []int) bool {
	for _, val := range card {
		if val == number {
			return true
		}
	}
	return false
}

func point_per_card(card [2][]int) int {
	points := 0
	for _, val := range card[1] {
		if is_int_in(val, card[0]) {
			points += 1
		}
	}
	return int(math.Pow(2, float64(points-1)))
}

func total_points(cards [][2][]int) int {
	total := 0
	for _, card := range cards {
		total += point_per_card(card)
	}
	return total
}

func match_per_card(card [2][]int) int {
	matches := 0
	for _, val := range card[1] {
		if is_int_in(val, card[0]) {
			matches += 1
		}
	}
	return matches
}

func sum(liste [198]int) int {
	somme := 0
	for _, val := range liste {
		somme += val
	}
	return somme
}

func total_scratchcards(cards [][2][]int) int {
	compteur := [198]int{}
	for i := 0; i < len(compteur); i++ {
		compteur[i] = 1
	}
	ind := 0
	for _, card := range cards {
		matches := match_per_card(card)
		for i := 1; i < matches+1; i++ {
			compteur[ind+i] += compteur[ind]
		}
		ind++
	}
	return (sum(compteur))
}
