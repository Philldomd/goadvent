/*--------------------------------------
Task1: 24542    [nil]   Time: 0s
Task2: 8736438  [nil]   Time: 4.0148ms
----------------------------------------*/

package solvers

import (
	"strconv"
	"strings"
)

type ScratchCards struct {
	Solve
}

func (scratchCards *ScratchCards) Task1(data string) string {
	var cards []string = strings.Split(data, "\n")
	sum := 0
	for _, card := range cards {
		if len(card) == 0 {
			continue
		}
		split := strings.Split(strings.Split(card, ":")[1], "|")
		s := 0
		for _, num := range strings.Fields(split[1]) {
			for _, v := range strings.Fields(split[0]) {
				if num == v {
					if s == 0 {
						s += 1
					} else {
						s *= 2
					}
				}
			}
		}
		sum += s
	}
	return strconv.Itoa(sum)
}

func (scratchCards *ScratchCards) Task2(data string) string {
	var cards []string = strings.Split(data, "\n")
	var mult []int
	for i := 0; i < len(cards); i++ {
		mult = append(mult, 0)
	}
	sum := 0
	for r, card := range cards {
		if len(card) == 0 {
			continue
		}
		split := strings.Split(strings.Split(card, ":")[1], "|")
		s := 0
		for _, num := range strings.Fields(split[1]) {
			for _, v := range strings.Fields(split[0]) {
				if num == v {
					s += 1
				}
			}
		}
		mult[r] += 1
		for j := r + 1; j < r+s+1; j++ {
			mult[j] += mult[r]
		}
	}
	for _, x := range mult {
		sum += x
		if x == 0 {
			break
		}
	}
	return strconv.Itoa(sum)
}
