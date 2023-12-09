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
	file, err := os.Open("C:/Users/pelis/Documents/Mines2A/Programminglang/adventofcode2023/day7/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []string{}
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	var hands [][2]interface{}
	for _, hand := range res {
		hands_sep := strings.Split(hand, " ")
		hands = append(hands, [2]interface{}{hands_sep[0], str_to_int(hands_sep[1])})
	}

	fmt.Println(total_winnings(tabtotal(hands)))
	fmt.Println(total_winnings(tabtotaljoker(hands)))

}

func interf_str(interf [2]interface{}) string {
	if str, ok := interf[0].(string); ok {
		return str
	}
	fmt.Println("Erreur interface")
	return ""
}

func interf3_str(interf [3]interface{}) string {
	if str, ok := interf[0].(string); ok {
		return str
	}
	fmt.Println("Erreur interface")
	return ""
}

func interf_int(interf [2]interface{}) int {
	if nbr, ok := interf[1].(int); ok {
		return nbr
	}
	fmt.Println("Erreur interface")
	return 0
}

func interf3_int2(interf [3]interface{}) int {
	if nbr, ok := interf[2].(int); ok {
		return nbr
	}
	fmt.Println("Erreur interface")
	return 0
}

func interf3_int1(interf [3]interface{}) int {
	if nbr, ok := interf[1].(int); ok {
		return nbr
	}
	fmt.Println("Erreur interface")
	return 0
}

func str_to_int(intstring string) int {
	value, err := strconv.Atoi(intstring)
	if err != nil {
		fmt.Println("Erreur de conversion")
		return value
	}
	return value
}

// renvoie les cartes différentes qu'il y a dans une manche hand  (dans l'ordre)
func diff_card(manche string) []string {
	diff_card := []string{string(manche[0])}
	for i := 1; i < len(manche); i++ {
		count := 0
		for j := 0; j < len(diff_card); j++ {
			if string(manche[i]) != diff_card[j] {
				count++
			}
		}
		if count == len(diff_card) {
			diff_card = append(diff_card, string(manche[i]))
		}
	}
	return diff_card
}

// compte le nbr de fois ou card apparaît dans la manche manche
func count_card(card string, manche string) int {
	count := 0
	for i := 0; i < len(manche); i++ {
		if string(manche[i]) == card {
			count++
		}
	}
	return count
}

func is_better(card1 string, card2 string) string {
	if card1 == card2 {
		fmt.Println("Problème : égalité")
		return "equal"
	}
	tab := [13]string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}
	for _, val := range tab {
		if val == card1 {
			return "true"
		}
		if val == card2 {
			return "false"
		}
	}
	fmt.Println("Erreur de comparaison des cartes")
	return "false"
}

func is_fiveofakind(hand string) bool {
	return len(diff_card(hand)) == 1
}

func is_fourofakind(hand string) bool {
	if len(diff_card(hand)) == 2 {
		first := diff_card(hand)[0]
		second := diff_card(hand)[1]
		return (count_card(first, hand) == 1 && count_card(second, hand) == 4) || (count_card(first, hand) == 4 && count_card(second, hand) == 1)
	}
	return false
}

func is_fullhouse(hand string) bool {
	if len(diff_card(hand)) == 2 {
		first := diff_card(hand)[0]
		second := diff_card(hand)[1]
		return (count_card(first, hand) == 2 && count_card(second, hand) == 3) || (count_card(first, hand) == 3 && count_card(second, hand) == 2)
	}
	return false
}

func is_threeofakind(hand string) bool {
	if len(diff_card(hand)) == 3 {
		for i := 0; i < 3; i++ {
			if count_card(diff_card(hand)[i], hand) == 3 {
				return true
			}
		}
	}
	return false
}

func is_twopairs(hand string) bool {
	if len(diff_card(hand)) == 3 {
		count := 0
		for i := 0; i < 3; i++ {
			if count_card(diff_card(hand)[i], hand) == 2 {
				count++
			}
		}
		return count == 2
	}
	return false
}

func is_onepair(hand string) bool {
	return len(diff_card(hand)) == 4
}

func is_highcard(hand string) bool {
	return len(diff_card(hand)) == 5
}

func order(hands [][2]interface{}, offset int) [][3]interface{} {
	tab := [][3]interface{}{}
	for _, handf := range hands {
		main := interf_str(handf)
		place := place(main, hands)
		tab_handf := [3]interface{}{main, interf_int(handf), place + offset + 1}
		tab = append(tab, tab_handf)
	}
	taborder := [][3]interface{}{}
	for i := 0; i < len(tab); i++ {
		for _, hand := range tab {
			if interf3_int2(hand) == i+1+offset {
				taborder = append(taborder, hand)
			}
		}
	}
	return taborder
}

