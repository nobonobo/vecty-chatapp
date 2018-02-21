package router

// Setuper ...
type Setuper interface {
	Setup()
}

// Teardowner ...
type Teardowner interface {
	Teardown()
}
