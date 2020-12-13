package main

import (
	"fmt"
	"regexp"
	"strings"
)

/**
--- Day 11: Seating System ---

Your plane lands with plenty of time to spare. The final leg of your journey is a ferry that goes directly to the tropical island where you can finally start your vacation. As you reach the waiting area to board the ferry, you realize you're so early, nobody else has even arrived yet!

By modeling the process people use to choose (or abandon) their seat in the waiting area, you're pretty sure you can predict the best place to sit. You make a quick map of the seat layout (your puzzle input).

The seat layout fits neatly on a grid. Each position is either floor (.), an empty seat (L), or an occupied seat (#). For example, the initial seat layout might look like this:

L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
Now, you just need to model the people who will be arriving shortly. Fortunately, people are entirely predictable and always follow a simple set of rules. All decisions are based on the number of occupied seats adjacent to a given seat (one of the eight positions immediately up, down, left, right, or diagonal from the seat). The following rules are applied to every seat simultaneously:

- If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
- If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
- Otherwise, the seat's state does not change.

Floor (.) never changes; seats don't move, and nobody sits on the floor.

After one round of these rules, every seat in the example layout becomes occupied:

#.##.##.#########.###.#.#..#..####.##.###.##.##.###.#####.##..#.#.....###########.######.##.#####.##

After a second round, the seats with four or more occupied adjacent seats become empty again:

#.LL.L#.###LLLLLL.L#L.L.L..L..#LLL.LL.L##.LL.LL.LL#.LLLL#.##..L.L.....#LLLLLLLL##.LLLLLL.L#.#LLLL.##

This process continues for three more rounds:

#.##.L#.###L###LL.L#L.#.#..#..#L##.##.L##.##.LL.LL#.###L#.##..#.#.....#L######L##.LL###L.L#.#L###.##

#.#L.L#.###LLL#LL.L#L.L.L..#..#LLL.##.L##.LL.LL.LL#.LL#L#.##..L.L.....#L#LLLL#L##.LLLLLL.L#.#L#L#.##

#.#L.L#.###LLL#LL.L#L.#.L..#..#L##.##.L##.#L.LL.LL#.#L#L#.##..L.L.....#L#L##L#L##.LLLLLL.L#.#L#L#.##

At this point, something interesting happens: the chaos stabilizes and further applications of these rules cause no seats to change state! Once people stop moving around, you count 37 occupied seats.

Simulate your seating area by applying the seating rules repeatedly until no seats change state. How many seats end up occupied?

Your puzzle answer was 2261.
*/

func day11_part1() {
	contents := getFilesContents("day11.input")
	rows := strings.Split(contents, "\n")
	grid := makeBlankGrid(len(rows), len(rows[0]))
	for i, line := range rows {
		col := strings.Split(line, "")
		for j, c := range col {
			grid[i][j] = c
		}
	}
	// printGrid(grid)
	ll := len(rows)
	ww := len(rows[0])

	itercount := 0
	var finalGrid string
	for true {
		itercount++
		res := rewriteGrid(grid, ll, ww)
		grd := mutateGrid(res, ll, ww, 4)
		a := fmt.Sprint(grid)
		b := fmt.Sprint(grd)
		if a == b {
			finalGrid = a
			break
		}
		// fmt.Println(itercount)
		// printGrid(grd)
		grid = grd
	}
	filled := filledRE.FindAllStringIndex(finalGrid, -1)
	fmt.Println("itercount", itercount-1, "until the same with", len(filled), "seats taken.")
}

func day11_part2() {
	contents := getFilesContents("day11.input")
	rows := strings.Split(contents, "\n")
	grid := makeBlankGrid(len(rows), len(rows[0]))
	for i, line := range rows {
		col := strings.Split(line, "")
		for j, c := range col {
			grid[i][j] = c
		}
	}
	ll := len(rows)
	ww := len(rows[0])

	itercount := 0
	var finalGrid string
	for true {
		itercount++
		res := rewriteGrid2(grid, ll, ww)
		grd := mutateGrid(res, ll, ww, 5)
		a := fmt.Sprint(grid)
		b := fmt.Sprint(grd)
		if a == b {
			finalGrid = a
			break
		}
		// fmt.Println(itercount)
		// printGrid(grd)
		grid = grd
	}

	filled := filledRE.FindAllStringIndex(finalGrid, -1)
	fmt.Println("part2: itercount", itercount-1, "until the same with", len(filled), "seats taken.")
}

func makeBlankGrid(length int, width int) [][]string {
	grid := make([][]string, length)
	for i := range grid {
		grid[i] = make([]string, width)
	}
	return grid
}

