package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/**
--- Day 4: Passport Processing ---

You arrive at the airport only to realize that you grabbed your North Pole Credentials instead of your passport. While these documents are extremely similar, North Pole Credentials aren't issued by a country and therefore aren't actually valid documentation for travel in most of the world.

It seems like you're not the only one having problems, though; a very long line has formed for the automatic passport scanners, and the delay could upset your travel itinerary.

Due to some questionable network security, you realize you might be able to solve both of these problems at the same time.

The automatic passport scanners are slow because they're having trouble detecting which passports have all required fields. The expected fields are as follows:

byr (Birth Year)
iyr (Issue Year)
eyr (Expiration Year)
hgt (Height)
hcl (Hair Color)
ecl (Eye Color)
pid (Passport ID)
cid (Country ID)
Passport data is validated in batch files (your puzzle input). Each passport is represented as a sequence of key:value pairs separated by spaces or newlines. Passports are separated by blank lines.

Here is an example batch file containing four passports:

ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
The first passport is valid - all eight fields are present. The second passport is invalid - it is missing hgt (the Height field).

The third passport is interesting; the only missing field is cid, so it looks like data from North Pole Credentials, not a passport at all! Surely, nobody would mind if you made the system temporarily ignore missing cid fields. Treat this "passport" as valid.

The fourth passport is missing two fields, cid and byr. Missing cid is fine, but missing any other field is not, so this passport is invalid.

According to the above rules, your improved system would report 2 valid passports.

Count the number of valid passports - those that have all required fields. Treat cid as optional. In your batch file, how many passports are valid?
*/

// " byr:1960\nhgt:183cm pid:764315947 eyr:2030\nhcl:#ceb3a1 ecl:brn\n-------------\n"

type passport struct {
	byr []string // (Birth Year)
	iyr []string // (Issue Year)
	eyr []string // (Expiration Year)
	hgt []string // (Height)
	hcl []string // (Hair Color)
	ecl []string // (Eye Color)
	pid []string // (Passport ID)
	// we ignore cid entirely
}

var byrRE = regexp.MustCompile(`byr:(#?\w+)`)
var iyrRE = regexp.MustCompile(`iyr:(#?\w+)`)
var eyrRE = regexp.MustCompile(`eyr:(#?\w+)`)
var hgtRE = regexp.MustCompile(`hgt:(#?\w+)`) // cm|inch
var hclRE = regexp.MustCompile(`hcl:(#?\w+)`)
var eclRE = regexp.MustCompile(`ecl:(#?\w+)`)
var pidRE = regexp.MustCompile(`pid:(#?\w+)`)

func day4_part1() {
	contents := getFilesContents("day05.input")
	passports := strings.Split(contents, "\n\n")
	validPassportsCount := 0
	for _, pass := range passports {
		fmt.Println("-------------\n", pass)
		data := passport{
			byr: byrRE.FindStringSubmatch(pass),
			iyr: iyrRE.FindStringSubmatch(pass),
			eyr: eyrRE.FindStringSubmatch(pass),
			hgt: hgtRE.FindStringSubmatch(pass),
			hcl: hclRE.FindStringSubmatch(pass),
			ecl: eclRE.FindStringSubmatch(pass),
			pid: pidRE.FindStringSubmatch(pass),
		}
		fmt.Print(data)
		if data.byr != nil &&
			data.iyr != nil &&
			data.eyr != nil &&
			data.hgt != nil &&
			data.hcl != nil &&
			data.ecl != nil &&
			data.pid != nil {
			validPassportsCount++
			fmt.Print("     ---    passport is valid \n")
		} else {
			fmt.Print("     ---    passport is NOT valid --- \n")
		}
	}
	fmt.Println("Found a total of", validPassportsCount, "valid passports.")
}

/**
--- Part Two ---

The line is moving more quickly now, but you overhear airport security talking about how passports with invalid data are getting through. Better add some data validation, quick!

You can continue to ignore the cid field, but each other field has strict rules about what values are valid for automatic validation:

byr (Birth Year) - four digits; at least 1920 and at most 2002.
iyr (Issue Year) - four digits; at least 2010 and at most 2020.
eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
hgt (Height) - a number followed by either cm or in:
If cm, the number must be at least 150 and at most 193.
If in, the number must be at least 59 and at most 76.
hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
pid (Passport ID) - a nine-digit number, including leading zeroes.
cid (Country ID) - ignored, missing or not.
Your job is to count the passports where all required fields are both present and valid according to the above rules.
*/

