package main

import (
	"fmt"
)

func day13_part1() {
	epoch := 1006697
	times := []int{13, 17, 19, 23, 29, 37, 41, 641, 661}

	for _, v := range times {
		mod := epoch % v
		dist := v - mod
		fmt.Println(v, dist, "answer", v*dist)
	}
}

// var times = []uint64{7, 13, 1, 1, 59, 1, 31, 19}
//
var times = []uint64{13, 1, 1, 41, 1, 1, 1, 1, 1, 1, 1, 1, 1, 641, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 19, 1, 1, 1, 1, 17, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 29, 1, 661, 1, 1, 1, 1, 1, 37, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 23}

var arrlen = uint64(len(times)) - 1
var answer = uint64(0)

// NOTE Solution adapted from blanchg https://github.com/blanchg/aoc202/blob/main/13.py
// I first tried to brute force it; then I was close but I had it backwards (epoch then bus rather than bus then epoch)
func day13_part2() {
	epoch := uint64(1)
	arr := [][]int{}
	for i, v := range times {
		if v > 1 {
			arr = append(arr, []int{int(v), i})
		}
	}
	fmt.Println(arr)
	res := uint64(arr[0][0])
	epoch = 1
	for _, v := range arr {
		num := v[0]
		idx := v[1]
		println("num", num, "epoch", epoch, "res", res)
		for {
			if (res+uint64(idx))%uint64(num) == 0 {
				break
			}
			res += epoch
		}
		epoch = epoch * uint64(num)
	}
	fmt.Println(epoch, res)
	fmt.Println("found the answer!", res)
}

func day13_part2_brute() {
	fmt.Println(times)
	epoch := uint64(0)
	for {
		go test(epoch)
		go test(epoch + 1)
		go test(epoch + 2)
		go test(epoch + 3)
		go test(epoch + 4)
		epoch += 5
		if answer != 0 {
			break
		}
	}
}

func test(epoch uint64) {
	i := uint64(0)
	for {
		if (epoch+i)%times[i] != 0 {
			// println("")
			// fmt.Println(epoch+uint64(i), "%", times[i], (epoch+uint64(i)%uint64(times[i])), "BAD")
			return
		} else {
			// print(".")
			// fmt.Println(epoch+uint64(i), "%", times[i], (epoch + uint64(i)%uint64(times[i])), "OK")
			if i >= arrlen {
				fmt.Println("found the answer!", epoch)
				answer = epoch
				return
			}
		}
		i++
	}
}

// time     bus 7   bus 13  bus 59  bus 31  bus 19
// 1068773    .       .       .       .       .
// 1068774    D       .       .       .       .
// 1068775    .       .       .       .       .
// 1068776    .       .       .       .       .
// 1068777    .       .       .       .       .
// 1068778    .       .       .       .       .
// 1068779    .       .       .       .       .
// 1068780    .       .       .       .       .
// 1068781    D       .       .       .       .
// 1068782    .       D       .       .       .
// 1068783    .       .       .       .       .
// 1068784    .       .       .       .       .
// 1068785    .       .       D       .       .
// 1068786    .       .       .       .       .
// 1068787    .       .       .       D       .
// 1068788    D       .       .       .       D
// 1068789    .       .       .       .       .
// 1068790    .       .       .       .       .
