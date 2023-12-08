package solvers

import (
	"sort"
	"strconv"
	"strings"
)

type CamelCards struct {
	Solve
}

type Found struct {
	key int 
	value int
}

type Hand struct {
	cards   []int
	points  int
	value   int
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func calcHand(h *Hand) {
	found := []Found{}
	fo := false
	for i, c := range h.cards {
		for j, f := range found {
			if f.key == c {
				found[j].value += 1
				fo = true
			}
		}
		for j := i + 1; j < len(h.cards); j++ {
			if c == h.cards[j] && !fo {
				f := Found{}
				f.key = c
				f.value = 1
				found = append(found, f)
				break
			}
		}
		fo = false
	}
	if len(found) == 2 {
		if found[0].value == 3 || found[1].value == 3 {
			h.value = found[0].value + found[1].value
		} else {
			h.value = 3
		}
	} else if len(found) == 1 {
    if found[0].value == 3 {
			h.value = 4
		} else if found[0].value > 3 {
			h.value = found[0].value + 2
		} else {
			h.value = found[0].value
		}
	} else {
		h.value = 1
	}
}

func calcHandJoker(h *Hand) {
	joker := 0
	found := []Found{}
	fo := false
	for i, c := range h.cards {
		for j, f := range found {
			if f.key == c {
				found[j].value += 1
				fo = true
			}
		}
		if c == 0 {
			joker++
		} else {
			for j := i + 1; j < len(h.cards); j++ {
				if c == h.cards[j] && !fo {
					f := Found{}
					f.key = c
					f.value = 1
					found = append(found, f)
					break
				}
			}
		}
		if len(found) == 2 {
			if found[0].value == 3 || found[1].value == 3 {
				h.value = found[0].value + found[1].value
			} else {
				h.value = 3
			}
		} else if len(found) == 1 {
			if found[0].value == 3 {
				h.value = 4
			} else if found[0].value > 3 {
				h.value = found[0].value + 2
			} else {
				h.value = found[0].value
			}
		} else {
			h.value = 1
		}
		fo = false
	}
	for joker > 0 {
		if h.value > 0 {
			if h.value == 1 {
				h.value += 1
			} else if h.value < 5 {
        h.value += 2
			} else {
				h.value += 1
			}
      
		} else {
			if joker > 4 {
				h.value = 7
				joker -= 5
			} else {
				h.value = 2
			}
		}
		joker--
	}
}

func sortCards(hands []Hand) []Hand {
	sort.Slice(hands[:], func(a, b int) bool {
		return hands[a].value > hands[b].value
	})
	start := 0
	end := 0
	for i := 7; i > 0; i-- {
    for j := start; j < len(hands); j++ {
			if hands[j].value < i {
				end = j
				break
			} else if  j == len(hands) - 1 {
				end = j+1
			}
		}
		temp_hand := hands[start:end]
		sort.Slice(temp_hand[:], func(a, b int) bool {
			for i := 0; i < len(temp_hand[a].cards); i++ {
				if temp_hand[a].cards[i] == temp_hand[b].cards[i] {
					continue
				}
				return temp_hand[a].cards[i] > temp_hand[b].cards[i]
			}
			return false
		})
		start = end
	}
	return hands
}

func poker(data string, c_id []string, joker bool) string {
	row := strings.Split(data, "\n")
	hands := []Hand{}
	sum := 0
	max_rank := len(row) - 1
	for _, r := range row {
		hand := Hand{}
		if len(r) > 0 {
			cards := strings.Split(strings.Fields(r)[0], "")
			for _, c := range cards {
				ic := indexOf(c, c_id)
				hand.cards = append(hand.cards, ic)
			}
			hand.points, _ = strconv.Atoi(strings.Fields(r)[1])
			if joker {
				calcHandJoker(&hand)
			} else {
				calcHand(&hand)
			}
			hands = append(hands, hand)
		}
	}
	hands = sortCards(hands)
	for _, h := range hands {
		sum += h.points * max_rank
	  max_rank--	
	}
	return strconv.Itoa(sum)
}

func (camelCards *CamelCards) Task1(data string) string {
	c_id := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	return poker(data, c_id, false)
}

func (camelCards *CamelCards) Task2(data string) string {
	c_id := []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}
	return poker(data, c_id, true)
}
