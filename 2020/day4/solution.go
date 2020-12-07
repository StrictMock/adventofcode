package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	data := readInput()
	counter := partOne(data)
	fmt.Println(counter)
}

func partOne(data []string) int {
	requiredElements := []string{"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid"}

	result := 0
	for _, passport := range data {
		valid := true
		for _, element := range requiredElements {
			if !strings.Contains(passport, element) {
				valid = false
			}
		}
		if valid {
			result++
		}
	}
	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() []string {
	f, err := ioutil.ReadFile("input.txt")
	check(err)

	splitted := regexp.
		MustCompile(`\n\s*\n`).
		Split(string(f), -1)

	for i, s := range splitted {
		data := strings.Map(func(r rune) rune {
			if r == '\n' {
				return ' '
			}
			return r
		}, s)
		splitted[i] = data
	}
	return splitted
}
