/*------------------------------------
Task1: 505379   [OK]    Time: 0s

--------------------------------------*/

package solvers

import (
	"fmt"
	"strconv"
)

type LensLibrary struct {
	Solve
}

type Box struct {
	label string
	lens int
}

var Boxes = make(map[int](Box))

func (lensLibrary *LensLibrary) Task1(data string) string {
	char_values := []int{}
	temp_char := 0
	for _, r := range data {
		switch r {
		case '\n':
			continue
		case ',':
			char_values = append(char_values, temp_char)
			temp_char = 0
		default:
			temp_char = ((temp_char + int(r)) * 17) % 256
		}
	}
	char_values = append(char_values, temp_char)
	return strconv.Itoa(Sum(char_values))
}

func (lensLibrary *LensLibrary) Task2(data string) string {
	box := 0
	label := ""
	lens := ""
	adding := false
	subtracting := false
	for _, r := range data {
		switch r {
		case '\n':
			continue
		case ',':
			if adding {
				if Boxes[box] == nil {
				  Boxes[box] = make(map[string]int)
				}
				Boxes[box][label], _ = strconv.Atoi(lens)
				adding = false
			} else if subtracting {
				if v, ok := Boxes[box]; ok {
					delete(v, label)
				}
				if len(Boxes[box]) == 0 {
					delete(Boxes, box)
				}
				subtracting = false
			}
			label = ""
			lens = ""
			box = 0
		case '-':
			subtracting = true
    case '=':
      adding = true
		default:
			if !adding {
				label = label + string(r)
				box = ((box + int(r)) * 17) % 256
			} else {
				lens = lens + string(r)
			}
		}
	}
	if adding {
		if Boxes[box] == nil {
			Boxes[box] = make(map[string]int)
		}
		Boxes[box][label], _ = strconv.Atoi(lens)
		adding = false
	} else if subtracting {
		if v, ok := Boxes[box]; ok {
			delete(v, label)
		}
		if len(Boxes[box]) == 0 {
			delete(Boxes, box)
		}
		subtracting = false
	}
	fmt.Println(Boxes)
	return "nil"
}
