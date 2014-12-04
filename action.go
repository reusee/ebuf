package ebuf

// Action wraps a function in which all operations being treated as single operation
func (b *Buffer) Action(fn func()) {
	b.skipping = true
	fn()
	b.skipping = false
	b.States[b.Current].Skip = false
}
