/*---------------------------------------------------
Task1: 899848294        [OK]        Time: 36.248ms
Task2: 247454898168563  [OK]        Time: 80.6327ms
-----------------------------------------------------*/

package snowsolver

import (
  "strconv"
  "strings"
)

type PulsePropagation struct {
  Solve
}

type chip struct {
  c_type rune
  //flip flop
  state bool
  //conjuntion
  inputs       map[string]int
  destinations []string
}

var highs int
var lows int

func flip(c *chip, from string, sig int) int {
  if c.c_type == '%' {
    return c.flip_signal(sig)
  } else {
    return c.conj_signal(from, sig)
  }
}

func (f *chip) flip_signal(sig int) int {
  if sig == 1 {
    f.state = !f.state
    if f.state {
      return 2
    } else {
      return 1
    }
  }
  return -1
}

func (c *chip) conj_signal(from string, sig int) int {
  c.inputs[from] = sig
  ret := 1
  for _, val := range c.inputs {
    if val == 1 {
      ret = 2
      break
    }
  }
  return ret
}

func (pulse *PulsePropagation) Task1(data string) string {
  highs, lows = 0, 0
  high := []int{}
  low := []int{}
  sequence_length := 0
  machine := make(map[string]chip)
  var broadcaster []string
  for _, row := range strings.Split(data, "\n") {
    if len(row) == 0 {
      continue
    }
    c_type := row[0]
    sp_row := strings.Split(row, "->")
    if c_type == 'b' {
      broadcaster = strings.Split(sp_row[1], ",")
      for i := range broadcaster {
        broadcaster[i] = strings.TrimLeft(broadcaster[i], " ")
      }
      continue
    }
    name := strings.TrimRight(sp_row[0][1:], " ")
    dest := strings.Split(sp_row[1], ",")
    for i := range dest {
      dest[i] = strings.TrimLeft(dest[i], " ")
    }
    machine[name] = chip{c_type: rune(c_type), state: false, inputs: map[string]int{}, destinations: dest}
  }
  for key, value := range machine {
    if value.c_type == '%' {
      for _, dest := range value.destinations {
        if machine[dest].c_type == '&' {
          machine[dest].inputs[key] = 1
        }
      }
    }
  }
  for {
    lows++
    queue := []string{}
    from := []string{}
    signal := []int{}
    for _, bc := range broadcaster {
      queue = append(queue, bc)
      from = append(from, "broadcaster")
      signal = append(signal, 1)
      lows++
    }
    for len(queue) != 0 {
      pop := queue[0]
      queue = queue[1:]
      fr := from[0]
      from = from[1:]
      sig := signal[0]
      signal = signal[1:]
      c_chip := machine[pop]
      c_sig := flip(&c_chip, fr, sig)
      machine[pop] = c_chip
      if c_sig > 0 {
        for _, dest := range machine[pop].destinations {
          if c_sig == 1 {
            lows++
          } else {
            highs++
          }
          if dest != "rx" {
            queue = append(queue, dest)
            from = append(from, pop)
            signal = append(signal, c_sig)
          }
        }
      }
    }
    stop := true
    for _, value := range machine {
      if value.state {
        stop = false
        break
      }
    }
    high = append(high, highs)
    low = append(low, lows)
    sequence_length++
    if stop || sequence_length == 1000 {
      break
    }
  }
  mod := 1000 % sequence_length
  if mod != 0 {
    highs += high[mod]
    lows += low[mod]
  }
  multi := (1000 - mod) / sequence_length
  high_ans := highs * multi
  low_ans := lows * multi
  ans := high_ans * low_ans
  return strconv.Itoa(ans)
}

func (pulse *PulsePropagation) Task2(data string) string {
  highs, lows = 0, 0
  buttons_pressed := []int{}
  pressed := 0
  machine := make(map[string]chip)
  var broadcaster []string
  for _, row := range strings.Split(data, "\n") {
    if len(row) == 0 {
      continue
    }
    c_type := row[0]
    sp_row := strings.Split(row, "->")
    if c_type == 'b' {
      broadcaster = strings.Split(sp_row[1], ",")
      for i := range broadcaster {
        broadcaster[i] = strings.TrimLeft(broadcaster[i], " ")
      }
      continue
    }
    name := strings.TrimRight(sp_row[0][1:], " ")
    dest := strings.Split(sp_row[1], ",")
    for i := range dest {
      dest[i] = strings.TrimLeft(dest[i], " ")
    }
    machine[name] = chip{c_type: rune(c_type), state: false, inputs: map[string]int{}, destinations: dest}
  }
  for key, value := range machine {
    if value.c_type == '%' {
      for _, dest := range value.destinations {
        if machine[dest].c_type == '&' {
          machine[dest].inputs[key] = 1
        }
      }
    }
  }
  for _, bc := range broadcaster {
    stop := false
    for {
      lows++
      pressed++
      queue := []string{}
      from := []string{}
      signal := []int{}

      queue = append(queue, bc)
      from = append(from, "broadcaster")
      signal = append(signal, 1)
      lows++
      for len(queue) != 0 {
        pop := queue[0]
        queue = queue[1:]
        fr := from[0]
        from = from[1:]
        sig := signal[0]
        signal = signal[1:]
        if pop == "rx" {
          if sig == 1 {
            stop = true
            break
          } else {
            continue
          }
        }
        c_chip := machine[pop]
        c_sig := flip(&c_chip, fr, sig)
        machine[pop] = c_chip
        if c_sig > 0 {
          for _, dest := range machine[pop].destinations {
            if c_sig == 1 {
              lows++
            } else {
              highs++
            }
            queue = append(queue, dest)
            from = append(from, pop)
            signal = append(signal, c_sig)
          }
        }
      }
      if stop {
        break
      }
    }
    buttons_pressed = append(buttons_pressed, pressed)
    pressed = 0
  }
  //my gcd was 1 so lcm is just multiply all pressed buttons
  val := int64(1)
  for i := range buttons_pressed {
    val *= int64(buttons_pressed[i])
  }
  return strconv.FormatInt(val, 10)
}
