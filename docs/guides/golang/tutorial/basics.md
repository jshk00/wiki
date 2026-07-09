---
tags:
    - golang
    - basics
    - guides
---

## Basic Info
---

1. Go programs are made up of packages
2. Execution of program starts from `main()` function in `main` package.
3. Example `math/rand` comprises of files that starts with `package rand` 
```go 
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")		
}
```

- use parenthesized import statement for multiple imports
```go 
package main

import (
	"math/rand"
	"fmt"
)
```

- **Exporting →** Everything which starts from capital letter is exportable in go. all remaining are accessible at package level .
```go 
package greeter

// greet is only accessible inside greet package
func greet(name string) string {
	return "Hello " + name
}

// GreetBomb can be accessed outside greet package
func GreetBomb() {
	for {
		fmt.Println("Hello unilimited")
	}
}
```
## How to execute
---

- Execution of program starts from `main()` function in main package
- There should be only one `main.go` file in single project.
- To run use command `go run main.go`
- To build use command `go build .`

> note: in order to use build command you should initialize go-modules for project directory

## Variables
---

- you can declare variables like `var name string` or `var name = "qwerty"` or `name := "qwerty"`.
- you can also declare multiple variables on same line using wallrus(:=) operator `a,b := 1,2`
- There are local and global variables local are bound the scope of block or function. global variables are accessible through out the package.

!!! warning
    global variables cannot be declared by wallrus operator

## Data Types
---

- everything in go is type just like everything in python is object.
- go provide basic types such as integer, float, boolean, string, byte, rune, complex
- `integer` is further devided into int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64.
    - `int` and `uint` are atleast 32 bit in nature but not alias for int32 or uint32 it can also be 64 bit depending upon operatinf system.
- `floats` can be declare as float32(4 bytes), float64(8 bytes) default float type inferred is float64.
- `complex` number are 2 types complex64 and complex128 default is complex128.
- `byte` is alias for uint8.
- `rune` is alias for int32. integer value is meant to represent unicode code point. unicode code reprents character like U+0030 but it doesn't know how to save it. this is where utf-8 comes in picture utf-8 saves character using 1,2,3,4 bytes ascii representation are saved using 1 byte. that's why rune is int32 so maximum storage length of utf-8 is 4 bytes and every string in go is encoded using utf-8. you can read more about unicode in this blog [https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/)
- `string` is slice of bytes and can be declare using double qoutes which respects escape sequence while raw string does not.
- `boolean` → `OR ||` `AND &&` `NEGATION !`

## Control flow
---

- Go provide control flow statement such as if - else, switch, for loop.
- Brackets in Go if else around the condition are omitted.
```go
if condition {
   //Do something
} else if condition {
   //Do something
} else {
   //Do something

```

- short statement can also be used such as assigning and putting condition on it.
```go
if statement; condition {
   //Do something
}

```

- switch statement
```go
switch statement; expression {
case expression1:
	//Dosomething
case expression2:
     //Dosomething
default:
     //Dosomething
}

```

- for loop → init and post part are optional remove them and you will get while loop. remove all and get while loop infinity. use `continue` and `break` keyword to continue and break flow of loop respectively. use range keyword to get iterator dynamically
```go
package main

import "fmt"

// loops in golang
// golang only provide for loop
func loop() {
	count := 5
	// normal loop
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}
	// while loop
	i := 0
	for i < count {
		fmt.Println(i)
		i++
	}
	// infinte while loop
	count = 0
	for {
		if count > 4 {
			break
		}
		fmt.Println(count)
		count++
	}
	// range loop
	names := []string{"alexa", "google homes", "siri", "cortana"}
	for idx, val := range names {
		fmt.Printf("index - %v value - %v\n", idx, val)
	}
}
```

## Functions
---

- basic function in golang
```go
func fn() {
	fmt.Println("simple function")
}
```

- function with paramters
```go
func fn(name string) {
	fmt.Println(name)	
}
```

- function with varidaic parameters
```go
func fn(nums ...int) {
	sum := 0
	for i:=0;i<len(nums);i++{
		sum += nums[i]
	}
	fmt.Println(sum)
}
```

