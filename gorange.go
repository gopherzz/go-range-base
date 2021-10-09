package gorange

type Range struct {
	from    int32
	to      int32
	current int32
}

type RangeFunc func(int32)

func (r Range) From(from int32) Range {
	return Range{from: from, current: from}
}

func (from Range) To(to int32) Range {
	return Range{
		from: from.from,
		to:   to,
	}
}

func (r *Range) For(f RangeFunc) {
	for i := r.from; i < r.to; i++ {
		f(i)
	}
}

func (r *Range) Front() (Range, error) {
	return Range{from: r.from, to: r.to, current: r.from}, nil
}

func (r *Range) Cur() int32 {
	return r.current
}

func (r *Range) IsLast() bool {
	if r.current >= r.to {
		return true
	}
	return false
}

func (r *Range) Next() Range {
	if r.IsLast() {
		return Range{from: r.from, to: r.to, current: r.from}
	}
	r.current++
	return Range{from: r.from, to: r.to, current: r.current}
}
