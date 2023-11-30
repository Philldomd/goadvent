package main

import (
	"fmt"
	"goadvent/taskmaster"
)

func main() {
  fmt.Println("Welcome to Go Advent!\nThis journey will start tomorrow")
  fmt.Println(taskmaster.GetPussleInput("2022", "1"))
}