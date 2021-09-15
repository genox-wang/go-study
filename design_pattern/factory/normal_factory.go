package factory

type Persion struct {
	Name string
}

func (p Persion) GetName() string {
	return p.Name
}

func NewPersion(name string) *Persion {
	return &Persion{
		Name: name,
	}
}
