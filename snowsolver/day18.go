/*----------------------------------------------
Task1: 72821    [OK]   Time: 13.3774ms
Task2: 127844509405501  [OK]    Time: 1.0496ms
------------------------------------------------*/

package snowsolver

import (
  "regexp"
  "strconv"
  "strings"
)

type LavaductLagoon struct {
  Solve
}

type LavaductBorder struct {
  Dir    string
  length int
  color  string
  wall   rune
}

func (lavaductLagoon LavaductLagoon) ReadLagoonBlueprint(data string) []LavaductBorder {
  re, _ := regexp.Compile(`\((.*)\)`)
  lb := []LavaductBorder{}
  rows := strings.Split(data, "\n")
  for i, row := range rows {
    if len(row) == 0 {
      continue
    }
    temp := strings.Fields(row)
    le, _ := strconv.Atoi(temp[1])
    wall := 'x'
    if temp[0] == "R" || temp[0] == "L" {
      b_idx, a_idx := i-1, i+1
      if b_idx < 0 {
        b_idx = len(rows) - 2
      }
      if a_idx == len(rows) {
        a_idx = 0
      }
      temp_B, temp_A := strings.Fields(rows[b_idx]), strings.Fields(rows[a_idx])
      if temp_B[0] != temp_A[0] {
        wall = 'n'
      } else {
        wall = 's'
      }
    }
    lb = append(lb, LavaductBorder{temp[0], le, re.FindStringSubmatch(temp[2])[1], wall})
  }
  return lb
}

func (lavaductLagoon LavaductLagoon) CalculateBorders(lava_blueprint []LavaductBorder) (int, int, Vector) {
  max_length := 0
  max_row := 0
  min_length := 0
  min_row := 0
  temp_length, temp_row := 0, 0
  for _, dig := range lava_blueprint {
    switch dig.Dir {
    case "D":
      temp_row += dig.length
    case "U":
      temp_row -= dig.length
    case "R":
      temp_length += dig.length
    case "L":
      temp_length -= dig.length
    }
    if temp_length > max_length {
      max_length = temp_length
    }
    if temp_row > max_row {
      max_row = temp_row
    }
    if temp_length < min_length {
      min_length = temp_length
    }
    if temp_row < min_row {
      min_row = temp_row
    }
  }
  if min_length < 0 {
    min_length = -min_length
  }
  if min_row < 0 {
    min_row = -min_row
  }
  return max_row + min_row + 3, max_length + min_length + 3, Vector{min_row + 1, min_length + 1}
}

func (lavaductLagoon LavaductLagoon) DigBorders(lava_blueprint []LavaductBorder, m_row int, m_len int, start_pos Vector) [][]string {
  dig_site := [][]string{}
  for i := 0; i < m_row; i++ {
    dig_row := []string{}
    for j := 0; j < m_len; j++ {
      dig_row = append(dig_row, ".")
    }
    dig_site = append(dig_site, dig_row)
  }
  position := start_pos
  for _, lb := range lava_blueprint {
    switch lb.Dir {
    case "R":
      for len := 0; len < lb.length; len++ {
        dig_site[position.x][position.y] = "#"
        position.y++
      }
    case "D":
      for len := 0; len < lb.length; len++ {
        dig_site[position.x][position.y] = "#"
        position.x++
      }
    case "L":
      for len := 0; len < lb.length; len++ {
        dig_site[position.x][position.y] = "#"
        position.y--
      }
    case "U":
      for len := 0; len < lb.length; len++ {
        dig_site[position.x][position.y] = "#"
        position.x--
      }
    }
  }
  return dig_site
}

func (lavaductLagoon *LavaductLagoon) Task1(data string) string {
  lava_blueprint := lavaductLagoon.ReadLagoonBlueprint(data)
  m_row, m_len, start_pos := lavaductLagoon.CalculateBorders(lava_blueprint)
  dig_site := lavaductLagoon.DigBorders(lava_blueprint, m_row, m_len, start_pos)
  inside := false
  sum := 0
  for i := 0; i < len(dig_site); i++ {
    for j := 0; j < len(dig_site[0]); j++ {
      if dig_site[i][j] == "#" {
        n_u, n_d, n_l, n_r := dig_site[i-1][j], dig_site[i+1][j], dig_site[i][j-1], dig_site[i][j+1]
        if n_l == "." && n_r == "." {
          inside = !inside
          sum += 1
          continue
        }
        if n_d == "#" && n_r == "#" {
          for k := j; k < len(dig_site[0])-1; k++ {
            if dig_site[i][k+1] == "." {
              j = k
              break
            }
            sum += 1
          }
          if dig_site[i-1][j] == "#" {
            inside = !inside
          }
          if !inside {
            sum += 1
          }
        }
        if n_u == "#" && n_r == "#" {
          for k := j; k < len(dig_site[0])-1; k++ {
            if dig_site[i][k+1] == "." {
              j = k
              break
            }
            sum += 1
          }
          if dig_site[i+1][j] == "#" {
            inside = !inside
          }
          if !inside {
            sum += 1
          }
        }
      }
      if inside {
        sum += 1
      }
    }
  }
  return strconv.Itoa(sum)
}