- closure or inline function in go. It can also be assigned to variable
```go
func fn() {
	// closure
	func() {
		fmt.Println("Closure")
	}()
	// closure with params
	func(nums ...int) {
		sum := 0
		for _, i := range nums {
			sum += i
		}
		fmt.Println(sum)
	}(5, 6, 7, 8)
}
```

- function with return type and return name and type declared
```go
// function with multiple return type and arguments
func getNames(name1, name2 string) (string, string) {
	return strings.ToTitle(name1), strings.ToUpper(name2)
}

// function with return variable mentioned
func returnMentioned() (count int) {
	count = 10
	return
}
```

- `init` function
- executes before main function and use cases include initilizing some value or configuration at runtime
- each  package have there own init function
```go
package main

func main(){
	fmt.Println("main function")
}

func init(){
	fmt.Println("init function")
}

//output
// Init function
// main function
```

## Errors
---

- Go’s way of dealing with an error is to explicitly return the error as a separate value
- you can check if returned error value from function is nil or not for error checking
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("non-existing.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file.Name() + "opened succesfully")
	}
}
```

- custom error can be defiened using **`errors.New("error something occurred")`**
- custom error with formatted message `fmt.Errorf("Err is: %s", "database connection issue")`
- advanced custom error messages
    - in below program `inputError` is a type of struct which implements the `Error()` interface method.
    - and `missingField` is addtional field in `inputError` to provide more information about error along with it’s method called `getMissingField`.
    - lastly `validate()` function will return `pointer to inputError that can be nil` and `advancedCustomError()` will type assert error and get the missing field name.
```go
// advanced custom error message
func advancedCustomError() {
	err := validate("", "")
	if err != nil {
		if err, ok := err.(*inputError); ok {
			fmt.Println(err)
			fmt.Printf("Missing Field is %s\n", err.getMissingField())
		}
	}
}

// function that return struct which implements Error() interface method
func validate(name, gender string) error {
	if name == "" {
		return &inputError{message: "Name is mandatory", missingField: "name"}
	}
	if gender == "" {
		return &inputError{message: "Gender is mandatory", missingField: "gender"}
	}
	return nil
}

// struct type for holding value
type inputError struct {
	message      string
	missingField string
}

// impleneting error interface
func (i *inputError) Error() string {
	return i.message
}

// addtional method for performing necessary action
func (i *inputError) getMissingField() string {
	return i.missingField
}
```

## Panic and recover
---

- you can wrap error using  `fmt.Errorf("E2: %w", e1)` %w is used to wrap an error.

## Defer keyword
---

- defer keyword send any execution to end of function
- deferred function always execute even if surrounding function execution failed abrruptly.
- deferred function will execute before surrounding function returns.
- mutliple defer keyword execute last in first out order
```go
package main
import "fmt"
func main() {
    i := 0
    i = 1
    defer fmt.Println(i)
    i = 2
    defer fmt.Println(i)
    i = 3
    defer fmt.Println(i)
}
// output
// 3
// 2
// 1
```

## Pointers
---

- pointer declaration `var sptr *string`
- pointer can be initialized uisng two way one is `sptr = new("alex")` and second is `sptr = &name`
- for priting the value of pointer you can use astricks operator before variable name `*sptr` effectively dereferencing it.
- below is the image of pointer deferencing →

![pointerInGo](../../../assets/images/golang/basics/pointersInGo.webp)

```go
package main

import "fmt"

func pointers() {
	// declaration of pointer
	var sptr *string
	name := "Jayesh"
	sptr = &name
	var ssptr **string = &sptr
	fmt.Println("original variable:", name)
	fmt.Println("pointer adress:", sptr)
	fmt.Println("referenced variable value:", *sptr)
	fmt.Println("double refernced ssptr:", ssptr)
	fmt.Println("ssptr value to original pointer:", *ssptr)
	fmt.Println("actual ssptr point value derefrenced:", **ssptr)
}

