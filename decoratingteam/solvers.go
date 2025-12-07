package decoratingteam

type Solve struct{}
type Solver interface {
  Task1(data string) string
  Task2(data string) string
}

func GetSolver(day int) Solver {
  solve := Solve{}
  s1 := &SafeCracker{solve}
  s2 := &GiftShop{solve}
  s3 := &Lobby{solve}
  s4 := &PrintingDepartment{solve}
  solvers := []Solver{Solver(s1),
  Solver(s2),Solver(s3),
  Solver(s4)}
  return solvers[day-1]
}