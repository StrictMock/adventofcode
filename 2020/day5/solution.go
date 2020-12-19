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

func getListedIDs(data []string) []int {
	ids := make([]int, 0)
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
		ids = append(ids, rows[0]*8+cols[0])
	}
	return ids
}

func getAllIDs() []int {
	ids := make([]int, 0)
	for _, row := range createSlice(128) {
		for _, col := range createSlice(8) {
			ids = append(ids, row*8+col)
		}
	}
	return ids
}

func isPresent(idToFind int, ids []int) bool {
	for _, id := range ids {
		if idToFind == id {
			return true
		}
	}
	return false
}

func partOne(data []string) int {
	id := 0
	for _, currentID := range getListedIDs(data) {
		if currentID > id {
			id = currentID
		}
	}
	return id
}

func partTwo(data []string) int {
	allIds := getAllIDs()
	listedIds := getListedIDs(data)
	for _, id := range allIds {
		if isPresent(id, listedIds) {
			continue
		}
		if isPresent(id+1, listedIds) && isPresent(id-1, listedIds) {
			return id
		}
	}
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
