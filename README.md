# go-range-base
## Simple Range for Golang
*For(func(int32))*
```go
package main

import (
  "fmt"
  
  gorng "github.com/gopherzz/go-range-base"
)

func main() {
  rng := gorng.Range{}.From(1).To(100)
  rng.For(func(i int32) { fmt.Println(i) })
}
```
*Next() int32*
```go
package main

import (
  "fmt"
  
  gorng "github.com/gopherzz/go-range-base"
)

func main() {
  rng := gorng.Range{}.From(0).To(3)
	var a int32
	for i := rng; !i.IsLast(); i.Next() {
		a++
	}
	if a != 3 {
		t.Errorf("Excepted a == 3, get a = %d", a)
	}
}
```
