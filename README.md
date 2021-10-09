# go-range-base
*Simple Range for Golang*

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
