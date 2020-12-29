Variadic Functions
-------------------------------

- Functions can take a variable number of arguments
- Use ellipsis **…**  to specify
- Treated as a slice inside function
- Can pass a slice to a variadic function
- Need the **…** suffix

```go
package main

import "fmt"
func getMax(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func main() {
    fmt.Println(getMax(1, 3, 6, 4))
    vslice := []int{1, 3, 6, 4}
    fmt.Println(getMax(vslice...))
}
```

Deferred Functions
------------------

- Call can be deferred until the surrounding function completes
- Typically used for cleanup activities
- Arguments of a deferred call are evaluated immediately. So, i will be printed as 2.
```go
package main

import "fmt"
func main(){
	i := 1
	defer fmt.Println(i+1)
	i++
	fmt.Println("Hello!")
}
```