```Go
package main

import "fmt"

func main() {
    var m map[string]int
    m = make(map[string]int)
    m["one"] = 1
    fmt.Println(m["one"]) // Output: 1
}
```