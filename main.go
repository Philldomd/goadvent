package main

import (
	"fmt"
  "goadvent/solvers"
	"goadvent/taskmaster"
	"strconv"
)

func main() {
	fmt.Println("Welcome to Go Advent!")
	data := taskmaster.GetPussleInput("2023", "1")
	fmt.Println("Task1: " + strconv.Itoa(solvers.GetSolver(1).Task1(data)))
	fmt.Println("Task2: " + strconv.Itoa(solvers.GetSolver(1).Task2(data)))
}
