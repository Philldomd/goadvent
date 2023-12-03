package solvers

import (
	"strconv"
	"strings"
	"unicode"
)

type GearRatios struct {
	Solve
}

func rowCheck(lines []string, row int, column int, substring string) int {
	upperBound := row - 1
	lowerBound := row + 1
	leftBound := column - len(substring)
	rightBound := column + 1
	for i := upperBound; i <= lowerBound; i++ {
		if i < 0 || i == len(lines)-1 {
			continue
		}
		for j := leftBound; j <= rightBound; j++ {
			if j < 0 || j == len(lines[row]) {
				continue
			}
			r := rune(lines[i][j])
			if !unicode.IsDigit(r) && r != '.' {
				n, _ := strconv.Atoi(substring)
				return n
			}
		}
	}
	return 0
}

func gearRatioCheck(lines []string, row int, column int) int {
	numbers := []int{}
	upperBound := row - 1
	lowerBound := row + 1
	leftBound := column - 1
	rightBound := column + 1
	n := ""
	for i := upperBound; i <= lowerBound; i++ {
		if i < 0 || i == len(lines)-1 {
			continue
		}
		j := leftBound
		if j < 0 {
			j = 0
		}
		r := rune(lines[i][j])
		L := leftBound
		R := rightBound
		for unicode.IsDigit(r) {
			if L > 0 {
				n = string(r) + n
				L--
				r = rune(lines[i][L])
			} else if L == 0 {
				n = string(r) + n
				break
			} else {
				break
			}
		}
		if unicode.IsDigit(rune(lines[i][column])) {
			if len(n) > 1 {
				n = n + string(lines[i][column]) + "|"
			} else {
				n = "|" + n + string(lines[i][column])
			}
		} else {
			n = n + "|"
		}
		if column + 1 < len(lines[row]) {
			r = rune(lines[i][column+1])
			for unicode.IsDigit(r) {
				if R < len(lines[row])-1 {
					n = n + string(r)
					R++
					r = rune(lines[i][R])
				} else if R == len(lines[row])-1 {
					n = n + string(r)
					break
				} else {
					break
				}
			}
		}
		if len(n) > 0 {
			num := strings.Split(n, "|")
			n = ""
			for _, i := range num {
				if i != "" {
					x, _ := strconv.Atoi(i)
					numbers = append(numbers, x)
				}
			}
		}
	}
	if len(numbers) == 2 {
		sum := numbers[0] * numbers[1]
		return sum
	}
	return 0
}

func (gearratios *GearRatios) Task1(data string) string {
	sum := 0
	var lines []string = strings.Split(data, "\n")
	substring := ""
	ans := 0
	for i, row := range lines {
		for j := 0; j < len(row); j++ {
			r := rune(row[j])
			if unicode.IsDigit(r) {
				substring = substring + string(r)
				j++
				for _, n := range row[j:] {
					if !unicode.IsDigit(n) {
						ans += rowCheck(lines, i, j-1, substring)
						break
					} else {
						j++
						substring = substring + string(n)
						if j == len(row) {
							ans += rowCheck(lines, i, j-1, substring)
						}
					}
				}
				substring = ""
			}
		}
		if ans > 0 {
			sum += ans
			ans = 0
		}
	}
	return strconv.Itoa(sum)
}

func (gearratios *GearRatios) Task2(data string) string {
	sum := 0
	var lines []string = strings.Split(data, "\n")
	symbol := "*"
	for i, row := range lines {
		for j := 0; j < len(row); j++ {
			r := string(row[j])
			if r == symbol {
				sum += gearRatioCheck(lines, i, j)
			}
		}
	}
	return strconv.Itoa(sum)
}
