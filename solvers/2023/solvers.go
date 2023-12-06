package solvers

type Solve struct{}
type Solver interface {
	Task1(data string) string
	Task2(data string) string
}

func GetSolver(day int) Solver {
	solve := Solve{}
	s1 := &Trebuchet{solve}
	s2 := &CodeConundrum{solve}
	s3 := &GearRatios{solve}
	s4 := &ScratchCards{solve}
	s5 := &Fertilizer{solve}
	s6 := &WaitForIt{solve}
	solvers := []Solver{Solver(s1), Solver(s2), Solver(s3), Solver(s4), Solver(s5), Solver(s6)}
	return solvers[day-1]
}