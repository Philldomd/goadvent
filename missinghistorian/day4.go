/*-------------------------------------
Task1: 2496     [OK]    Time: 1.667ms
Task2: 1967     [OK]    Time: 1.0123ms
---------------------------------------*/

package missinghistorian

import (
	"strconv"
	"strings"
)

type Ceressearch struct {
	Solve
}

type Vector struct {
	x, y int
}

func readRuneArrOfArr(data string) [][]rune {
	var ret [][]rune
	for _, r := range strings.Split(data, "\n") {
		if len(r) == 0 {
			continue
		}
		ret = append(ret, []rune(r))
	}
	return ret
}

func checkDir(d Vector, pos Vector, word []rune, word_pos int, puzzle [][]rune) bool {
	len_row := len(puzzle[pos.x])
	len_col := len(puzzle)
	for x := word_pos; x < len(word); x++ {
		x_pos := pos.x + d.x * x
		y_pos := pos.y + d.y * x
		if x_pos < 0 || x_pos >= len_row {
			continue
		}
		if y_pos < 0 || y_pos >= len_col {
			continue
		}
		if puzzle[x_pos][y_pos] != word[x] {
			return false
		}
		if x == len(word)-1 {
			return true
		}
	}
	return false
}

func search(pos Vector, word []rune, word_pos int, dir []Vector, puzzle [][]rune) int {
	ret := 0
	for _, d := range dir {
		if checkDir(d, pos, word, word_pos, puzzle) {
			ret += 1
		}
	}
	return ret
}

func findWordInPuzzel(i int, row []rune, word []rune, dir []Vector, puzzle [][]rune) int {
  ret := 0
	word_pos := 0
	for j, x := range row {
		if x == word[word_pos] {
			word_pos = 1
      ret += search(Vector{i, j}, word, word_pos, dir, puzzle)
		}
		word_pos = 0
	}
	return ret
}

func (Ceressearch *Ceressearch) Task1(data string) string {
	ret := 0
	wordSearch := readRuneArrOfArr(data)
	word := []rune{'X','M','A','S'}
	// TL TM TR
	// ML XX MR
	// BL BM BR
	directions := []Vector{{-1,-1},{-1,0},{-1,1}, //TL TM TR
		{0,-1},{0,1}, //ML MR
		{1,-1},{1,0},{1,1}} //BL BM BT
	for i, r := range wordSearch {
		ret += findWordInPuzzel(i, r, word, directions, wordSearch)
	}
	checkDir(Vector{0,1}, Vector{0, 5}, word, 0, wordSearch)
	return strconv.Itoa(ret)
}

func testCorners(pos Vector, corners [][]rune, dir []Vector, puzzle [][]rune) int {
	ret := 0
	for _, c := range corners {
		if c[0] == puzzle[pos.x + dir[0].x][pos.y + dir[0].y] && 
			c[1] == puzzle[pos.x + dir[1].x][pos.y + dir[1].y] && 
			c[2] == puzzle[pos.x + dir[2].x][pos.y + dir[2].y] && 
			c[3] == puzzle[pos.x + dir[3].x][pos.y + dir[3].y] {
			ret += 1
			break
		}
	}
	return ret
}

func findXMAS(i int, row []rune, corners [][]rune, dir []Vector, puzzle [][]rune) int {
	ret := 0
  for j, x := range row {
		if  j == 0 || j == len(row)-1 {
			continue
		}
		if x == 'A' {
			ret += testCorners(Vector{i, j}, corners, dir, puzzle)
		}
	}
	return ret
}

func (Ceressearch *Ceressearch) Task2(data string) string {
	ret := 0
	wordSearch := readRuneArrOfArr(data)
  corners := [][]rune{
		{'M', 'S', 'M', 'S'},
		{'S', 'M', 'S', 'M'},
		{'S', 'S', 'M', 'M'},
		{'M', 'M', 'S', 'S'},
	}
	directions := []Vector{{-1, -1},{-1, 1},{1, -1},{1, 1}}
	for i, r := range wordSearch {
		if i == 0 || i == len(wordSearch)-1 {
			continue
		}
		ret += findXMAS(i, r, corners, directions, wordSearch)
	}
  return strconv.Itoa(ret)
}