package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	data := readInput()
	fmt.Println("part one:", partOne(data))
	fmt.Println("part two:", partTwo(data))
}

type FieldValidators map[string]string

func areAllRequiredFieldsPresent(fields []string, validators FieldValidators) bool {
	for key := range validators {
		found := false
		for _, field := range fields {
			if key == strings.Split(field, ":")[0] {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func validateFields(fields []string, validators FieldValidators) bool {
	for _, field := range fields {
		splitted := strings.Split(field, ":")
		regex, ok := validators[splitted[0]]
		if ok {
			if !regexp.MustCompile(regex).MatchString(splitted[1]) {
				return false
			}
		}
	}
	return true
}

func validatePassport(data string, validators FieldValidators) bool {
	fields := strings.Fields(data)

	if !areAllRequiredFieldsPresent(fields, validators) {
		return false
	}

	if !validateFields(fields, validators) {
		return false
	}

	return true
}

func partOne(data []string) int {
	requiredElements := map[string]string{"byr": ".*",
		"iyr": ".*",
		"eyr": ".*",
		"hgt": ".*",
		"hcl": ".*",
		"ecl": ".*",
		"pid": ".*"}

	result := 0
	for _, passport := range data {
		if validatePassport(passport, requiredElements) {
			result++
		}
	}
	return result
}

func partTwo(data []string) int {
	requiredElementWithValidators := map[string]string{"byr": "^([1][9][2-9][0-9]|[2][0][0][0-2])$",
		"iyr": "^([2][0][1][0-9]|[2][0][2][0])$",
		"eyr": "^([2][0][2][0-9]|[2][0][3][0])$",
		"hgt": "^([1][5-8][0-9]cm|[1][9][0-3]cm|[5][9]in|[6][0-9]in|[7][0-6]in)$",
		"hcl": "^#([a-f]|\\d){6}$",
		"ecl": "(amb|blu|brn|gry|grn|hzl|oth)",
		"pid": "^\\d{9}$"}

	result := 0
	for _, passport := range data {
		if validatePassport(passport, requiredElementWithValidators) {
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