func place(main string, tab [][2]interface{}) int {
	place := 0
	ind := 0
	bool := false
	for i := 0; i < len(tab); i++ {
		if main != interf_str(tab[i]) {
			ind = 0
			bool = false
			for j := 0; j < len(main); j++ {
				if string(main[j]) == string(interf_str(tab[i])[j]) && !bool {
					ind++
				} else {
					bool = true
				}
			}
			if is_better(string(main[ind]), string(interf_str(tab[i])[ind])) == "true" {
				place++
			}
		}
	}
	return place
}

func make_tabfive(hands [][2]interface{}, offset int) [][3]interface{} {
	tab_five := [][2]interface{}{}
	for _, hand := range hands {
		main := interf_str(hand)
		mise := interf_int(hand)
		if is_fiveofakind(main) {
			tab_five = append(tab_five, [2]interface{}{main, mise})
		}
	}
	tab := order(tab_five, offset)
	return tab
}

func make_tab_four(hands [][2]interface{}, offset int) ([][3]interface{}, int) {
	tab_four := [][2]interface{}{}
	len := 0
	for _, hand := range hands {
		main := interf_str(hand)
		mise := interf_int(hand)
		if is_fourofakind(main) {
			tab := [2]interface{}{main, mise}
			tab_four = append(tab_four, tab)
			len++
		}
	}
	tab := order(tab_four, offset)
	return tab, len
}

func make_tabfull(hands [][2]interface{}, offset int) ([][3]interface{}, int) {
	tab_full := [][2]interface{}{}
	len := 0
	for _, hand := range hands {
		main := string(interf_str(hand))
		mise := interf_int(hand)
		if is_fullhouse(main) {
			tab_full = append(tab_full, [2]interface{}{main, mise})
			len++
		}
	}
	tab := order(tab_full, offset)
	return tab, len
}

func make_tabthree(hands [][2]interface{}, offset int) ([][3]interface{}, int) {
	tab_three := [][2]interface{}{}
	len := 0
	for _, hand := range hands {
		main := string(interf_str(hand))
		mise := interf_int(hand)
		if is_threeofakind(main) {
			tab_three = append(tab_three, [2]interface{}{main, mise})
			len++
		}
	}
	tab := order(tab_three, offset)
	return tab, len
}

func make_tabtwopairs(hands [][2]interface{}, offset int) ([][3]interface{}, int) {
	tab_twopairs := [][2]interface{}{}
	len := 0
	for _, hand := range hands {
		main := string(interf_str(hand))
		mise := interf_int(hand)
		if is_twopairs(main) {
			tab_twopairs = append(tab_twopairs, [2]interface{}{main, mise})
			len++
		}
	}
	tab := order(tab_twopairs, offset)
	return tab, len
}

func make_tabonepair(hands [][2]interface{}, offset int) ([][3]interface{}, int) {
	tab_onepair := [][2]interface{}{}
	len := 0
	for _, hand := range hands {
		main := string(interf_str(hand))
		mise := interf_int(hand)
		if is_onepair(main) {
			tab_onepair = append(tab_onepair, [2]interface{}{main, mise})
			len++
		}
	}
	tab := order(tab_onepair, offset)
	return tab, len
}

func make_tabhighcard(hands [][2]interface{}, offset int) ([][3]interface{}, int) {
	tab_highcard := [][2]interface{}{}
	len := 0
	for _, hand := range hands {
		main := string(interf_str(hand))
		mise := interf_int(hand)
		if is_highcard(main) {
			tab_highcard = append(tab_highcard, [2]interface{}{main, mise})
			len++
		}
	}
	tab := order(tab_highcard, offset)
	return tab, len
}

func tabtotal(hands [][2]interface{}) [][3]interface{} {
	total := [][3]interface{}{}
	offset := 0
	highcard, len_h := make_tabhighcard(hands, offset)
	for _, hand := range highcard {
		total = append(total, hand)
	}
	offset += len_h
	onepair, len_o := make_tabonepair(hands, offset)
	for _, hand := range onepair {
		total = append(total, hand)
	}
	offset += len_o
	twopairs, len_2 := make_tabtwopairs(hands, offset)
	for _, hand := range twopairs {
		total = append(total, hand)
	}
	offset += len_2
	three, len_t := make_tabthree(hands, offset)
	for _, hand := range three {
		total = append(total, hand)
	}
	offset += len_t
	full, len_f := make_tabfull(hands, offset)
	for _, hand := range full {
		total = append(total, hand)
	}
	offset += len_f
	four, len_4 := make_tab_four(hands, offset)
	for _, hand := range four {
		total = append(total, hand)
	}
	offset += len_4
	five := make_tabfive(hands, offset)
	for _, hand := range five {
		total = append(total, hand)
	}
	return total
}

