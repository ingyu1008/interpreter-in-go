package reader

type Reader struct {
	input string
	pos   int
}

func New(input string) *Reader {
	r := &Reader{input: input}
	r.pos = 0
	return r
}

func (r *Reader) ReadPos(pos int) int {
	if pos < 0 || pos >= len(r.input) {
		return -1
	}

	return int(r.input[pos])
}
