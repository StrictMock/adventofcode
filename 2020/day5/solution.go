package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	data := readInput()
	fmt.Println("part one:", partOne(data))
	fmt.Println("part two:", partTwo(data))
}

func createSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = i
	}
	return slice
}

func partOne(data []string) int {
	id := 0
	for _, seat := range data {
		rows := createSlice(128)
		cols := createSlice(8)
		for _, character := range seat {
			if character == 'B' {
				rows = rows[len(rows)/2 : len(rows)]
			} else if character == 'F' {
				rows = rows[0 : len(rows)/2]
			} else if character == 'L' {
				cols = cols[0 : len(cols)/2]
			} else if character == 'R' {
				cols = cols[len(cols)/2 : len(cols)]
			}
		}
		if currentID := rows[0]*8 + cols[0]; currentID > id {
			id = currentID
		}
	}
	return id
}

func partTwo(data []string) int {
	return 0
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() []string {
	f, err := ioutil.ReadFile("input.txt")
	check(err)

	return regexp.
		MustCompile(`\s*\n`).
		Split(string(f), -1)
}
