package decoratingteam

type Solve struct{}
type Solver interface {
  Task1(data string) string
  Task2(data string) string
}

func GetSolver(day int) Solver {
  solve := Solve{}
  s1 := &SafeCracker{solve}
  solvers := []Solver{Solver(s1)}
  return solvers[day-1]
}