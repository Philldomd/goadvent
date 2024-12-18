package main

import (
  "fmt"
  snowsolvers "goadvent/snowsolver"
  missinghistorian "goadvent/missinghistorian"
  "goadvent/taskmaster"
  "os"
  "strconv"
  "time"
)

var year string = "0"
var day string = "0"
var task string = "0"
var test bool = false
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
  if year == "2023" {
    if !test {
      data := taskmaster.GetPussleInput(year, day)
      if task == "1" || task == "0" {
        start := time.Now()
        task1 := snowsolvers.GetSolver(iday).Task1(data)
        stop := time.Since(start)
        if send {
          fmt.Printf("Task1: %s\t[%s]\tTime: %s\n", task1, taskmaster.PostAnswer(year, day, "1", task1), stop)
        } else {
          fmt.Printf("Task1: %s\t[%s]\tTime: %s\n", task1, "nil", stop)
        }
      }
      if task == "2" || task == "0" {
        start := time.Now()
        task2 := snowsolvers.GetSolver(iday).Task2(data)
        stop := time.Since(start)
        if send {
          fmt.Printf("Task2: %s\t[%s]\tTime: %s\n", task2, taskmaster.PostAnswer(year, day, "2", task2), stop)
        } else {
          fmt.Printf("Task2: %s\t[%s]\tTime: %s\n", task2, "nil", stop)  
        }
      }
    } else {
      data := ``
      if task == "1" || task == "0" {
        task1 := snowsolvers.GetSolver(iday).Task1(data)
        fmt.Printf("Task1: %s\t[%s]\n", task1, "nil")
      }
      if task == "2" || task == "0" {
        task2 := snowsolvers.GetSolver(iday).Task2(data)
        fmt.Printf("Task2: %s\t[%s]\n", task2, "nil")
      }
    }
  } else if year == "2024" {
    if !test {
      data := taskmaster.GetPussleInput(year, day)
      if task == "1" || task == "0" {
        start := time.Now()
        task1 := missinghistorian.GetSolver(iday).Task1(data)
        stop := time.Since(start)
        if send {
          fmt.Printf("Task1: %s\t[%s]\tTime: %s\n", task1, taskmaster.PostAnswer(year, day, "1", task1), stop)
        } else {
          fmt.Printf("Task1: %s\t[%s]\tTime: %s\n", task1, "nil", stop)
        }
      }
      if task == "2" || task == "0" {
        start := time.Now()
        task2 := missinghistorian.GetSolver(iday).Task2(data)
        stop := time.Since(start)
        if send {
          fmt.Printf("Task2: %s\t[%s]\tTime: %s\n", task2, taskmaster.PostAnswer(year, day, "2", task2), stop)
        } else {
          fmt.Printf("Task2: %s\t[%s]\tTime: %s\n", task2, "nil", stop)  
        }
      }
    } else {
      data := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`
      if task == "1" || task == "0" {
        task1 := missinghistorian.GetSolver(iday).Task1(data)
        fmt.Printf("Task1: %s\t[%s]\n", task1, "nil")
      }
      if task == "2" || task == "0" {
        task2 := missinghistorian.GetSolver(iday).Task2(data)
        fmt.Printf("Task2: %s\t[%s]\n", task2, "nil")
      }
    }
  }
}
