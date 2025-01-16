```Go
package main

import "fmt"

func main() {
    var m map[string]int
    fmt.Println(m[1]) // This will panic because the map is not initialized
}
```