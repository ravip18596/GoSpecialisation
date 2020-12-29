Object Oriented Programming in Go
---------------------------------

## Classes and Encapsulation

- Collection of data fields and functions that share a well-defined responsibility
- Example: Point class
  - Used in a geometry program
  - Data: x coordinate, y coordinate
  - Functions:
    - DistToOrigin(), Quadrant()
    - AddXOffset(), AddYOffset()
    - SetX(), SetY()
- Classes are a template which its objects fill them with data
- Contain data fields, not data

### Object

- Instance of a class
- Contains real data
- Example: Point class

### Encapsulation

- Data can be protected from the programmer (i.e. prevent anti-patterns or misuse)
- Data can be accessed only using methods which you have defined
- Maybe we don’t trust the programmer to keep data consistent
- Example: Double distance to origin
    - Option 1: Make method DoubleDist()
    - Option 2: Trust programmer to double X and Y directly

## Support for classes

There is no **class** keyword in golang as in other languages

### Associating Methods with Data

- Method has a receiver type that it is associated with
  - Method can be defined on package level type i.e. why we can not define
    methods on std lib data types like ints or floats
- Use dot notation to call the method

```go
package main

import "fmt"

type MyInt int //using alias

func (mi MyInt) Double() int {
  return int(mi * 2)
}
func main() {
  v := MyInt(2)
  fmt.Println(v.Double())
}
```

### Implicit Method Argument

- Object v (in above example) is an implicit argument to the method
- Double method has MyInt obj as implicit argument passed as call by value

### Structs

- Struct types compose data fields
```go
package main

import (
	"fmt"
	"math"
)
type Point struct{
	x float64
	y float64
}
func (p Point) DistToOrig() {
    t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
    return math.Sqrt(t)
}
func main() {
    p1 := Point(3, 4)
    fmt.Println(p1.DistToOrig())
}
```
- Structs represents traditional feature of classes

### Encapsulation

- Making data fields or methods hidden from the programmer
- Might use a private keyword in another language
- Example: Point struct, Scale() method
- Scale() should multiply x and y coordinates by a constant
- if you don’t trust this to the programmer
  - Might scale one coordinate but not the other
  - Coordinates could become inconsistent
  - Need to hide x and y coordinates
- In case you are building ORM interface and data access methods

#### Hiding in a package

- Go can only hide data/methods in a package
- Variables/functions are only exported if their names start with a capital letter

```go
package data
var X int = 1
var Y int = 2
```
```go
package main

import (
	"data"
	"fmt"
)
func main(){
	fmt.Println(Y)
	fmt.Println(X)
}
```

- Can define *public functions* to allow access to hidden data
- Hide fields of structs by starting field name with a lower-case letter

```go
package data
import "fmt"
type Point struct {
	x,y float64
}
func (p *Point) InitMe(xn,yn float64) {
	p.x = xn
	p.y = yn
}
func (p *Point) Scale(v float64) {
  p.x = p.x * v
  p.y = p.y * v
}
func (p *Point) PrintMe(){
  fmt.Println(p.x, p.y)
}
```
```go
package main

import "data"
func main() {   
	var p data.Point
	p.InitMe(3,4)
	p.Scale(2)
	p.PrintMe()
}
```

### Pointer Receivers

- Regular receiver in method require object copying when they are passed 
  as implicit argument to the method, which can lead to higher memory footprint
- Receiver can be a pointer to a type
- Call by reference, pointer is passed to the method
```go
func (p *Point) OffsetX(v float64) {
	p.x = p.x + v
}
```
- Do not need to reference when calling the method
```go
func main() {
	p := Point{3, 4}
	p.OffsetX(5)
	fmt.Println(p.x)
}
```
- No need to dereference
  - Point is referenced as p, not *p
  - Dereferencing is automatic with . operator

- Good programming practice:
  - All methods for a type have pointer receivers, or
  - All methods for a type have non-pointer receivers
  - Mixing pointer/non-pointer receivers for a type will get confusing
  - Pointer receiver allows modification
