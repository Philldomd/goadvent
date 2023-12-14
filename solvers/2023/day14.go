/*-----------------------------------------
Task1: 109345   [OK]    Time: 810.8µs
Task2: 112452   [OK]    Time: 19.4415061s
-------------------------------------------*/

package solvers

import (
	"strconv"
	"strings"
)

type ParabolicReflectorDish struct {
	Solve
}

type Rock struct {
	v       Vector
	stone   rune
	score   int
	at_rest bool
}

// make map over at_rest stones
// make loop inside this function for each stone instead of outside
func (parabolicReflectorDish ParabolicReflectorDish) RocknRoll(height int, width int, rock_list []Rock, idx int, dir Vector) Vector {
	pos, limit := 0, 0
	score := height
	_x, _y := rock_list[idx].v.x, rock_list[idx].v.y
	switch dir {
	case Vector{-1, 0}:
		pos = _x
	case Vector{0, -1}:
		pos = _y
	case Vector{1, 0}:
		pos = _x
		limit = height
	case Vector{0, 1}:
		pos = _y
		limit = width
	}
	if pos == limit {
		rock_list[idx].at_rest = true
		rock_list[idx].score = score - rock_list[idx].v.x
		return Vector{_x, _y}
	}
	for !rock_list[idx].at_rest {
		moving_rock := rock_list[idx]
		moving_rock.v.x += dir.x
		moving_rock.v.y += dir.y
		if moving_rock.v.x < 0 || moving_rock.v.y < 0 {
			rock_list[idx].at_rest = true
			rock_list[idx].score = score - rock_list[idx].v.x
			break
		}
		if moving_rock.v.x == height || moving_rock.v.y == width {
			rock_list[idx].at_rest = true
			rock_list[idx].score = score - rock_list[idx].v.x
			break
		}
		for i, rock := range rock_list {
			if rock.v == moving_rock.v {
				if !rock.at_rest {
					rock_roll := parabolicReflectorDish.RocknRoll(height, width, rock_list, i, dir)
					rock_list[idx].at_rest = true
					rock_list[idx].v.x = rock_roll.x - dir.x
					rock_list[idx].v.y = rock_roll.y - dir.y
					rock_list[idx].score = score - rock_list[idx].v.x
					break
				} else {
					rock_list[idx].at_rest = true
					rock_list[idx].score = score - rock_list[idx].v.x
					break
				}
			}
		}
		if !rock_list[idx].at_rest {
			rock_list[idx].v = moving_rock.v
		}
	}
	return rock_list[idx].v
}

func (parabolicReflectorDish ParabolicReflectorDish) RockRollNorth(picture [][]rune, rock_list []Rock) int {
	score := len(picture)
	sum := 0
	for _, rock := range rock_list {
		if rock.v.x == 0 {
			rock.score = score
		}
		at_rest := false
		for !at_rest {
			if rock.v.x == 0 {
				rock.score = score
				at_rest = true
			} else {
				switch picture[rock.v.x-1][rock.v.y] {
				case '.':
					picture[rock.v.x][rock.v.y] = '.'
					picture[rock.v.x-1][rock.v.y] = 'O'
					rock.v.x -= 1
				case '#':
					rock.score = score - rock.v.x
					at_rest = true
				case 'O':
					rock.score = score - rock.v.x
					at_rest = true
				}
			}
		}
		sum += rock.score
	}
	return sum
}

func (parabolicReflectorDish ParabolicReflectorDish) CollectRocks(data string, all bool) ([][]rune, []Rock) {
	picture := [][]rune{}
	rock_list := []Rock{}
	for i, row := range strings.Split(data, "\n") {
		if len(row) == 0 {
			continue
		}
		rock_row := []rune{}
		for j, rock := range row {
			rock_row = append(rock_row, rock)
			if rock == 'O' {
				rock_list = append(rock_list, Rock{Vector{i, j}, 'O', 0, false})
			}
			if rock == '#' && all {
				rock_list = append(rock_list, Rock{Vector{i, j}, '#', 0, true})
			}
		}
		picture = append(picture, rock_row)
	}
	return picture, rock_list
}

func (parabolicReflectorDish *ParabolicReflectorDish) Task1(data string) string {
	return strconv.Itoa(parabolicReflectorDish.RockRollNorth(parabolicReflectorDish.CollectRocks(data, false)))
}