func total_winnings(hands [][3]interface{}) int {
	total := 0
	for _, hand := range hands {
		total += interf3_int1(hand) * interf3_int2(hand)
	}
	return total
}

// PART 2

func is_betterwithjoker(card1 string, card2 string) string {
	if card1 == card2 {
		fmt.Println("Problème : égalité")
		return "equal"
	}
	tab := [13]string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}
	for _, val := range tab {
		if val == card1 {
			return "true"
		}
		if val == card2 {
			return "false"
		}
	}
	fmt.Println("Erreur de comparaison des cartes")
	return "false"
}

func orderjoker(hands [][2]interface{}, offset int) [][3]interface{} {
	tab := [][3]interface{}{}
	for _, handf := range hands {
		main := interf_str(handf)
		place := placejoker(main, hands)
		tab_handf := [3]interface{}{main, interf_int(handf), place + offset + 1}
		tab = append(tab, tab_handf)
	}
	taborder := [][3]interface{}{}
	for i := 0; i < len(tab); i++ {
		for _, hand := range tab {
			if interf3_int2(hand) == i+1+offset {
				taborder = append(taborder, hand)
			}
		}
	}
	return taborder
}

func placejoker(main string, tab [][2]interface{}) int {
	place := 0
	ind := 0
	bool := false
	for i := 0; i < len(tab); i++ {
		if main != interf_str(tab[i]) {
			ind = 0
			bool = false
			for j := 0; j < len(main); j++ {
				if string(main[j]) == string(interf_str(tab[i])[j]) && !bool {
					ind++
				} else {
					bool = true
				}
			}
			if is_betterwithjoker(string(main[ind]), string(interf_str(tab[i])[ind])) == "true" {
				place++
			}
		}
	}
	return place
}

func is_fiveofakindjoker(hand string) bool {
	if len(diff_card(hand)) == 1 && count_card("J", hand) == 0 {
		return true
	}
	if len(diff_card(hand)) == 2 {
		return count_card("J", hand) > 0
	}
	if count_card("J", hand) == 5 {
		return true
	}
	return false
}

func is_fourofakindjoker(hand string) bool {
	if len(diff_card(hand)) == 2 && count_card("J", hand) == 0 {
		first := diff_card(hand)[0]
		second := diff_card(hand)[1]
		return (count_card(first, hand) == 1 && count_card(second, hand) == 4) || (count_card(first, hand) == 4 && count_card(second, hand) == 1)
	}
	if len(diff_card(hand)) == 3 && count_card("J", hand) > 0 {
		diff := diff_card(hand)
		if count_card("J", hand) == 1 {
			for i := 0; i < 3; i++ {
				if count_card(diff[i], hand) == 3 {
					return true
				}
			}
		}
		if count_card("J", hand) == 2 {
			return (count_card(diff[0], hand) == 1 || count_card(diff[1], hand) == 1 || count_card(diff[2], hand) == 1)
		}
		if count_card("J", hand) == 3 {
			return true
		}
	}
	return false
}

func is_fullhousejoker(hand string) bool {
	if len(diff_card(hand)) == 2 && count_card("J", hand) == 0 {
		first := diff_card(hand)[0]
		second := diff_card(hand)[1]
		return (count_card(first, hand) == 2 && count_card(second, hand) == 3) || (count_card(first, hand) == 3 && count_card(second, hand) == 2)
	}
	if count_card("J", hand) == 1 {
		diff := diff_card(hand)
		if len(diff) == 3 {
			return (count_card(diff[0], hand) == 2 || count_card(diff[1], hand) == 2)
		}
	}

	return false
}

func is_threeofakindjoker(hand string) bool {
	if len(diff_card(hand)) == 3 && count_card("J", hand) == 0 {
		for i := 0; i < 3; i++ {
			if count_card(diff_card(hand)[i], hand) == 3 {
				return true
			}
		}
	}
	if count_card("J", hand) == 1 {
		return len(diff_card(hand)) == 4
	}
	if count_card("J", hand) == 2 {
		return len(diff_card(hand)) == 4
	}
	return false
}

func is_twopairsjoker(hand string) bool {
	if len(diff_card(hand)) == 3 && count_card("J", hand) == 0 {
		count := 0
		for i := 0; i < 3; i++ {
			if count_card(diff_card(hand)[i], hand) == 2 {
				count++
			}
		}
		return count == 2
	}
	return false
}

func is_onepairjoker(hand string) bool {
	if len(diff_card(hand)) == 4 && count_card("J", hand) == 0 {
		return true
	}
	if count_card("J", hand) == 1 {
		return len(diff_card(hand)) == 5
	}
	return false
}

