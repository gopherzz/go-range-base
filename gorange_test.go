package gorange_test

import (
	"fmt"
	"testing"

	gorng "github.com/gopherzz/go-range-base"
)

// it's stupid test i don't know why it here...
func TestRangeFor(t *testing.T) {
	rng := gorng.Range{}.From(0).To(100)
	rng.For(func(i int32) { fmt.Println(i) })
}

func TestRangeNext(t *testing.T) {
	rng := gorng.Range{}.From(0).To(3)
	var a int32
	for i := rng; !i.IsLast(); i.Next() {
		a++
	}
	if a != 3 {
		t.Errorf("Excepted a == 3, get a = %d", a)
	}
}
