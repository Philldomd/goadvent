package solvers

import (
	"strconv"
	"strings"
)

type MirageMaintenance struct {
	Solve
}

func (mirageMaintenance MirageMaintenance) reverseInts(in []int) []int {
	for i, j := 0, len(in)-1; i<j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
	return in
}

func (mirageMaintenance MirageMaintenance) convert(s []string) []int {
	i := []int{}
	for _, x := range s {
		y, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		i = append(i, y)
	}
	return i
}

func (mirageMaintenance MirageMaintenance) sequence(i []int) int {
	s := 0
	for _, x := range i {
		s += x
	}
	if s == 0 {
		return 0
	}
	diffs := []int{}
	for n := 0; n < len(i)-1; n += 1 {
		diffs = append(diffs, i[n+1]-i[n])
	}
	return i[len(i)-1] + mirageMaintenance.sequence(diffs)
}

func (mirageMaintenance *MirageMaintenance) Task1(data string) string {
	sum := 0
	for _, D := range strings.Split(data, "\n") {
		if len(D) == 0 {
			continue
		}
		sum += mirageMaintenance.sequence(mirageMaintenance.convert(strings.Fields(D)))
	}
	return strconv.Itoa(sum)
}

func (mirageMaintenance *MirageMaintenance) Task2(data string) string {
	sum := 0
	for _, D := range strings.Split(data, "\n") {
		if len(D) == 0 {
			continue
		}
		sum += mirageMaintenance.sequence(mirageMaintenance.reverseInts(mirageMaintenance.convert(strings.Fields(D))))
	}
	return strconv.Itoa(sum)
}
