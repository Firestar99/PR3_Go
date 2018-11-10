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
	//select initial layout here
	var layout layout = Glider

	//setup
	height := layout.height
	width := layout.width
	var array = CreateArray(height, width)
	var next = CreateArray(height, width)
	ParseLayout(array, height, width, layout.initial)

	for i := 0; i < layout.generations; i++ {
		//insert code here!

		//sleep
		time.Sleep(200 * time.Millisecond)
	}
}

func Next(array *[][]bool, height, width int, y, x int) bool {
	//insert code here!
}

//premade functions
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
