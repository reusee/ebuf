package ebuf

type Watcher interface {
	Insert(op Op)
	Delete(op Op)
}
