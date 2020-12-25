Function Types
--------------

## First class values

- Functions can be treated like other types
    - Variables can be declared with a function type
    - Can be **created dynamically**
    - Can be **passed as arguments** and **returned as values**
    - Can be **stored in data structures**
    
## Variables as functions

- Function can be declared as variable
- Example
```go
package main

import "fmt"

var funcVar func(int) int
func incFn(x int) int { 
   return x + 1
} 
func main() { 
   funcVar = incFn
   fmt.Print(funcVar(1))
} 
  ```
  
## Function as arguments

- Function can be passed as argument to another function
```go
package main
import "fmt"
func applyIt(afunct func (int) int,val int) int {
	return afunct(val)
}

func incFn(x int) int { return x+1 }
func decFn(x int) int { return x-1 }

func main(){
	fmt.Println(applyIt(incFn,2))
	fmt.Println(applyIt(decFn,2))
}
```

## Anonymous Functions

It is not necessary to name a function
```go
package main

import "fmt"

func applyIt (afunct func (int) int, val int) int{
	return afunct(val)
} 
func main(){
	result := applyIt(func (x int) int { 
		return x+1
	} , 2)
	fmt.Println(result)
}
```