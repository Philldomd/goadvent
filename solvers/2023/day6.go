/*---------------------------------------
Task1: 840336   [nil]   Time: 0s
Task2: 41382569 [nil]   Time: 46.9754ms
-----------------------------------------*/

package solvers

import (
	"regexp"
	"strconv"
	"strings"
)

type WaitForIt struct {
	Solve
}

func (waitForIt *WaitForIt) Task1(data string) string {
	rows := strings.Split(data, "\n")
	time := []int{}
	distance := []int{}
	winning := []int{}
	for _, r := range rows {
		switch strings.Split(r, ":")[0] {
		case "Time":
			fields := strings.Fields(strings.Split(r, ":")[1])
			for i := 0; i < len(fields); i++ {
				it, _ := strconv.Atoi(fields[i])
				time = append(time, it)
			}
		case "Distance":
			fields := strings.Fields(strings.Split(r, ":")[1])
			for i := 0; i < len(fields); i++ {
				id, _ := strconv.Atoi(fields[i])
				distance = append(distance, id)
			}
		}
	}
	for i := 0; i < len(time); i++ {
		w := 0
		l := 0
		even := false
		if time[i]%2 == 1 {
			l = (time[i] / 2) + 1
		} else {
			even = true
			l = time[i] / 2
		}
		for j := 0; j < l; j++ {
			if j*(time[i]-j) > distance[i] {
				w += 2
			}
		}
		if even {
			if l*(time[i]-l) > distance[i] {
				w++
			}
		}
		winning = append(winning, w)
	}
	w := 1
	for _, win := range winning {
		w *= win
	}
	return strconv.Itoa(w)
}

func (waitForIt *WaitForIt) Task2(data string) string {
	re, _ := regexp.Compile(`[[:blank:]]`)
	data = re.ReplaceAllString(data, "")
	return waitForIt.Task1(data)
}