var byr2RE = regexp.MustCompile(`byr:(\d+)(\W|$)`)
var iyr2RE = regexp.MustCompile(`iyr:(\d+)(\W|$)`)
var eyr2RE = regexp.MustCompile(`eyr:(\d+)(\W|$)`)
var hgt2RE = regexp.MustCompile(`hgt:(\d+)(in|cm)(\W|$)`) // cm|inch
var hcl2RE = regexp.MustCompile(`hcl:#(([a-f]|[0-9]){6})(\W|$)`)
var ecl2RE = regexp.MustCompile(`ecl:(\w+)(\W|$)`)
var pid2RE = regexp.MustCompile(`pid:(\d{9})(\D|$)`)

func day4_part2() {
	file, err := os.Open("day04.input")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	filebyt, err := ioutil.ReadAll(file)
	filestr := string(filebyt)
	passports := strings.Split(filestr, "\n\n")
	validPassportsCount := 0
	for _, pass := range passports {
		fmt.Println("-------------\n", pass)
		data := passport{
			byr: byr2RE.FindStringSubmatch(pass), // 4 digits 1920-2002
			iyr: iyr2RE.FindStringSubmatch(pass), // 4 digits 2010-2020
			eyr: eyr2RE.FindStringSubmatch(pass), // 4 digits 2020-2030
			hgt: hgt2RE.FindStringSubmatch(pass), // \d cm 150-193, in 59-76
			hcl: hcl2RE.FindStringSubmatch(pass), // # + 6 [0-9]|[a-f]
			ecl: ecl2RE.FindStringSubmatch(pass), // [amb blu brn gry grn hzl oth] + len(3)
			pid: pid2RE.FindStringSubmatch(pass), // 9 \d including leading 0's
		}
		fmt.Println(data)

		if data.byr == nil ||
			data.iyr == nil ||
			data.eyr == nil ||
			data.hgt == nil ||
			data.hcl == nil ||
			data.ecl == nil ||
			data.pid == nil {
			fmt.Println("Missing information or invalid match")
			continue
		}

		if byr, err := strconv.Atoi(data.byr[1]); err != nil {
			log.Panicln(err)
		} else {
			if byr < 1920 || byr > 2002 {
				fmt.Println("found byr but not in range 1920-2002")
				continue
			}
		}

		if iyr, err := strconv.Atoi(data.iyr[1]); err != nil {
			log.Panicln(err)
		} else {
			if iyr < 2010 || iyr > 2020 {
				fmt.Println("found iyr but not in range 2010-2020")
				continue
			}
		}

		if eyr, err := strconv.Atoi(data.eyr[1]); err != nil {
			log.Panicln(err)
		} else {
			if eyr < 2020 || eyr > 2030 {
				fmt.Println("found eyr but not in range 2020-2030")
				continue
			}
		}

		if hgt, err := strconv.Atoi(data.hgt[1]); err != nil {
			log.Panicln(err)
		} else {
			if data.hgt[2] == "cm" {
				if hgt < 150 || hgt > 193 {
					fmt.Println("height in cm but not in range 150-193cm")
					continue
				}
			} else if data.hgt[2] == "in" {
				if hgt < 59 || hgt > 76 {
					fmt.Println("height in in but not in range 59-76in")
					continue
				}
			} else {
				fmt.Println("invalid height")
				continue
			}
		}

		ecl := data.ecl[1]
		if len(ecl) != 3 {
			fmt.Println("found ecl but not 3 chars long")
			continue
		}
		ecls := "amb blu brn gry grn hzl oth"
		if !strings.Contains(ecls, ecl) {
			fmt.Println("invalid ecl, not in range [amb blu brn gry grn hzl oth]")
			continue
		}

		pid := data.pid[1]
		if len(pid) != 9 {
			fmt.Println("found pid but not 9 chars long")
			continue
		}

		validPassportsCount++
		fmt.Println("VALID")
	}
	fmt.Println("\nFound a total of", validPassportsCount, "valid passports.\n")
}
