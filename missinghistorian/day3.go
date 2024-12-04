/*-------------------------------------
Task1: 188741603 [OK]    Time: 0s
Task2: 67269798  [OK]    Time: 2.2557ms
---------------------------------------*/

package missinghistorian

import (
	"regexp"
	"strconv"
	"strings"
)

type Mullitover struct {
	Solve
}

func (Mullitover * Mullitover) Task1(data string) string {
	ret := 0
	var re = regexp.MustCompile(`mul\((\d+,\d+)\)`)
	found := re.FindAllStringSubmatch(data, -1)
	for _, val := range found {
		r_l := strings.Split(val[1], ",")
		r, _ := strconv.Atoi(r_l[0])
		l, _ := strconv.Atoi(r_l[1])
		ret += r * l
	}
	return strconv.Itoa(ret)
}

func (Mullitover *Mullitover) Task2(data string) string {
	ret := 0
	var re = regexp.MustCompile(`mul\((\d+,\d+)\)`)
	var redo = regexp.MustCompile(`do\(\)`)
	var redont = regexp.MustCompile(`don't\(\)`)
	found_d := re.FindAllStringSubmatchIndex(data, -1)
	found := re.FindAllStringSubmatch(data, -1)
	dos := redo.FindAllStringIndex(data, -1)
	donts := redont.FindAllStringIndex(data, -1)
	donts_counter := 0
	dos_counter := 0
	count := true
	for i, val := range found {
		if donts_counter < len(donts) && donts[donts_counter][0] < found_d[i][0] {
			donts_counter += 1
			count = false
		}
		if dos_counter < len(dos) && dos[dos_counter][0] < found_d[i][0] {
			dos_counter += 1
			count = true
		}
		if count {
			r_l := strings.Split(val[1], ",")
			r, _ := strconv.Atoi(r_l[0])
			l, _ := strconv.Atoi(r_l[1])
			ret += r * l
		}
	}
	return strconv.Itoa(ret)
}