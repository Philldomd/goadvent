package solvers

type Solve struct{}
type Solver interface {
	Task1(data string) string
	Task2(data string) string
}

func Sum(arr []int) int {
  sum := 0
	for _, i := range arr {
    sum += i
	}
	return sum
}

func GetSolver(day int) Solver {
	solve := Solve{}
	s1 := &Trebuchet{solve}
	s2 := &CodeConundrum{solve}
	s3 := &GearRatios{solve}
	s4 := &ScratchCards{solve}
	s5 := &Fertilizer{solve}
	s6 := &WaitForIt{solve}
	s7 := &CamelCards{solve}
	s8 := &HauntedWasteland{solve}
	s9 := &MirageMaintenance{solve}
	s10 := &PipeMaze{solve}
	s11 := &CosmicExpansion{solve}
	s12 := &HotSprings{solve}
	solvers := []Solver{Solver(s1), 
		Solver(s2), Solver(s3), 
		Solver(s4), Solver(s5), 
		Solver(s6), Solver(s7), 
		Solver(s8), Solver(s9), 
		Solver(s10), Solver(s11),
	  Solver(s12)}
	return solvers[day-1]
}