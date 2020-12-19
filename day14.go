package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func convertDecimalToBinary(number int) int {
	binary := 0
	counter := 1
	remainder := 0

	for number != 0 {
		remainder = number % 2
		number = number / 2
		binary += remainder * counter
		counter *= 10

	}
	return binary
}

/**
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
*/
var memre = regexp.MustCompile(`^mem\[(\d+)\]`)

func day14_part1() {
	contents := getFilesContents("day14.input")
	lines := strings.Split(contents, "\n")
	program := parseProgram(lines)
	// fmt.Println(program)
	dict := make(map[int]int64)
	mask := ""
	for _, cmd := range program {
		if cmd[0] == "mask" {
			mask = cmd[1]
			// fmt.Println("new mask", mask)
		} else {
			num, _ := strconv.Atoi(cmd[1])
			idx, _ := strconv.Atoi(cmd[2])
			st1 := int_to_bitstr(num)
			// fmt.Println(st1)
			res := filterThroughMask(st1, mask)
			// fmt.Println(res)
			dict[idx] = bitstr_to_int(res)
			// fmt.Println(dict)
		}
	}

	acc := int64(0)
	for _, v := range dict {
		acc += v
	}

	fmt.Println("result", acc)
	// a := 0001011
	// res, _ := strconv.ParseInt(string(a), 10, 64)
	// fmt.Println(res)
	// binary := "00001001001"
	// output, _ := strconv.ParseInt(binary, 2, 64)
	// fmt.Println(11, "is", int_to_bitstr(11))
	// fmt.Println(bitstr_to_int(int_to_bitstr(11)))
	// fmt.Println(73, "is", int_to_bitstr(73))
	// fmt.Println(bitstr_to_int(int_to_bitstr(73)))
}

func day14_part2() {
	contents := getFilesContents("day14.mock2")
	lines := strings.Split(contents, "\n")
	// fmt.Println(program)
	// dict := make(map[int][]int64)

	// digits := regexp.MustCompile(`(\d)`)

	// NOTE TODO

	var mask string
	mem1, part1 := map[int]int{}, 0
	mem2, part2 := map[int]int{}, 0
	for _, s := range lines {
		if _, err := fmt.Sscanf(s, "mask = %s", &mask); err == nil {
			continue
		}
		var addr, value int
		fmt.Sscanf(s, "mem[%d] = %d", &addr, &value)

		for i, x := 0, strings.Count(mask, "X"); i < 1<<x; i++ {
			mask := strings.NewReplacer("X", "x", "0", "X").Replace(mask)
			for _, r := range fmt.Sprintf("%0*b", x, i) {
				mask = strings.Replace(mask, "x", string(r), 1)
			}

			addr := apply(mask, addr)
			part2, mem2[addr] = part2+value-mem2[addr], value
		}

		value = apply(mask, value)
		part1, mem1[addr] = part1+value-mem1[addr], value
	}
	fmt.Println(part1)
	fmt.Println(part2)

	// acc := int64(0)
	// for _, v := range dict {
	// 	acc += v
	// }
	//
	// fmt.Println("result", acc)
}
func apply(mask string, value int) int {
	and, _ := strconv.ParseUint(strings.ReplaceAll(mask, "X", "1"), 2, 0)
	or, _ := strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 0)
	return value&int(and) | int(or)
}

// https://www.geeksforgeeks.org/generate-all-the-binary-strings-of-n-bits/
func generateAllBinaryStrings(n int, arr []string, i int, dic map[string]int) {
	if i == n {
		dic[fmt.Sprintf("%s", arr)] = len(arr)
		fmt.Println(arr, n)
		return
	}

	arr[i] = "0"
	generateAllBinaryStrings(n, arr, i+1, dic)

	arr[i] = "1"
	generateAllBinaryStrings(n, arr, i+1, dic)
}

func sliceIndices(str string, letter rune) []int {
	lst := make([]int, 1)
	runes := []rune(str)
	for i, v := range runes {
		if v == letter {
			lst = append(lst, i)
		}
	}
	return lst
}

func filterThroughMask(astr, mask string) string {
	result := []rune(astr)
	b := []rune(mask)
	for i, v := range b {
		if string(v) != "X" {
			result[i] = b[i]
		}
	}
	return string(result)
}

func filterThroughMaskv2(astr, mask string) string {
	result := []rune(astr)
	b := []rune(mask)
	for i, v := range b {
		if string(v) == "X" {
			result[i] = 'X'
		} else if string(v) == "1" {
			result[i] = b[i]
		}
	}
	return string(result)
}

func parseProgram(lines []string) [][]string {
	var program [][]string
	for _, v := range lines {
		spl := strings.Split(v, " = ")
		res := memre.FindStringSubmatch(spl[0])
		memc := ""
		if len(res) > 0 {
			memc = res[1]
		}
		// fmt.Println(isMem, spl[0], memc, spl[1])
		program = append(program, []string{spl[0], spl[1], memc})
	}
	return program
}

func int_to_bitstr(num int) string {
	binaroutput := strconv.FormatInt(int64(num), 2)
	return fmt.Sprintf("%036s", binaroutput)
}

func bitstr_to_int(bitstr string) int64 {
	output, err := strconv.ParseInt(bitstr, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return output
}
