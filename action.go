package ebuf

func (b *Buffer) Action(fn func()) {
	b.skipping = true
	fn()
	b.skipping = false
	b.States[b.Current].Skip = false
}
