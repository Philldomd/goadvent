package main

import (
	"fmt"
	"goadvent/solvers"
	"goadvent/taskmaster"
	"strconv"
)

var year string = "2023"
var day string = "2"
/*var temp string = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`*/

func main() {
	iday, _ := strconv.Atoi(day)
	fmt.Println("Welcome to Go Advent!")
	data := taskmaster.GetPussleInput(year, day)
	task1 := solvers.GetSolver(iday).Task1(data)
	task2 := solvers.GetSolver(iday).Task2(data)
	fmt.Printf("Task1: %s\t[%s]\n", task1, taskmaster.PostAnswer(year, day, "1", task1))
	fmt.Printf("Task2: %s\t[%s]\n", task2, taskmaster.PostAnswer(year, day, "2", task2))
}
