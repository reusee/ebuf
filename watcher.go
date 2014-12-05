package ebuf

// Watcher is watching insert/delete operations of a buffer
type Watcher interface {
	Operate(Op)
}
