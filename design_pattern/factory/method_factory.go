package factory

type Student struct {
	Name  string
	Score int
}

func NewStudent(name string) func(score int) *Student {
	return func(score int) *Student {
		return &Student{
			Name:  name,
			Score: score,
		}
	}
}