// output
//WARN: this addresses may changes on your system
original variable: Jayesh
pointer adress: 0xc00004c230
referenced variable value: Jayesh
double refernced ssptr: 0xc00000e028
ssptr value to original pointer: 0xc00004c230
actual ssptr point value derefrenced: Jayesh
```

## Structs
---

- struct is collection of fields.
- member of struct can be accessed by dot notation
- struct are pass by value by default.
    - if you pass struct to function or new variable a copy of struct is passed
    - to manage struct state they need to be pass by reference (pointers)
- structs can be nested.
- nested field are also accessed by dot notation
```go
// simple struct
type employee struct {
	name   string
	age    int
	salary int
}
// nested struct
type employee struct {
	name   string
	age    int
	salary int
	company employer
}
type employer struct {
	name   string
	location string
}
```

## Slices
---

- slices in golang represents three follwing things →
    - pointer to underlying array →`always pass by reference`
    - current lenght of underlying array →`len()`
    - total capacity which is the maximum capacity to which the underlying array can expand → `cap()`
    ```go
    type SliceHeader struct {
            Pointer uintptr
            Len  int
            Cap  int
    }
    ```
    ![slices](../../../assets/images/golang/basics/slice.webp)
    
- default value of slice is usable.

## Maps
---

- default nil value of map cannot be used.
- allowed key types are →
    - boolean
    - numeric
    - string
    - pointer
    - channel
    - interface types
    - structs – if all it’s field type is comparable
    - array – if the type of value of array element is comparable
- not alllowed key types →
    - Slice
    - Map
    - Function
- if there are duplicate key the recent value for key is considerd and previous value is discarded
```go
salary := map[string]int{"manoj": 2000}
	salary["manoj"] = 3000
	salary["nitin"] = 1000
	fmt.Println("salary map:", salary)

// output
salary map: map[manoj:3000 nitin:1000]
```

- you can iterate over map using for range loop
```go
for k, v := range salary{
	fmt.Printf("key: %v, value: %v\n", k,v)
}
```

- map are refrenced data type and is not safe for concurrent use. two variable assigned to each other containing map will point to same map and changes will be reflected wise versa.
```go
func main(){
	user := map[string]string{
		"name":       "John wick",
		"profession": "assasination",
	}
	fmt.Println("original -> ", user)
	mapPassed(user)
	fmt.Println("modified -> ", user)
}

// changes to map will reflect to in parent function to
func mapPassed(user map[string]string) {
	user["status"] = "rich"
}
// output
// map[name:John wick profession:assasination]
// map[name:John wick profession:assasination status:rich]
```

## Methods
---

- methods are receiver function on specific type, eg struct or other function or interface.
- methods can access properties of receiver and other methods on that type.
- basics syntax `func (receiver receiver_type) some_func_name(arguments) return_values`.
```go
type user struct {
	id          int
	name        string
	email       string
	password    string
	description string
}

// method syntax func (receiver receiver_type) some_func_name(arguments) return_values
func (u *user) validate() error {
	switch {
	case !strings.Contains(u.email, "@"):
		return errors.New("enter proper mail format")
	case len(u.password) < 8:
		return errors.New("password should be greater than 8 characters")
	case u.description == "":
		return errors.New("descrition should not be empty")
	}
	return nil
}
```

- methods on function type
```go
type Greeting func(name string) string

func (g Greeting) exclamation(name string) string {
	return g(name) + "!"
}

func (g Greeting) upper(name string) {
	fmt.Println(strings.ToUpper(name))
}

func main() {
	english := Greeting(func(name string) string {
		return "Hello, " + name
	})
	fmt.Println(english("ANisus"))
	fmt.Println(english.exclamation("ANisus"))
	english.upper("ANisus")
}
```

## Interfaces
---

- basic syntax
```go
type interface_name interface{
	method_name(argument_name argument_type) returnvalues and types 
}

// example
type animal interface {
	breathe()
	walk()
}
```

- default value of interface is nil.
- Interface are implemented implicitly
- Helps write more modular and decoupled code between different parts of codebase – It can help reduce dependency between different parts of codebase and provide loose coupling this is application of interface.
- A type implements an interface if it defines all methods of an interface. 
- If that defines all methods of another interface then it implements that interface. In essence, a type can implement multiple interfaces.
![interfaces](../../../assets/images/golang/basics/interface_1.webp)

- type asserting syntax `val := i.({type})`
- to get concret value of interface type assertion is used in golang.
- An empty interface has no methods , hence by default all concrete types implement the empty interface. If you write a function that accepts an empty interface then you can pass any type to that function
