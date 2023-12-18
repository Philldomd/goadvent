package main

import (
	"fmt"
	solvers "goadvent/solvers/2023"
	"goadvent/taskmaster"
	"os"
	"strconv"
	"time"
)

var year string = "2023"
var day string = "17"
var task string = "1"
var test bool = true
var send bool = false

func main() {
	for i, arg := range os.Args {
		switch arg {
		case "-d":
			day = os.Args[i+1]
		case "-y":
			year = os.Args[i+1]
		case "-t":
			task = os.Args[i+1]
		case "-f":
			test = true
		case "-s":
			send = true
		default:
			continue
		}
	}
	iday, _ := strconv.Atoi(day)
	fmt.Printf("Welcome to Go Advent! It's day %s, year %s!\n", day, year)
	if !test {
		data := taskmaster.GetPussleInput(year, day)
		if task == "1" || task == "0" {
			start := time.Now()
			task1 := solvers.GetSolver(iday).Task1(data)
			stop := time.Since(start)
			if send {
				fmt.Printf("Task1: %s\t[%s]\tTime: %s\n", task1, taskmaster.PostAnswer(year, day, "1", task1), stop)
			} else {
				fmt.Printf("Task1: %s\t[%s]\tTime: %s\n", task1, "nil", stop)
			}
		}
	  if task == "2" || task == "0" {
			start := time.Now()
			task2 := solvers.GetSolver(iday).Task2(data)
			stop := time.Since(start)
			if send {
			  fmt.Printf("Task2: %s\t[%s]\tTime: %s\n", task2, taskmaster.PostAnswer(year, day, "2", task2), stop)
			} else {
			  fmt.Printf("Task2: %s\t[%s]\tTime: %s\n", task2, "nil", stop)	
			}
		}
	} else {
		data := `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533
`
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
