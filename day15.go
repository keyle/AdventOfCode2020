package main

import "fmt"

// var input2 = []int{2, 15, 0, 9, 1, 20}

func day15_part1() {
	var expected = []int{0, 3, 6, 0, 3, 3, 1, 0, 4, 0}
	var input = []int{0, 3, 6}
	var dict = map[int][]int{}
	runfor, i := 9, 0
	for i <= runfor {
		v := input[i]
		input = append(input, v)
		if len(dict[v]) > 0 {
			dist := dict[v][len(dict[v])-1] - dict[v][len(dict[v])-2]
			fmt.Println(dict[v][len(dict[v])-1], dict[v][len(dict[v])-2])
			// fmt.Println("\ti", i, "- dict", dict[v][1], " = dist ", dist)
			input[i+1] = dist
		} else {
			dict[v] = append(dict[v], i)
			fmt.Println("new")
		}
		fmt.Println("Turn", i+1, " = ", v, "\t", expected[i])
		dict[v] = append(dict[v], i)
		// fmt.Println(dict)
		i++
	}

}
