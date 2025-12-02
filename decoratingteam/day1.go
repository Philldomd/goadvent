/*---------------------------------------
Task1: 1177     [OK]    Time: 808.8µs
Task2: 6768     [OK]    Time: 1.0427ms
---------------------------------------*/

package decoratingteam

import (
	"strconv"
	"strings"
)

type SafeCracker struct {
	Solve
}

func (safeCracker *SafeCracker) Task1(data string) string {
	ans := 0
	start := 50
	d := strings.Split(data, "\n")
	for _, line := range d {
		if len(line) == 0 {
			continue
		}
		dir := string(line[0])
		number, _ := strconv.Atoi(line[1:])
		switch dir {
		case "L":
			start = (start - number)
		case "R":
			start = (start + number)
		}
		if start%100 == 0 {ans++}
	}
  return strconv.Itoa(ans)
}

func (safeCracker *SafeCracker) Task2(data string) string {
  start := 50
	ans := 0
	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {continue}
		a, _ := strconv.Atoi(line[1:])
		ans += a / 100
		if line[0] == 'L' {
	  	a -= a*2
		}
		a = a - (a / 100 * 100)
		stop := start + a
		switch {
		case stop == 0 && a != 0:
			start = 0
			ans++
		case stop > 99:
			start = stop - 100
			ans++
		case stop < 0 && start > 0:
			start = 100 + stop
			ans++
		default:
			if stop < 0 {
				start = 100 + stop
			} else {
			  start = stop
			}
	  }
	}

	return strconv.Itoa(ans)
}