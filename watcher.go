package ebuf

import "container/list"

// Watcher is watching insert/delete operations of a buffer
type Watcher interface {
	// if op.Type is Insert, elem is current state; if op.Type is Delete, elem is the state before deletion
	Operate(op Op, elem *list.Element)
}
