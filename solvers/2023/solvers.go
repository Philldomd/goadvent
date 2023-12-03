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

	solvers := []Solver{Solver(s1), Solver(s2), Solver(s3)}
	return solvers[day-1]
}