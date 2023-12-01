package solvers

type Solve struct{}
type Solver interface {
	Task1(data string) int
  Task2(data string) int
}

func GetSolver(day int) Solver {
  solve := Solve{}
  s1 := &Trebuchet{solve}

  solvers := []Solver{Solver(s1)}
  return solvers[day-1]
}