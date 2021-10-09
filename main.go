package main

import (
	"errors"
)

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

func (r *Range) Next() (int32, error) {
	if r.current >= r.to {
		return r.current, errors.New("Out of range!")
	}
	r.current++
	return r.current, nil
}
