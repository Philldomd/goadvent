/*-------------------------------------
Task1: 5166     [OK]    Time: 29.3182ms
Task2: 4679     [OK]    Time: 796.8082ms
---------------------------------------*/

package missinghistorian

import (
	"strconv"
	"strings"
)

type PrintQueue struct {
	Solve
}

func findIndex(arr []string, value string) (int, bool) {
	for x, v := range arr {
		if v == value {
      return x, true
		}
	}
	return 0, false
}

func (printQueue *PrintQueue) Task1(data string) string {
	var rules = make(map[string][]string)
	ret := 0
	failed := false
	rulesInput := strings.Split(data, "\n\n")
	for _, rule := range strings.Split(rulesInput[0], "\n") {
    r := strings.Split(rule, "|")
		rules[r[0]] = append(rules[r[0]], r[1])
	}
	for _, update := range strings.Split(rulesInput[1], "\n") {
		if len(update) == 0 {
			continue
		}
		numbers := strings.Split(update, ",")
		for pos, number := range numbers {
			if val, ok := rules[number]; ok {
				for _, values := range val {
					if idx, ok := findIndex(strings.Split(update, ","), values); ok {
						if idx < pos {
							failed = true
							break
						}
					}
				}
			}
      if failed {
				break
			}
		}
		if failed {
			failed = false
			continue
		} 
		middle := len(numbers) / 2
		r, _ := strconv.Atoi(numbers[middle])
		ret += r
	}
	return strconv.Itoa(ret)
}

func solveByRule(rules map[string][]string, values []string) ([]string, bool) {
	var ret []string
	copy(ret, values)
	for pos, number := range values {
		if val, ok := rules[number]; ok {
			for _, value := range val {
				if idx, ok := findIndex(values, value); ok {
					if idx < pos {
						for x := idx; x < pos; x++ {
							temp := values[x]
							values[x] = values[x+1]
							values[x+1] = temp
						}
						return values, false
					}
				}
			}
		}
	}
	return ret, true
}

func (printQueue *PrintQueue) Task2(data string) string {
	var rules = make(map[string][]string)
	ret := 0
	failed := false
	rulesInput := strings.Split(data, "\n\n")
	var failedRows [][]string
	for _, rule := range strings.Split(rulesInput[0], "\n") {
    r := strings.Split(rule, "|")
		rules[r[0]] = append(rules[r[0]], r[1])
	}
	for _, update := range strings.Split(rulesInput[1], "\n") {
		if len(update) == 0 {
			continue
		}
		numbers := strings.Split(update, ",")
		for pos, number := range numbers {
			if val, ok := rules[number]; ok {
				for _, values := range val {
					if idx, ok := findIndex(numbers, values); ok {
						if idx < pos {
							failed = true
							break
						}
					}
				}
			}
      if failed {
				failedRows = append(failedRows, numbers)
				break
			}
		}
		if failed {
			failed = false
			continue
		}
	}
	for x := range failedRows {
		for {
			if _, ok := solveByRule(rules, failedRows[x]); ok {
				middle := len(failedRows[x]) / 2
				r, _ := strconv.Atoi(failedRows[x][middle])
				ret += r
				break
			}
		}
	}
	return strconv.Itoa(ret)
}