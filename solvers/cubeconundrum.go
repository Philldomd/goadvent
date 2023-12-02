package solvers

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

type CodeConundrum struct {
	Solve
}

func (codeconundrum *CodeConundrum) Task1(data string) string {
	scanner := bufio.NewScanner(strings.NewReader(data))
	sum := 0
  for scanner.Scan() {
		row := strings.Split(scanner.Text(), ":")
		possible := true
		game, _ := strconv.Atoi(strings.Fields(row[0])[1])
		for _, set := range strings.Split(row[1], ";"){
			if !possible { break }
			for _, cubes := range strings.Split(set, ","){
				color := strings.Fields(cubes)[1]
				number, _ := strconv.Atoi(strings.Fields(cubes)[0])
				switch color{
				case "red":
					if number > 12 {
            possible = false
					} 
				case "green":
					if number > 13 {
						possible = false
					}
				case "blue":
					if number > 14 {
						possible = false
					}
				default:
					log.Fatalln("Error!")
				}
				if !possible { break }
			}
		}
		if possible {	sum += game	}
	}
	return strconv.Itoa(sum)
}

func (codeconundrum *CodeConundrum) Task2(data string) string {
	scanner := bufio.NewScanner(strings.NewReader(data))
	sum := 0
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ":")
		possible := true
		red := 0
		green := 0
		blue := 0
		for _, set := range strings.Split(row[1], ";"){
			if !possible { break }
			for _, cubes := range strings.Split(set, ","){
				color := strings.Fields(cubes)[1]
				number, _ := strconv.Atoi(strings.Fields(cubes)[0])
				switch color{
				case "red":
					if number > red {
            red = number
					} 
				case "green":
					if number > green {
						green = number
					}
				case "blue":
					if number > blue {
						blue = number
					}
				default:
					log.Fatalln("Error!")
				}
			}
		}
		sum += red * green * blue
	} 
	return strconv.Itoa(sum)
}
