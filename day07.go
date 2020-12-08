package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

/**
--- Day 7: Handy Haversacks ---

You land at the regional airport in time for your next flight. In fact, it looks like you'll even have time to grab some food: all flights are currently delayed due to issues in luggage processing.

Due to recent aviation regulations, many rules (your puzzle input) are being enforced about bags and their contents; bags must be color-coded and must contain specific quantities of other color-coded bags. Apparently, nobody responsible for these regulations considered how long they would take to enforce!

For example, consider the following rules:

light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
These rules specify the required contents for 9 bag types. In this example, every faded blue bag is empty, every vibrant plum bag contains 11 bags (5 faded blue and 6 dotted black), and so on.

You have a shiny gold bag. If you wanted to carry it in at least one other bag, how many different bag colors would be valid for the outermost bag? (In other words: how many colors can, eventually, contain at least one shiny gold bag?)

In the above rules, the following options would be available to you:

A bright white bag, which can hold your shiny gold bag directly.
A muted yellow bag, which can hold your shiny gold bag directly, plus some other bags.
A dark orange bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
A light red bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
So, in this example, the number of bag colors that can eventually contain at least one shiny gold bag is 4.

How many bag colors can eventually contain at least one shiny gold bag? (The list of rules is quite long; make sure you get all of it.)

Your puzzle answer was 235.
*/

func day7_part1() {
	contents := getFilesContents("day07.input")
	rawlines := strings.Split(contents, "\n")
	var lines [][]string
	for _, rawline := range rawlines {
		elements := strings.Split(rawline, ", ")
		lines = append(lines, elements)
	}

	findlist := []string{"shiny gold"}
	var tofind string
	results := make(map[string]int) // dict of all bags upstream
	for len(findlist) > 0 {
		tofind, findlist = findlist[0], findlist[1:] // "slice from front"
		for _, line := range lines {
			for _, v := range line[1:] {
				if strings.Contains(v, tofind) {
					findlist = append(findlist, line[0]) // add it to inspection list
					results[line[0]] = 1
				}
			}
		}
		fmt.Println(len(findlist), findlist)
	}
	count := 0
	for k, v := range results {
		count++
		fmt.Println(k, v, count)
	}
}

/**
--- Part Two ---

It's getting pretty expensive to fly these days - not because of ticket prices, but because of the ridiculous number of bags you need to buy!

Consider again your shiny gold bag and the rules from the above example:

faded blue bags contain 0 other bags.
dotted black bags contain 0 other bags.
vibrant plum bags contain 11 other bags: 5 faded blue bags and 6 dotted black bags.
dark olive bags contain 7 other bags: 3 faded blue bags and 4 dotted black bags.
So, a single shiny gold bag must contain 1 dark olive bag (and the 7 bags within it) plus 2 vibrant plum bags (and the 11 bags within each of those): 1 + 1*7 + 2 + 2*11 = 32 bags!

Of course, the actual rules have a small chance of going several levels deeper than this example; be sure to count all of the bags, even if the nesting becomes topologically impractical!

Here's another example:

shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.
In this example, a single shiny gold bag must contain 126 other bags.

How many individual bags are required inside your single shiny gold bag?

Your puzzle answer was 158493.
*/

type Bag struct {
	name     string
	mess     []string
	contents []Bag
}

var reference = make(map[string][]string)

func day7_part2() {
	contents := getFilesContents("day07.input")
	start := time.Now()
	rawlines := strings.Split(contents, "\n")
	for _, rawline := range rawlines {
		elements := strings.Split(rawline, ", ")
		reference[elements[0]] = elements[1:] // e.g. {shiny gold: ['2 dark red reference', '2 olive plum']}
	}
	bag := Bag{name: "shiny gold"}
	bag.mess = reference[bag.name]
	parseBag(bag)
	elapsed := time.Since(start)
	fmt.Println("Elapsed", elapsed, count-1)
}

var count uint = 0

func parseBag(pbag Bag) {
	// timeTrack(time.Now(), pbag.name)
	// fmt.Println(pbag, count)
	count++
	for _, b := range pbag.mess {
		tokens := strings.Split(b, " ")
		if tokens[0] != "no" {
			numbags, _ := strconv.Atoi(tokens[0])
			for i := 0; i < numbags; i++ {
				newbag := Bag{name: strings.Join(tokens[1:], " ")}
				newbag.mess = reference[newbag.name]
				pbag.contents = append(pbag.contents, newbag)
				parseBag(newbag)
			}
		}
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
