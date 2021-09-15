package template

type Animal interface {
	eat() string
	run() string
	jump() string
}

type Cat struct{}

func (c *Cat) eat() string {
	return "Cat eat"
}

func (c *Cat) run() string {
	return "Cat run"
}

func (c *Cat) jump() string {
	return "Cat jump"
}

type WhiteCat struct {
	Cat
}

func (c *WhiteCat) run() string {
	return "WhiteCat run"
}
