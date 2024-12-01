/*--------------------------------------
Task1: 505379   [OK]    Time: 0s
Task2: 263211   [OK]    Time: 1.5935ms
----------------------------------------*/

package snowsolver

import (
	"strconv"
)

type LensLibrary struct {
	Solve
}

type Box struct {
	label string
	lens  int
}

var Boxes = make(map[int]([]Box))

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

func (lensLibrary LensLibrary) ReadLensData(data string) {
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
				le, _ := strconv.Atoi(lens)
				found := false
				for i, v := range Boxes[box] {
					if v.label == label {
						Boxes[box][i].lens = le
						found = true
					}
				}
				if !found {
					Boxes[box] = append(Boxes[box], Box{label, le})
				}
				adding = false
			} else if subtracting {
				if v, ok := Boxes[box]; ok {
					for i, b := range v {
						if b.label == label {
							Boxes[box] = append(Boxes[box][:i], Boxes[box][i+1:]...)
						}
					}
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
		le, _ := strconv.Atoi(lens)
		found := false
		for i, v := range Boxes[box] {
			if v.label == label {
				Boxes[box][i].lens = le
				found = true
			}
		}
		if !found {
			Boxes[box] = append(Boxes[box], Box{label, le})
		}
		adding = false
	} else if subtracting {
		if v, ok := Boxes[box]; ok {
			for i, b := range v {
				if b.label == label {
					Boxes[box] = append(Boxes[box][:i], Boxes[box][i+1:]...)
				}
			}
		}
		subtracting = false
	}
}

func (LensLibrary LensLibrary) CalcFocusPower() int {
	value := 0
	for i, b := range Boxes {
		for j, l := range b {
			value += (i + 1) * (j + 1) * l.lens
		}
	}
	return value
}

func (lensLibrary *LensLibrary) Task2(data string) string {
	lensLibrary.ReadLensData(data)
	return strconv.Itoa(lensLibrary.CalcFocusPower())
}
