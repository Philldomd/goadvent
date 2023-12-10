package main

import (
	"fmt"
	solvers "goadvent/solvers/2023"
	"goadvent/taskmaster"
	"os"
	"strconv"
)

var year string = "2023"
var day string = "10"
var task string = "2"
var test bool = true

func main() {
	for i, arg := range os.Args {
		switch arg {
		case "-d":
			day = os.Args[i+1]
		case "-y":
			year = os.Args[i+1]
		case "-t":
			task = os.Args[i+1]
		default:
			continue
		}
	}
	iday, _ := strconv.Atoi(day)
	fmt.Printf("Welcome to Go Advent! It's day %s, year %s!\n", day, year)
	if !test {
		data := taskmaster.GetPussleInput(year, day)
		if task == "1" || task == "0" {
			task1 := solvers.GetSolver(iday).Task1(data)
			fmt.Printf("Task1: %s\t[%s]\n", task1, taskmaster.PostAnswer(year, day, "1", task1))
		}
		if task == "2" || task == "0" {
			task2 := solvers.GetSolver(iday).Task2(data)
			fmt.Printf("Task2: %s\t[%s]\n", task2, taskmaster.PostAnswer(year, day, "2", task2))
		}
	} else {
		data := `
.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`
		if task == "1" || task == "0" {
			task1 := solvers.GetSolver(iday).Task1(data)
			fmt.Printf("Task1: %s\t[%s]\n", task1, "nil")
		}
		if task == "2" || task == "0" {
			task2 := solvers.GetSolver(iday).Task2(data)
			fmt.Printf("Task2: %s\t[%s]\n", task2, "nil")
		}
	}
}
