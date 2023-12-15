package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("C:/Users/pelis/Documents/Mines2A/Programminglang/adventofcode2023/day15/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	strings := strings.Split(res[0], ",")

	fmt.Println(total_result(strings))

	fmt.Println(lens_totalpower(final_boxes(strings)))

}

func hash_algorithm(ascii int, currentvalue int) int {
	value := currentvalue + ascii
	value = value * 17
	value = value % 256
	return value
}

func hash_chain(chain string) int {
	value_chain := 0
	for _, char := range chain {
		val := hash_algorithm(int(char), value_chain)
		value_chain = val
	}
	return value_chain
}

func total_result(strings []string) int {
	sum := 0
	for _, chain := range strings {
		sum += hash_chain(chain)
	}
	return sum
}

// PART2

func ascii_to_int(ascii byte) int {
	if ascii == '9' {
		return 9
	}
	if ascii == '8' {
		return 8
	}
	if ascii == '7' {
		return 7
	}
	if ascii == '6' {
		return 6
	}
	if ascii == '5' {
		return 5
	}
	if ascii == '4' {
		return 4
	}
	if ascii == '3' {
		return 3
	}
	if ascii == '2' {
		return 2
	}
	if ascii == '1' {
		return 1
	}
	return 0
}

func get_label(lens [2]interface{}) string {
	if str, ok := lens[0].(string); ok {
		return str
	}
	fmt.Println("Erreur interface")
	return ""
}

func get_focallength(lens [2]interface{}) int {
	if focal_length, ok := lens[1].(int); ok {
		return focal_length
	}
	fmt.Println("Erreur interface")
	return 0
}

func read_lens(chain string) ([2]interface{}, string, int) {
	currentchar := 0
	str := ""
	for chain[currentchar] != '-' && chain[currentchar] != '=' {
		str += string(chain[currentchar])
		currentchar++
	}
	hash := hash_chain(str)
	if chain[currentchar] == '=' {
		focal_length := ascii_to_int(chain[currentchar+1])
		lens := [2]interface{}{str, focal_length}
		return lens, string(chain[currentchar]), hash
	}
	if chain[currentchar] == '-' {
		lens := [2]interface{}{str, 0}
		return lens, string(chain[currentchar]), hash
	}
	fmt.Println("Probleme read lens")
	return [2]interface{}{"", 0}, "", 0
}

func next_step(lens [2]interface{}, instruction string, ind_box int, boxes [256][][2]interface{}) [][2]interface{} {
	label := get_label(lens)
	new_box := [][2]interface{}{}
	ind_samelabel := is_in_the_box(label, boxes[ind_box])
	if instruction == "=" {
		if ind_samelabel == -1 {
			for _, lensparc := range boxes[ind_box] {
				new_box = append(new_box, lensparc)
			}
			new_box = append(new_box, lens)
			return new_box
		} else {
			for i := 0; i < len(boxes[ind_box]); i++ {
				if i == ind_samelabel {
					new_box = append(new_box, lens)
				} else {
					new_box = append(new_box, boxes[ind_box][i])
				}
			}
			return new_box
		}
	} else {
		if ind_samelabel != -1 {
			for i := 0; i < len(boxes[ind_box]); i++ {
				if i != ind_samelabel {
					new_box = append(new_box, boxes[ind_box][i])
				}
			}
			return new_box
		}
		return boxes[ind_box]
	}
}

func next_step_boxes(chain string, boxes [256][][2]interface{}) [256][][2]interface{} {
	lens, instruction, ind_box := read_lens(chain)
	finalboxes := [256][][2]interface{}{}
	for i := 0; i < len(boxes); i++ {
		if i == ind_box {
			finalboxes[i] = next_step(lens, instruction, i, boxes)
		} else {
			finalboxes[i] = boxes[i]
		}
	}
	return finalboxes
}

func is_in_the_box(label string, box [][2]interface{}) int {
	ind := 0
	for _, lens := range box {
		if get_label(lens) == label {
			return ind
		}
		ind++
	}
	return -1
}

func final_boxes(strings []string) [256][][2]interface{} {
	final_boxes := [256][][2]interface{}{}
	for _, str := range strings {
		final_boxes = next_step_boxes(str, final_boxes)
	}
	return final_boxes
}

func lens_totalpower(boxes [256][][2]interface{}) int {
	sum := 0
	ind_box := 1
	for _, box := range boxes {
		ind_lens := 1
		for _, lens := range box {
			sum += ind_box * ind_lens * get_focallength(lens)
			ind_lens++
		}
		ind_box++
	}
	return sum
}
