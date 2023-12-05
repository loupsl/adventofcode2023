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

	file, err := os.Open("C:/Users/pelis/Documents/Mines2A/Programminglang/adventofcode2023/day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []string{}
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	games := []string{}
	for i := 0; i < len(res); i++ {
		game_i := strings.Split(res[i], ": ")[1]
		games = append(games, game_i)
	}
	games_sep := [][][]string{}
	for _, game := range games {
		manches := strings.Split(game, "; ")
		manches_colors := [][]string{}
		for _, manche := range manches {
			colors := strings.Split(manche, ", ")
			manches_colors = append(manches_colors, colors)
		}
		games_sep = append(games_sep, manches_colors)
	}

	fmt.Println(count_gamepossible(games_sep))
	fmt.Println(total_power(games_sep))
}

// renvoie le nbr de cubes tirés
func nbr_cubes(tirage string) int {
	spaceIndex := strings.Index(tirage, " ")
	charAsString := string(tirage[:spaceIndex])
	value, err := strconv.Atoi(charAsString)
	if err != nil {
		fmt.Println("Erreur de conversion en entier:", err)
		return 0
	}
	return value
}

// renvoie la couleur des boules tirées
func color(tirage string) string {
	spaceIndex := strings.Index(tirage, " ")
	return tirage[spaceIndex+1:]
}

// retourne true si le tirage est possible
func isPossible_tirage(tirage string) bool {
	if color(tirage) == "blue" {
		if nbr_cubes(tirage) < 15 {
			return true
		}
	}
	if color(tirage) == "green" {
		if nbr_cubes(tirage) < 14 {
			return true
		}
	}
	if color(tirage) == "red" {
		if nbr_cubes(tirage) < 13 {
			return true
		}
	}
	return false
}

func isPossible_manche(manche []string) bool {
	for _, tirage := range manche {
		if !isPossible_tirage(tirage) {
			return false
		}
	}
	return true
}

func isPossible_game(game [][]string) bool {
	for _, manche := range game {
		if !isPossible_manche(manche) {
			return false
		}
	}
	return true
}

func count_gamepossible(games [][][]string) int {
	count := 0
	for i := 0; i < len(games); i++ {
		if isPossible_game(games[i]) {
			count += i + 1
		}
	}
	return count
}

func cubes_needed(game [][]string, col string) int {
	max_cubes_color := 0
	for _, manche := range game {
		for _, tirage := range manche {
			if color(tirage) == col {
				if nbr_cubes(tirage) > max_cubes_color {
					max_cubes_color = nbr_cubes(tirage)
				}
			}
		}
	}
	return max_cubes_color
}

func power(game [][]string) int {
	pow := 0
	blue_cubes := cubes_needed(game, "blue")
	green_cubes := cubes_needed(game, "green")
	red_cubes := cubes_needed(game, "red")
	pow = blue_cubes * green_cubes * red_cubes
	return pow
}

func total_power(games [][][]string) int {
	count := 0
	for _, game := range games {
		count += power(game)
	}
	return count
}
