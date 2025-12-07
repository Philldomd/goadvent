/*----------------------------------------------
Task1: 44487518055      [OK]   Time: 116.398ms
Task2: 53481866137      [OK]   Time: 174.2529ms
------------------------------------------------*/

package decoratingteam

import (
  "strings"
  "strconv"
  "math"
)

type GiftShop struct {
  Solve
}

func (giftShop *GiftShop) Task1(data string) string {
  d := strings.Split(data, ",")
  ans := 0
  for _, idRange := range d {
    ids := strings.Split(idRange,"-")
    high, _ := strconv.Atoi(ids[1])
    start, _ := strconv.Atoi(ids[0])
    if len(ids[0])%2 != 0 && len(ids[1])%2 != 0 {continue}

    for start <= high {
      val := strconv.Itoa(start)
      l := len(val)/2
      if len(val)%2 != 0 {
        start = int(math.Pow(float64(10), float64(len(val))))
        continue
      }
      f_seq := val[:l]
      l_seq := val[l:]
      if f_seq == l_seq {
        ans += start
      }
      start++
    }
  }
  return strconv.Itoa(ans)
}

func (giftShop *GiftShop) iterPow(depth int, start int, inc int) int {
  if depth <= 0 {
    if depth == 0 {
      return start
    } else {
      return 0
    }
  }
  return start * int(math.Pow10(depth)) + giftShop.iterPow(depth-inc, start, inc)
}

func (giftShop *GiftShop) Task2(data string) string {
  d := strings.Split(data, ",")
  ans := 0
  for _, idRange := range d {
    ids := strings.Split(idRange,"-")
    high, _ := strconv.Atoi(ids[1])
    start, _ := strconv.Atoi(ids[0])
    found := map[int]bool{}
    len_high := len(ids[1])
    itr := 1
    for (itr*2 <= len_high) {
        len_low := len(ids[0])
        for {
          if (len_low % itr == 0 && len_low / itr >= 2) {break}
          len_low++
        }
        last_digit_low := giftShop.iterPow(len_low-1, 1, itr)
        for last_digit_low <= high {
          if (last_digit_low < int(math.Pow10(len_low)) && last_digit_low >= start) {
            if !found[last_digit_low] {
              ans += last_digit_low
              found[last_digit_low] = true
            }
          } else if (last_digit_low > int(math.Pow10(len_low))) {
            len_low += itr
            last_digit_low = giftShop.iterPow(len_low-1, 1, itr)
            continue
          }
          last_digit_low += giftShop.iterPow(len_low-itr, 1, itr)
        }
      itr++
    }
  }
  return strconv.Itoa(ans)
}