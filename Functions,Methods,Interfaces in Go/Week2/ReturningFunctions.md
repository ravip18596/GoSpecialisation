Returning Functions
-------------------

## Functions as return value

- Functions can return functions 
- Might create a function with controllable parameters
- Example: Distance to Origin function
  - Takes a point (x, y, coordinates)
  - Returns distance to origin
- What if I want to change the origin?
  - Option 1: Pass origin as argument
  - Option 2: Define function with new origin
    
## Function Defines a Function

- Origin(o_x,o_y) location is passed as an argument 
- Origin(o_x,o_y) is built into the returned function fn
- Dist1() and Dist2() have different origins

```go
package main

import (
	"fmt"
	"math"
)

func MakeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {
	fn := func(x, y float64) float64 {
		// dist = sqrt((x-o_x)^2 - (y-o_y)^2)
		dist := math.Sqrt(math.Pow((x-o_x), 2) + math.Pow((y-o_y), 2))
		return dist
	}
	return fn
}

func main() {
	Dist1 := MakeDistOrigin(0, 0)
	Dist2 := MakeDistOrigin(2, 1)

	fmt.Println(Dist1(2,2))
	fmt.Println(Dist2(2,2))
}
```

## Function Environment

- Environment of function include all the variable names defined locally in the function
- **Closure**
  - Function + its environment
  - When functions are passed/returned, their environment comes with them!
  - In above function `MakeDistOrigin`, o_x and o_y variable are in the closure of fn()