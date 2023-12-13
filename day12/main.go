package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("C:/Users/pelis/Documents/Mines2A/Programminglang/adventofcode2023/day12/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []string{}
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	tab := make_tab(res)
	fmt.Println(total_poss_arr(tab))

}

func make_tab(file []string) [][2]interface{} {
	tab := [][2]interface{}{}
	for i := 0; i < len(file); i++ {
		chain_i := ""
		tabint_i := []int{}
		for j := 0; j < len(file[i]); j++ {
			ascii := file[i][j]
			if ascii == '#' || ascii == '.' || ascii == '?' {
				if ascii == '#' {
					chain_i += "#"
				}
				if ascii == '.' {
					chain_i += "."
				}
				if ascii == '?' {
					chain_i += "?"
				}
			}
			if ascii > '0' && ascii <= '9' {
				tabint_i = append(tabint_i, ascii_to_int(ascii))
			}
		}
		tab = append(tab, [2]interface{}{chain_i, tabint_i})
	}
	return tab
}

func ascii_to_int(ascii byte) int {
	if ascii == '1' {
		return 1
	}
	if ascii == '2' {
		return 2
	}
	if ascii == '3' {
		return 3
	}
	if ascii == '4' {
		return 4
	}
	if ascii == '5' {
		return 5
	}
	if ascii == '6' {
		return 6
	}
	if ascii == '7' {
		return 7
	}
	if ascii == '8' {
		return 8
	}
	if ascii == '9' {
		return 9
	}
	fmt.Println("Problème de code ASCII")
	return 0
}

func get_chain(tab [2]interface{}) string {
	if chain, ok := tab[0].(string); ok {
		return chain
	}
	fmt.Println("Problème")
	return ""
}

func get_tabint(tabligne [2]interface{}) []int {
	if tabint, ok := tabligne[1].([]int); ok {
		return tabint
	}
	fmt.Println("problème")
	return []int{0}
}

func is_possible(arrangement string, tabligne [2]interface{}) bool {
	if len(arrangement) == len(get_chain(tabligne)) {
		count := 0
		tabint := []int{}
		bool := false
		for i := 0; i < len(arrangement); i++ {
			if arrangement[i] == '#' {
				bool = true
				count++
				if i == len(arrangement)-1 {
					tabint = append(tabint, count)
				}
			}
			if arrangement[i] == '.' && bool {
				bool = false
				tabint = append(tabint, count)
				count = 0
			}

		}
		return is_equal(get_tabint(tabligne), tabint)
	}
	return false
}

func is_equal(tab1 []int, tab2 []int) bool {
	if len(tab1) == len(tab2) {
		for i := 0; i < len(tab1); i++ {
			if tab1[i] != tab2[i] {
				return false
			}
		}
		return true
	}
	return false
}

func h_to_place(tabligne [2]interface{}) int {
	hash_to_place := 0
	tabint := get_tabint(tabligne)
	for i := 0; i < len(tabint); i++ {
		hash_to_place += tabint[i]
	}
	chain := get_chain(tabligne)
	for i := 0; i < len(chain); i++ {
		if chain[i] == '#' {
			hash_to_place--
		}
	}
	return hash_to_place
}

func place_h(chain string, index int, groupSizes []int, currentCombination *strings.Builder, combinations *[]string) {
	if index == len(chain) {
		if len(groupSizes) == 0 {
			*combinations = append(*combinations, currentCombination.String())
		}
		return
	}

	if chain[index] == '?' {
		if len(groupSizes) > 0 && groupSizes[0] > 0 {
			newGroupSizes := make([]int, len(groupSizes))
			copy(newGroupSizes, groupSizes)
			newGroupSizes[0]--

			currentCombination.WriteByte('#')
			place_h(chain, index+1, newGroupSizes, currentCombination, combinations)
			removeLastChar(currentCombination)
		}

		newGroupSizes := make([]int, len(groupSizes))
		copy(newGroupSizes, groupSizes)
		if len(newGroupSizes) > 0 && newGroupSizes[0] == 0 {
			newGroupSizes = newGroupSizes[1:]
		}

		currentCombination.WriteByte('.')
		place_h(chain, index+1, newGroupSizes, currentCombination, combinations)
		removeLastChar(currentCombination)

	} else {
		currentCombination.WriteByte(chain[index])
		if chain[index] == '#' && len(groupSizes) > 0 {
			groupSizes[0]--
			if groupSizes[0] == 0 {
				groupSizes = groupSizes[1:]
			}
		}
		place_h(chain, index+1, groupSizes, currentCombination, combinations)
	}
}

func removeLastChar(builder *strings.Builder) {
	str := builder.String()
	if len(str) > 0 {
		builder.Reset()
		builder.WriteString(str[:len(str)-1])
	}
}

func truncateBuilder(builder *strings.Builder) {
	str := builder.String()
	if len(str) > 0 {
		builder.Reset()
		builder.WriteString(str[:len(str)-1])
	}
}

func tab_possible_arr(tabligne [2]interface{}) []string {
	chain := get_chain(tabligne)
	tabint := get_tabint(tabligne)
	//n := h_to_place(tabligne)
	var combinations []string
	currentCombination := strings.Builder{}
	place_h(chain, 0, tabint, &currentCombination, &combinations)
	return combinations
}

func count_possible_arr(tabligne [2]interface{}) int {
	sum := 0
	tab_poss := tab_possible_arr(tabligne)
	for _, poss := range tab_poss {
		if is_possible(poss, tabligne) {
			sum++
		}
	}
	return sum
}

func total_poss_arr(tab [][2]interface{}) int {
	sum := 0
	for _, tabligne := range tab {
		sum += count_possible_arr(tabligne)
	}
	return sum
}
