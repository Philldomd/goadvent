/*-------------------------------------
Task1: 2430334  [OK]    Time: 981.6µs
Task2: 28786472 [OK]    Time: 690.3µs
---------------------------------------*/

package missinghistorian

import (
  "math"
  "sort"
  "strconv"
  "strings"
)

type Hysteria struct {
  Solve
}

func (Hysteria *Hysteria) Task1(data string) string {
  var ret = 0
  rows := strings.Split(data, "\n")
  var left []int
  var right []int
  for _, r := range rows {
    if len(r) == 0 {
      continue
    }
    ileft, _ := strconv.Atoi(strings.Split(r,"   ")[0])
    iright, _ := strconv.Atoi(strings.Split(r, "   ")[1])
    left = append(left, ileft)
    right = append(right, iright)
  }
  sort.Ints(left)
  sort.Ints(right)
  for i, lval := range left {
    ret += int(math.Abs(float64(lval - right[i])))
  }
  return strconv.Itoa(ret)
}

func (Hysteria *Hysteria) Task2(data string) string {
  var ret = 0
  rows := strings.Split(data, "\n")
  var left = make(map[int]int)
  var right = make(map[int]int)
  for _, r := range rows {
    if len(r) == 0 {
      continue
    }
    lval, _ := strconv.Atoi(strings.Split(r, "   ")[0])
    rval, _ := strconv.Atoi(strings.Split(r, "   ")[1])
    if val, ok := left[lval]; ok {
      left[lval] = val + 1
    } else {
      left[lval] = 1
    }
    if val, ok := right[rval]; ok {
      right[rval] = val + 1
    } else {
      right[rval] = 1
    }
  }
  for key, val := range left {
    ret += key * val * right[key]
  }
  return strconv.Itoa(ret)
}
