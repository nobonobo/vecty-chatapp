package actions

// Action ...
type Action interface {
	action()
}

type base struct{}

func (b base) action() {}
