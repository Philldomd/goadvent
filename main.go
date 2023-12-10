package main

import (
	"fmt"
	solvers "goadvent/solvers/2023"
	"goadvent/taskmaster"
	"os"
	"strconv"
)

var year string = "2023"
var day string = "9"
var task string = "0"
var test bool = false

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
		data := ``
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
