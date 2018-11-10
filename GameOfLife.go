package main

import (
	"fmt"
	"time"
)

type layout struct {
	height      int
	width       int
	generations int
	initial     []string
}

var Oscillator1 layout = layout{
	5,
	5,
	10,
	[]string{
		"     ",
		"  #  ",
		"  #  ",
		"  #  ",
		"     ",
	},
}

var Glider layout = layout{
	10,
	10,
	24,
	[]string{
		"          ",
		"  #       ",
		"   #      ",
		" ###      ",
		"          ",
		"          ",
		"          ",
		"          ",
		"          ",
		"          ",
	},
}

var Oscillator2 layout = layout{
	19,
	20,
	30,
	[]string{
		"                    ",
		"                    ",
		"                    ",
		"                    ",
		"                    ",
		"                    ",
		"                    ",
		"                    ",
		"       #    #       ",
		"     ## #### ##     ",
		"       #    #       ",
		"                    ",
		"                    ",
		"                    ",
		"                    ",
		"                    ",
		"                    ",
		"                    ",
		"                    ",
	},
}

var SpaceShip layout = layout{
	11,
	50,
	50,
	[]string{
		"                                                  ",
		"                                                  ",
		"                                                  ",
		"                                                  ",
		"                                                  ",
		" ####    #####    ######                          ",
		"#   #   #    #   #     #                          ",
		"    #        #         #                          ",
		"#  #    #   #    #    #                           ",
		"          #        ##                             ",
		"                                                  ",
	},
}

var Giant layout = layout{
	29,
	29,
	100,
	[]string{
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
	},
}

func main() {
	var layout layout = Glider

	height := layout.height
	width := layout.width
	var array = CreateArray(height, width)
	var next = CreateArray(height, width)
	ParseLayout(array, height, width, layout.initial)

	for i := 0; i < layout.generations; i++ {
		printOut(array, height, width)

		//step
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				(*next)[y][x] = Next(array, height, width, y, x)
			}
		}

		//swap buffers
		temp := array
		array = next
		next = temp

		//sleep
		time.Sleep(200 * time.Millisecond)
	}
	printOut(array, height, width)
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

func Next(array *[][]bool, height, width int, y, x int) bool {
	//count the surrounding fields which are alive
	var count = 0
	for offy := -1; offy <= 1; offy++ {
		for offx := -1; offx <= 1; offx++ {
			if !(offx == 0 && offy == 0) {
				if ArrayGetOrDefault(array, height, width, offy+y, offx+x, false) {
					count++
				}
			}
		}
	}

	//set the new state depending on count
	if count < 2 {
		return false
	}
	if count == 2 {
		return (*array)[y][x]
	}
	if count == 3 {
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
