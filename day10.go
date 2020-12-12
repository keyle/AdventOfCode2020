package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/**

jolts

adapters

output jotlage

1 2 3 jolts lower than its rating -> ok

device can take 3 jolts higher than the findHighest rated adapter

3 9 6 -> 12 jv as 9+3=12 jv

charging outlate 0 jv

use every adapter
distribution jv diff between the charging outlet, the adapters and your device

for example,s uppose that in your bar, you have adapters with the following joltage ratings:

16
10
15
5
1
11
7
19
6
12
4
rated for 22 (19+3)

because adpaters can only connect to a source 1-3 jv lower tha tits rating, in order ot use every adapter you'd need to choose them like this
- from 0 you can only go to 1 2 3, you have a 1
- from 1, your only choice is 4
- from 4 adapters rated 5 6 7 are valid however pick lowest and closest, 5
- then 6 7 following that logic
- 7 -> 10 is only 3 diff
- from 10 go 11/12, pick 11
- then 12.
- after 12, 15
- then 16, 19.
- finally your device builtin adapter is always 3 higher than th findHighest, so it's 19+3 = 22.

find how many differences of 5 jots and 3 jolts 22 and 10. 22 * 10 = 220.
*/

func day10_part1() {
	contents := getFilesContents("day10.input")
	data := strings.Split(contents, "\n")
	adapters := []int{}
	for _, v := range data {
		u, _ := strconv.Atoi(v)
		adapters = append(adapters, u)
	}

	sort.Ints(adapters)

	cur := 0
	threes := 0
	ones := 0
	for _, v := range adapters {
		if v-cur == 3 {
			cur = v
			threes++
		} else if v-cur == 1 {
			cur = v
			ones++
		}
		fmt.Println(cur)
	}

	threes++ // last adapter to device

	fmt.Println("threes", threes, "ones", ones, "total", threes*ones)

	// highest := findHighest(adapters)
	// var cur int = 0
	// for i <= highest {
	// 	nextSmallest := findSmallest(adapters)
	//
	// }
}

func day10_part2() {
	contents := getFilesContents("day10.input")
	data := strings.Split(contents, "\n")
	adapters := []int{0}
	for _, v := range data {
		u, _ := strconv.Atoi(v)
		adapters = append(adapters, u)
	}
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	fmt.Println(adapters)

	marks := make(map[int]int)
	marks[0] = 1

	for i := 0; i < len(adapters); i++ {
		for j := i + 1; j < len(adapters); j++ {
			if adapters[j] > adapters[i]+3 {
				break
			}
			marks[j] += marks[i]
		}
	}
	fmt.Println(marks[len(adapters)-1])
}
