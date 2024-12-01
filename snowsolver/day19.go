/*---------------------------------------------
Task1: 368523   [OK]    Time: 2.1638ms
Task2: 124167549767307  [OK]    Time: 3.261ms
-----------------------------------------------*/

package snowsolver

import (
  "regexp"
  "strconv"
  "strings"
)

type Aplenty struct {
  Solve
}

type functions struct {
  which    rune
  operator rune
  compare  int
  result   string
}

type part struct {
  lower int64
  upper int64
}

type partRanges struct {
  xmas  map[rune]part
  funcs string
}

func (p *partRanges) init() {
  p.xmas = make(map[rune]part)
  p.xmas['x'] = part{1, 4000}
  p.xmas['m'] = part{1, 4000}
  p.xmas['a'] = part{1, 4000}
  p.xmas['s'] = part{1, 4000}
  p.funcs = "in"
}

func (p partRanges) addRange() int64 {
  x := int64(p.xmas['x'].upper - int64(p.xmas['x'].lower-1))
  m := int64(p.xmas['m'].upper - int64(p.xmas['m'].lower-1))
  a := int64(p.xmas['a'].upper - int64(p.xmas['a'].lower-1))
  s := int64(p.xmas['s'].upper - int64(p.xmas['s'].lower-1))
  return x * m * a * s
}

var function_list map[string][]functions
var values_list []map[rune]int

func (aplenty Aplenty) findFunctions(funcs string) {
  function_list = make(map[string][]functions)
  re, _ := regexp.Compile(`{(.*)}`)
  re_num, _ := regexp.Compile(`([0-9]+):`)
  for _, row := range strings.Split(funcs, "\n") {
    f := strings.Split(row, "{")[0]
    rules := re.FindStringSubmatch(row)[1]
    func_list := []functions{}
    for _, r := range strings.Split(rules, ",") {
      fu := functions{}
      if strings.Contains(r, ":") {
        fu.which = rune(r[0])
        fu.operator = rune(r[1])
        fu.compare, _ = strconv.Atoi(re_num.FindStringSubmatch(r)[1])
        fu.result = strings.Split(r, ":")[1]
      } else {
        fu.result = r
      }
      func_list = append(func_list, fu)
    }
    function_list[f] = func_list
  }
}

func (aplenty Aplenty) findData(values string) {
  re, _ := regexp.Compile(`{(.*)}`)
  for _, row := range strings.Split(values, "\n") {
    if len(row) == 0 {
      continue
    }
    val := make(map[rune]int)
    r := re.FindStringSubmatch(row)[1]
    for _, t := range strings.Split(r, ",") {
      v := strings.Split(t, "=")
      k_v := rune(v[0][0])
      i_v, _ := strconv.Atoi(v[1])
      val[k_v] = i_v
    }
    values_list = append(values_list, val)
  }
}

func (aplenty Aplenty) interpretator(values map[rune]int, function string) bool {
  if function == "A" {
    return true
  } else if function == "R" {
    return false
  } else {
    for _, funcs := range function_list[function] {
      if funcs.compare == 0 && funcs.operator == 0 && funcs.which == 0 {
        return aplenty.interpretator(values, funcs.result)
      }
      switch funcs.operator {
      case '<':
        if values[funcs.which] < funcs.compare {
          return aplenty.interpretator(values, funcs.result)
        }
      case '>':
        if values[funcs.which] > funcs.compare {
          return aplenty.interpretator(values, funcs.result)
        }
      }
    }
  }
  return false
}

func (aplenty Aplenty) interpret() int {
  sum := 0
  for _, val := range values_list {
    if aplenty.interpretator(val, "in") {
      sum += val['x'] + val['m'] + val['a'] + val['s']
    }
  }
  return sum
}

func (aplenty *Aplenty) Task1(data string) string {
  d := strings.Split(data, "\n\n")
  aplenty.findFunctions(d[0])
  aplenty.findData(d[1])
  return strconv.Itoa(aplenty.interpret())
}

func (aplenty Aplenty) calculatePartRanges(range_list partRanges) int64 {
  sum := int64(0)
  queue := []partRanges{range_list}
  for len(queue) != 0 {
    t_range_list := queue[0]
    queue = queue[1:]
    if t_range_list.funcs == "A" {
      sum += t_range_list.addRange()
    } else if t_range_list.funcs == "R" {
      continue
    }
    for _, f := range function_list[t_range_list.funcs] {
      switch f.operator {
      case '<':
        temp_l := partRanges{}
        temp_l.xmas = make(map[rune]part)
        for key, value := range t_range_list.xmas {
          temp_l.xmas[key] = value
        }
        temp_l.funcs = f.result
        temp_l.xmas[f.which] = part{t_range_list.xmas[f.which].lower, int64(int64(f.compare) - 1)}
        t_range_list.xmas[f.which] = part{int64(f.compare), t_range_list.xmas[f.which].upper}
        queue = append(queue, temp_l)
      case '>':
        temp_h := partRanges{}
        temp_h.xmas = make(map[rune]part)
        for key, value := range t_range_list.xmas {
          temp_h.xmas[key] = value
        }
        temp_h.funcs = f.result
        temp_h.xmas[f.which] = part{int64(f.compare) + 1, t_range_list.xmas[f.which].upper}
        t_range_list.xmas[f.which] = part{t_range_list.xmas[f.which].lower, int64(f.compare)}
        queue = append(queue, temp_h)
      case 0:
        t_range_list.funcs = f.result
        queue = append(queue, t_range_list)
      }
    }
  }
  return sum
}

func (aplenty *Aplenty) Task2(data string) string {
  partRange := partRanges{}
  partRange.init()
  aplenty.findFunctions(strings.Split(data, "\n\n")[0])
  return strconv.FormatInt(aplenty.calculatePartRanges(partRange), 10)
}