func is_highcardjoker(hand string) bool {
	return len(diff_card(hand)) == 5 && count_card("J", hand) == 0
}

func make_tabfivejoker(hands [][2]interface{}, offset int) [][3]interface{} {
	tab_five := [][2]interface{}{}
	for _, hand := range hands {
		main := interf_str(hand)
		mise := interf_int(hand)
		if is_fiveofakindjoker(main) {
			tab_five = append(tab_five, [2]interface{}{main, mise})
		}
	}
	tab := orderjoker(tab_five, offset)
	return tab
}

func make_tab_fourjoker(hands [][2]interface{}, offset int) ([][3]interface{}, int) {
	tab_four := [][2]interface{}{}
	len := 0
	for _, hand := range hands {
		main := interf_str(hand)
		mise := interf_int(hand)
		if is_fourofakindjoker(main) {
			tab := [2]interface{}{main, mise}
			tab_four = append(tab_four, tab)
			len++
		}
	}
	tab := orderjoker(tab_four, offset)
	return tab, len
}

func make_tabfulljoker(hands [][2]interface{}, offset int) ([][3]interface{}, int) {
	tab_full := [][2]interface{}{}
	len := 0
	for _, hand := range hands {
		main := string(interf_str(hand))
		mise := interf_int(hand)
		if is_fullhousejoker(main) {
			tab_full = append(tab_full, [2]interface{}{main, mise})
			len++
		}
	}
	tab := orderjoker(tab_full, offset)
	return tab, len
}

func make_tabthreejoker(hands [][2]interface{}, offset int) ([][3]interface{}, int) {
	tab_three := [][2]interface{}{}
	len := 0
	for _, hand := range hands {
		main := string(interf_str(hand))
		mise := interf_int(hand)
		if is_threeofakindjoker(main) {
			tab_three = append(tab_three, [2]interface{}{main, mise})
			len++
		}
	}
	tab := orderjoker(tab_three, offset)
	return tab, len
}

func make_tabtwopairsjoker(hands [][2]interface{}, offset int) ([][3]interface{}, int) {
	tab_twopairs := [][2]interface{}{}
	len := 0
	for _, hand := range hands {
		main := string(interf_str(hand))
		mise := interf_int(hand)
		if is_twopairsjoker(main) {
			tab_twopairs = append(tab_twopairs, [2]interface{}{main, mise})
			len++
		}
	}
	tab := orderjoker(tab_twopairs, offset)
	return tab, len
}

func make_tabonepairjoker(hands [][2]interface{}, offset int) ([][3]interface{}, int) {
	tab_onepair := [][2]interface{}{}
	len := 0
	for _, hand := range hands {
		main := string(interf_str(hand))
		mise := interf_int(hand)
		if is_onepairjoker(main) {
			tab_onepair = append(tab_onepair, [2]interface{}{main, mise})
			len++
		}
	}
	tab := orderjoker(tab_onepair, offset)
	return tab, len
}

func make_tabhighcardjoker(hands [][2]interface{}, offset int) ([][3]interface{}, int) {
	tab_highcard := [][2]interface{}{}
	len := 0
	for _, hand := range hands {
		main := string(interf_str(hand))
		mise := interf_int(hand)
		if is_highcardjoker(main) {
			tab_highcard = append(tab_highcard, [2]interface{}{main, mise})
			len++
		}
	}
	tab := orderjoker(tab_highcard, offset)
	return tab, len
}

func tabtotaljoker(hands [][2]interface{}) [][3]interface{} {
	total := [][3]interface{}{}
	offset := 0
	highcard, len_h := make_tabhighcardjoker(hands, offset)
	for _, hand := range highcard {
		total = append(total, hand)
	}
	offset += len_h
	onepair, len_o := make_tabonepairjoker(hands, offset)
	for _, hand := range onepair {
		total = append(total, hand)
	}
	offset += len_o
	twopairs, len_2 := make_tabtwopairsjoker(hands, offset)
	for _, hand := range twopairs {
		total = append(total, hand)
	}
	offset += len_2
	three, len_t := make_tabthreejoker(hands, offset)
	for _, hand := range three {
		total = append(total, hand)
	}
	offset += len_t
	full, len_f := make_tabfulljoker(hands, offset)
	for _, hand := range full {
		total = append(total, hand)
	}
	offset += len_f
	four, len_4 := make_tab_fourjoker(hands, offset)
	for _, hand := range four {
		total = append(total, hand)
	}
	offset += len_4
	five := make_tabfivejoker(hands, offset)
	for _, hand := range five {
		total = append(total, hand)
	}
	return total
}
