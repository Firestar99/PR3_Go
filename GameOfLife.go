package main

import (
	"fmt"
	"time"
)

func main() {
	height := 29
	width := 29
	var array = CreateArray(height, width)
	var next = CreateArray(height, width)
	ParseLayout(array, height, width, []string{
		"                             ",
		"                             ",
		"                             ",
		"                             ",
		"                             ",
		"                             ",
		"                             ",
		"                             ",
		"                             ",
		"                             ",
		"             ###             ",
		"             # #             ",
		"             # #             ",
		"              +              ",
		"             # #             ",
		"             # #             ",
		"             ###             ",
		"                             ",
		"                             ",
		"                             ",
		"                             ",
		"                             ",
		"                             ",
		"                             ",
		"                             ",
		"                             ",
		"                             ",
		"                             ",
		"                             ",
	})

	for i := 0; i < 100; i++ {
		printOut(array, height, width)
		Step(array, next, height, width)

		temp := array
		array = next
		next = temp
		time.Sleep(200 * time.Millisecond)
	}
}

func CreateArray(height, width int) *[][]bool {
	array := make([][]bool, height)
	for y := 0; y < height; y++ {
		array[y] = make([]bool, width)
	}
	return &array
}

func ParseLayout(array *[][]bool, height, width int, input []string) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			(*array)[y][x] = input[y][x] == '#'
		}
	}
}

func Step(array, next *[][]bool, height, width int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			(*next)[y][x] = Next(array, height, width, y, x)
		}
	}
}

func Next(array *[][]bool, height, width int, y, x int) bool {
	var surrounding int = 0
	for offy := -1; offy <= 1; offy++ {
		for offx := -1; offx <= 1; offx++ {
			if !(offx == 0 && offy == 0) {
				if ArrayGetOrDefault(array, height, width, offy+y, offx+x, false) {
					surrounding++
				}
			}
		}
	}

	if surrounding < 2 {
		return false
	}
	if surrounding == 2 {
		return (*array)[y][x]
	}
	if surrounding == 3 {
		return true
	}
	return false
}

func ArrayGetOrDefault(array *[][]bool, height, width int, y, x int, def bool) bool {
	if 0 > y || y >= height {
		return def
	}
	line := (*array)[y]
	if 0 > x || x >= width {
		return def
	}
	return line[x]
}

func printOut(array *[][]bool, height, width int) {
	var out = ""
	out += "\x0c"
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if (*array)[y][x] {
				out += "#"
			} else {
				out += " "
			}
		}
		out += "\n"
	}
	fmt.Print(out)
}
