/*---------------------------------------------
Task1: 12599    [nil]   Time: 0s
Task2: 8245452805243    [nil]   Time: 8.108ms
-----------------------------------------------*/

package snowsolver

import (
  "regexp"
  "strconv"
  "strings"
)

type HauntedWasteland struct {
  Solve
}

func (hauntedWasteland HauntedWasteland) GCD(a, b int) int {
  for b != 0 {
    t := b
    b = a % b
    a = t
  }
  return a
}

// find Least Common Multiple (LCM) via GCD
func (hauntedWasteland HauntedWasteland) LCM(a, b int, integers ...int) int {
  result := a * b / hauntedWasteland.GCD(a, b)

  for i := 0; i < len(integers); i++ {
    result = hauntedWasteland.LCM(result, integers[i])
  }

  return result
}

type Node struct {
  value      string
  L          string
  IL         int
  R          string
  IR         int
  start_stop int
}

func (hauntedWasteland HauntedWasteland) walk_list(instruction []rune, nl []Node) int {
  index := 0
  node := Node{}
  for i, first := range nl {
    if first.value == "AAA" {
      node = nl[i]
    }
  }
  depth := 0
  for node.value != "ZZZ" {
    inst := instruction[index]
    if inst == 'L' {
      node = nl[node.IL]
    } else {
      node = nl[node.IR]
    }
    depth += 1
    index++
    if index == len(instruction) {
      index = 0
    }
  }
  return depth
}

func (hauntedWasteland HauntedWasteland) walk_list_stop(instruction []rune, nl []Node) int {
  loops := []int{}
  node := Node{}
  for _, n := range nl {
    if n.start_stop != 2 {
      continue
    }
    node = n
    index := 0
    depth := 0
    for node.start_stop != 1 {
      inst := instruction[index]
      if inst == 'L' {
        node = nl[node.IL]
      } else {
        node = nl[node.IR]
      }
      depth += 1
      index++
      if index == len(instruction) {
        index = 0
      }
    }
    loops = append(loops, depth)
  }
  return hauntedWasteland.LCM(loops[0], loops[1], loops...)
}

func (node Node) indexOf(list []Node) Node {
  r_node := node
  for i, n := range list {
    if n.value == node.L {
      r_node.IL = i
    }
    if n.value == node.R {
      r_node.IR = i
    }
  }
  return r_node
}

func (hauntedWasteland *HauntedWasteland) CreateNodeList(rows []string) []Node {
  nodes := []Node{}
  re := regexp.MustCompile(`([0-Z]+)`)
  for _, r := range rows {
    if len(r) == 0 {
      continue
    }
    node := Node{}
    info := strings.Split(r, " = ")
    node.value = info[0]
    LR := strings.Split(info[1], ", ")
    node.L = re.FindString(LR[0])
    node.R = re.FindString(LR[1])
    last_char := string(node.value[2])
    switch last_char {
    case "A":
      node.start_stop = 2
    case "Z":
      node.start_stop = 1
    default:
      node.start_stop = 0
    }
    nodes = append(nodes, node)
  }
  for i, node := range nodes {
    nodes[i] = node.indexOf(nodes)
  }
  return nodes
}

func (hauntedWasteland *HauntedWasteland) Task1(data string) string {
  temp := strings.Split(data, "\n\n")
  instructions := []rune(temp[0])
  rows := strings.Split(temp[1], "\n")
  sum := hauntedWasteland.walk_list(instructions, hauntedWasteland.CreateNodeList(rows))
  return strconv.Itoa(sum)
}

func (hauntedWasteland *HauntedWasteland) Task2(data string) string {
  temp := strings.Split(data, "\n\n")
  instructions := []rune(temp[0])
  rows := strings.Split(temp[1], "\n")
  sum := hauntedWasteland.walk_list_stop(instructions, hauntedWasteland.CreateNodeList(rows))
  return strconv.Itoa(sum)
}
