package missinghistorian

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
  s1 := &Hysteria{solve}
  s2 := &Rudolf{solve}
  s3 := &Mullitover{solve}
  s4 := &Ceressearch{solve}
  s5 := &PrintQueue{solve}
  solvers := []Solver{Solver(s1),
  Solver(s2), Solver(s3),
  Solver(s4), Solver(s5)}
  return solvers[day-1]
}