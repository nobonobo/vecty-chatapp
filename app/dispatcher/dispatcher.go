package dispatcher

import "github.com/nobonobo/vecty-sample/app/actions"

// ID ...
type ID int

var callbacks = []func(action actions.Action){}

// Dispatch ...
func Dispatch(action actions.Action) {
	for _, c := range callbacks {
		c(action)
	}
}

// Register ...
func Register(callback func(action actions.Action)) ID {
	id := ID(len(callbacks))
	callbacks = append(callbacks, callback)
	return id
}

// Unregister ...
func Unregister(id ID) {
	callbacks = callbacks[:int(id)]
	remain := callbacks[int(id):]
	if len(remain) > 1 {
		callbacks = append(callbacks, remain[id+1:]...)
	}
}
