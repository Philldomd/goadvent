/*-------------------------------------
Task1: 55386    [nil]   Time: 3.765ms
Task2: 54824    [nil]   Time: 3.4017ms
---------------------------------------*/

package snowsolver

import (
  "bufio"
  "regexp"
  "strconv"
  "strings"
)

type Trebuchet struct {
  Solve
}

type match struct {
  index  int
  number int
}

func (trebuchet *Trebuchet) Task1(data string) string {
  scanner := bufio.NewScanner(strings.NewReader(data))
  first := regexp.MustCompile("[0-9]")
  last := regexp.MustCompile("([0-9])[^0-9]*$")
  sum := 0
  for scanner.Scan() {
    f := first.FindString(scanner.Text())
    l := last.FindStringSubmatch(scanner.Text())[1]
    number, _ := strconv.Atoi(f + l)
    sum += number
  }
  return strconv.Itoa(sum)
}

func (trebuchet *Trebuchet) Task2(data string) string {
  sum := 0
  scanner := bufio.NewScanner(strings.NewReader(data))
  num := []string{"0", "zero", "1", "one", "2", "two", "3", "three", "4", "four", "5", "five", "6", "six", "7", "seven", "8", "eight", "9", "nine"}
  for scanner.Scan() {
    first := match{4098, 0}
    last := match{-1, 0}
    for i, n := range num {
      index := strings.Index(scanner.Text(), n)
      if index != -1 {
        if first.index > index {
          first.index = index
          first.number = i / 2
        }
      }
      last_index := strings.LastIndex(scanner.Text(), n)
      if last_index != -1 {
        if last.index < last_index {
          last.index = last_index
          last.number = i / 2
        }
      }
    }
    sum += first.number*10 + last.number
  }
  return strconv.Itoa(sum)
}
