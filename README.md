just a handy package to help compile and run go code at runtime, taken from https://lemire.me/blog/2021/10/14/calling-a-dynamically-compiled-function-from-go/


Example usage:


```
package main

import (
	"fmt"

	"github.com/jackdoe/go-evalish"
)

func main() {
	dir := "/tmp"

	code := `
package main
    
func Sum(a int,b int) int {
    return a + b
}
`
 
	// compile the .so
	plug := evalish.CompileP(code, dir, "go")

	// get the function
	f := evalish.LookupP(plug, "Sum").(func(int, int) int)

	// use the function
	sum := f(5, 6)
	fmt.Printf("%d", sum)
}

```
