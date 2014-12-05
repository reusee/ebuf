package ebuf

type Watcher interface {
	Operate(Op)
}
