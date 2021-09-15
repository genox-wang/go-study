package strategy

type Behavior interface {
	Do() string
}

type Run struct{}

func (r Run) Do() string {
	return "Run"
}

type Jump struct{}

func (j Jump) Do() string {
	return "Jump"
}

type Strategy struct {
	behavior Behavior
}

func (s *Strategy) setBehavior(behavior Behavior) {
	s.behavior = behavior
}

func (s *Strategy) Do() string {
	return s.behavior.Do()
}