// rewriteGrid makes a grid with the seats logic of length ll and width ww
// each item is a string with chars in this order: Self, East of self, South East, South, etc. Clockwise
// making each item exactly 9 chars long
func rewriteGrid(grid [][]string, ll int, ww int) [][]string {
	newGrid := makeBlankGrid(ll, ww)
	for r, row := range grid {
		for c, seat := range row {
			if seat == "." {
				newGrid[r][c] = "."
				continue
			}
			st := seat
			if c+1 < ww {
				st += grid[r][c+1] // E
			} else {
				st += "x"
			}
			if r+1 < ll && c+1 < ww {
				st += grid[r+1][c+1] // SE
			} else {
				st += "x"
			}
			if r+1 < ll {
				st += grid[r+1][c] // S
			} else {
				st += "x"
			}
			if r+1 < ll && c > 0 {
				st += grid[r+1][c-1] // SW
			} else {
				st += "x"
			}
			if c > 0 {
				st += grid[r][c-1] // W
			} else {
				st += "x"
			}
			if r > 0 && c > 0 {
				st += grid[r-1][c-1] // NW
			} else {
				st += "x"
			}
			if r > 0 {
				st += grid[r-1][c] // N
			} else {
				st += "x"
			}
			if r > 0 && c+1 < ww {
				st += grid[r-1][c+1] // NE
			} else {
				st += "x"
			}
			newGrid[r][c] = st
		}
	}
	return newGrid
}

// lookFwd in grid of ll length and ww width, for row r, and col c, with delta dr applied to rows and delta dc applied to columns
// finds a L or # or gives x out of bounds @TODO double check what if floor runs to the wall
func lookFwd(grid [][]string, ll int, ww int, r int, c int, dr int, dc int) string {
	st := ""
	for true {
		if r+dr < ll && r+dr >= 0 && c+dc < ww && c+dc >= 0 {
			if grid[r+dr][c+dc] == "L" ||
				grid[r+dr][c+dc] == "#" {
				st += grid[r+dr][c+dc]
				break
			}
		} else {
			st += "x"
			break
		}
		c += dc
		r += dr
	}
	return st
}

// rewriteGrid makes a grid with the seats logic of length ll and width ww
// each item is a string with chars in this order: Self, East of self, South East, South, etc. Clockwise
// making each item exactly 9 chars long
func rewriteGrid2(grid [][]string, ll int, ww int) [][]string {
	newGrid := makeBlankGrid(ll, ww)
	for r, row := range grid {
		for c, seat := range row {
			if seat == "." {
				newGrid[r][c] = "."
				continue
			}
			st := seat
			st += lookFwd(grid, ll, ww, r, c, 0, +1)  // E
			st += lookFwd(grid, ll, ww, r, c, +1, +1) // SE
			st += lookFwd(grid, ll, ww, r, c, +1, 0)  // S
			st += lookFwd(grid, ll, ww, r, c, +1, -1) // SW
			st += lookFwd(grid, ll, ww, r, c, 0, -1)  // W
			st += lookFwd(grid, ll, ww, r, c, -1, -1) // NW
			st += lookFwd(grid, ll, ww, r, c, -1, 0)  // N
			st += lookFwd(grid, ll, ww, r, c, -1, +1) // NE
			// // E
			// for true {
			// 	if c+1 < ww {
			// 		if grid[r][c+1] == "L" ||
			// 			grid[r][c+1] == "#" {
			// 			st += grid[r][c+1]
			// 			break
			// 		}
			// 	} else {
			// 		st += "x"
			// 		break
			// 	}
			// 	c++
			// }
			newGrid[r][c] = st
		}
	}
	return newGrid
}

// var emptyRE = regexp.MustCompile("L")
var filledRE = regexp.MustCompile("#")

// mutateGrid takes a grid and moves it following the rules by 1 generation
func mutateGrid(grid [][]string, ll int, ww int, tolerance int) [][]string {
	newGrid := makeBlankGrid(ll, ww)
	for r, row := range grid {
		for c, seat := range row {
			// fmt.Print(r, c, "- ")
			if rune(seat[0]) == rune("."[0]) {
				newGrid[r][c] = "."
				// fmt.Println("floor", grid[r][c], newGrid[r][c])
				continue
			}
			// fmt.Println("seat", "is as before", newGrid[r][c])
			newGrid[r][c] = string(rune((grid[r][c])[0]))
			filled := filledRE.FindAllStringIndex(seat[1:], -1)
			if rune(seat[0]) == rune("L"[0]) { // seat is empty
				// - If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
				if len(filled) == 0 {
					newGrid[r][c] = "#"
					// fmt.Println("seat", "is now occupied", "was", grid[r][c], "now", newGrid[r][c])
				}
			} else if rune(seat[0]) == rune("#"[0]) { // seat is occupied
				// - If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
				if len(filled) >= tolerance {
					newGrid[r][c] = "L"
					// fmt.Println("seat", "is now freed", "was", grid[r][c], "now", newGrid[r][c])
				}
			}
		}
	}
	return newGrid
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		for _, v := range row {
			fmt.Print(v + " ")
		}
		fmt.Println("")
	}
	fmt.Println("")
}
