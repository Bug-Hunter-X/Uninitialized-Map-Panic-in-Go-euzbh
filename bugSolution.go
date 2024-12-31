```go
package main

import "fmt"

func main() {
    var m = make(map[string]int)
    m["a"] = 1
    fmt.Println(m["a"])
}
```