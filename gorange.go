package gorange

type Range struct {
	from    int32
	to      int32
	current int32
}

type RangeFunc func(int32)

func New(from, to int32) Range {
	return Range{from: from, to: to}
}

// Set Range Start value
func (r Range) From(from int32) Range {
	return Range{from: from, current: from}
}

// Set Range End value
func (from Range) To(to int32) Range {
	return Range{
		from: from.from,
		to:   to,
	}
}

// Iterate Range, and run function 'f' with current range value
func (r *Range) For(f RangeFunc) {
	for i := r; !i.IsLast(); i.Next() {
		f(i.Cur())
	}
}

// Return Copy of Range with changed cursor position
func (r *Range) Next() Range {
	if r.IsLast() {
		return *r
	}
	r.current++
	return *r
}

func (r *Range) IsLast() bool {
	if r.current >= r.to {
		return true
	}
	return false
}

func (r *Range) Cur() int32 {
	return r.current
}
