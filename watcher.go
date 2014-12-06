package ebuf

import "github.com/reusee/rope"

// Watcher is watching insert/delete operations of a buffer
type Watcher interface {
	Operate(Op, *rope.Rope)
}