type LavaVectors struct {
  //start   Vector
  stop Vector
  //state   int
  //turn    rune
  //counted bool
}

func (lavaductLagoon LavaductLagoon) cornerAdj(corner string) float32 {
  adj := float32(0)
  if corner == "RD" || corner == "DL" || corner == "LU" || corner == "UR" {
    adj += 0.5
  } else {
    adj -= 0.5
  }
  return adj
}

func (lavaductLagoon LavaductLagoon) createLavaVector(l_b []LavaductBorder, start_pos Vector) []LavaVectors {
  lava_vector_list := []LavaVectors{}
  l_b = append(l_b, l_b[0])
  to_corner := l_b[len(l_b)-2].Dir + l_b[0].Dir
  sp := start_pos
  for i, border := range l_b {
    if i == len(l_b)-1 {
      continue
    }
    from_corner := to_corner
    to_corner = border.Dir + l_b[i+1].Dir
    lv := LavaVectors{}
    //lv.start = sp
    lv.stop = sp
    adj := lavaductLagoon.cornerAdj(from_corner)
    adj += lavaductLagoon.cornerAdj(to_corner)
    length := border.length + int(adj)
    //lv.turn = border.wall
    switch border.Dir {
    case "R":
      lv.stop.y = lv.stop.y + length
      //lv.state = 0
    case "L":
      lv.stop.y = lv.stop.y - length
      //lv.state = 0
    case "D":
      lv.stop.x = lv.stop.x + length
      //lv.state = 1
    case "U":
      lv.stop.x = lv.stop.x - length
      //lv.state = 1
    }
    sp = lv.stop
    lava_vector_list = append(lava_vector_list, lv)
  }
  return lava_vector_list
}

func (lavaductLagoon LavaductLagoon) ReadLagoonBlueprintColor(data string) []LavaductBorder {
  re, _ := regexp.Compile(`\((.*)\)`)
  lb := []LavaductBorder{}
  rows := strings.Split(data, "\n")
  for i, row := range rows {
    if len(row) == 0 {
      continue
    }
    temp := re.FindStringSubmatch(strings.Fields(row)[2])[1]
    dir := ""
    switch temp[len(temp)-1] {
    case '0':
      dir = "R"
    case '1':
      dir = "D"
    case '2':
      dir = "L"
    case '3':
      dir = "U"
    }
    le, _ := strconv.ParseInt(temp[1:len(temp)-1], 16, 64)
    wall := 'x'
    if dir == "R" || dir == "L" {
      b_idx, a_idx := i-1, i+1
      if b_idx < 0 {
        b_idx = len(rows) - 2
      }
      if a_idx == len(rows) {
        a_idx = 0
      }
      temp_B, temp_A := re.FindStringSubmatch(strings.Fields(rows[b_idx])[2])[1], re.FindStringSubmatch(strings.Fields(rows[a_idx])[2])[1]
      if temp_B[len(temp_B)-1] != temp_A[len(temp_A)-1] {
        wall = 'n'
      } else {
        wall = 's'
      }
    }
    lb = append(lb, LavaductBorder{dir, int(le), temp, wall})
  }
  return lb
}

func (lavaductLagoon *LavaductLagoon) Task2(data string) string {
  lava_blueprint := lavaductLagoon.ReadLagoonBlueprintColor(data)
  _, _, start_pos := lavaductLagoon.CalculateBorders(lava_blueprint)
  lava_vector_list := lavaductLagoon.createLavaVector(lava_blueprint, start_pos)
  sum := int64(0)
  temp_sum := float64(0)
  for i := 0; i < len(lava_vector_list); i++ {
    this, after := i, i+1
    if i == len(lava_vector_list)-1 {
      after = 0
    }
    P1 := lava_vector_list[this].stop
    P2 := lava_vector_list[after].stop
    temp := float64((P1.x * P2.y) - (P1.y * P2.x))
    temp_sum += temp
  }
  if temp_sum < 0 {
    temp_sum = -temp_sum
  }
  temp_sum = temp_sum * 0.5
  sum = int64(temp_sum)
  return strconv.Itoa(int(sum))
}