func (parabolicReflectorDish *ParabolicReflectorDish) Task2(data string) string {
	pic, r_l := parabolicReflectorDish.CollectRocks(data, true)
	height := len(pic)
	width := len(pic[0])
	old_sum := 0
	cstart := 0
	limit := 0
	pic_temp := ""
	for i := range pic {
		pic_temp = pic_temp + string(pic[i])
	}
	stored_pics := []string{pic_temp}
	for ; limit < 1_000_000_000; limit++ {
		for i, r := range pic {
			for j, c := range r {
				if c == 'O' {
					pic[i][j] = '.'
				}
			}
		}
		for i := range r_l {
			parabolicReflectorDish.RocknRoll(height, width, r_l, i, Vector{-1, 0})
		}
		for i := range r_l {
			if r_l[i].stone == 'O' {
				r_l[i].at_rest = false
			}
		}
		for i := range r_l {
			parabolicReflectorDish.RocknRoll(height, width, r_l, i, Vector{0, -1})
		}
		for i := range r_l {
			if r_l[i].stone == 'O' {
				r_l[i].at_rest = false
			}
		}
		for i := range r_l {
			parabolicReflectorDish.RocknRoll(height, width, r_l, i, Vector{1, 0})
		}
		for i := range r_l {
			if r_l[i].stone == 'O' {
				r_l[i].at_rest = false
			}
		}
		for i := range r_l {
			parabolicReflectorDish.RocknRoll(height, width, r_l, i, Vector{0, 1})
		}
		for i := range r_l {
			if r_l[i].stone == 'O' {
				pic[r_l[i].v.x][r_l[i].v.y] = 'O'
				r_l[i].at_rest = false
			}
		}
		sum := 0
		for _, r := range r_l {
			if r.stone == 'O' {
				sum += r.score
			}
		}
		temp_pic := ""
		for i := range pic {
			temp_pic = temp_pic + string(pic[i])
		}
		found := false
		for i, stored := range stored_pics {
			if stored == temp_pic {
				found = true
				cstart = i
				for i := range r_l {
					parabolicReflectorDish.RocknRoll(height, width, r_l, i, Vector{-1, 0})
				}
				for i := range r_l {
					if r_l[i].stone == 'O' {
						r_l[i].at_rest = false
					}
				}
				for i := range r_l {
					parabolicReflectorDish.RocknRoll(height, width, r_l, i, Vector{0, -1})
				}
				for i := range r_l {
					if r_l[i].stone == 'O' {
						r_l[i].at_rest = false
					}
				}
				for i := range r_l {
					parabolicReflectorDish.RocknRoll(height, width, r_l, i, Vector{1, 0})
				}
				for i := range r_l {
					if r_l[i].stone == 'O' {
						r_l[i].at_rest = false
					}
				}
				for i := range r_l {
					parabolicReflectorDish.RocknRoll(height, width, r_l, i, Vector{0, 1})
				}
				for i := range r_l {
					if r_l[i].stone == 'O' {
						r_l[i].at_rest = false
					}
				}
				limit++
				break
			}
		}
		if !found {
			stored_pics = append(stored_pics, temp_pic)
		} else {
			break
		}
	}
	rem := (1_000_000_000 - limit - 1) % (limit - cstart)
	for j := 0; j < rem; j++ {
		for i := range r_l {
			parabolicReflectorDish.RocknRoll(height, width, r_l, i, Vector{-1, 0})
		}
		for i := range r_l {
			if r_l[i].stone == 'O' {
				r_l[i].at_rest = false
			}
		}
		for i := range r_l {
			parabolicReflectorDish.RocknRoll(height, width, r_l, i, Vector{0, -1})
		}
		for i := range r_l {
			if r_l[i].stone == 'O' {
				r_l[i].at_rest = false
			}
		}
		for i := range r_l {
			parabolicReflectorDish.RocknRoll(height, width, r_l, i, Vector{1, 0})
		}
		for i := range r_l {
			if r_l[i].stone == 'O' {
				r_l[i].at_rest = false
			}
		}
		for i := range r_l {
			parabolicReflectorDish.RocknRoll(height, width, r_l, i, Vector{0, 1})
		}
		for i := range r_l {
			if r_l[i].stone == 'O' {
				r_l[i].at_rest = false
			}
		}
	}
	for _, r := range r_l {
		if r.stone == 'O' {
			old_sum += r.score
		}
	}
	for _, r := range r_l {
		if r.stone != '#' {
			pic[r.v.x][r.v.y] = 'O'
		}
	}
	return strconv.Itoa(old_sum)
}
