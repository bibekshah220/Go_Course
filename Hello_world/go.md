# Go Programming Notes

A simple Go reference with examples and explanations.

---

## 1. Hello World

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

### Explanation

* `package main` — Defines the main package.
* `import "fmt"` — Imports Go's formatting package.
* `func main()` — Program execution starts here.
* `fmt.Println()` — Prints text to the terminal.

Run:

```bash
go run .
```

---

## 2. Variables

```go
package main

import "fmt"

func main() {
	var name string = "Bibek"
	age := 25

	fmt.Println(name)
	fmt.Println(age)
}
```

### Explanation

```go
var name string = "Bibek"
```

Creates a variable with an explicitly defined type.

```go
age := 25
```

Go automatically determines that `age` is an `int`.

---

## 3. Constants

```go
package main

import "fmt"

const appName = "My App"

func main() {
	fmt.Println(appName)
}
```

### Explanation

A constant cannot be changed after it is declared.

```go
const appName = "My App"
```

---

## 4. Functions

```go
package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	result := add(10, 20)

	fmt.Println(result)
}
```

### Explanation

```go
func add(a int, b int) int
```

This function:

* accepts two `int` values
* returns an `int`
* adds the two values

Output:

```text
30
```

---

## 5. Multiple Return Values

Go functions can return multiple values.

```go
package main

import "fmt"

func divide(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	result, remainder := divide(10, 3)

	fmt.Println("Result:", result)
	fmt.Println("Remainder:", remainder)
}
```

Output:

```text
Result: 3
Remainder: 1
```

---

## 6. If / Else

```go
package main

import "fmt"

func main() {
	age := 20

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}
}
```

### Explanation

Go uses `if` and `else` for conditions.

Unlike some languages, you don't need parentheses around the condition.

---

## 7. Switch

```go
package main

import "fmt"

func main() {
	status := "active"

	switch status {
	case "active":
		fmt.Println("User is active")
	case "inactive":
		fmt.Println("User is inactive")
	default:
		fmt.Println("Unknown status")
	}
}
```

---

## 8. For Loop

Go mainly uses `for` for loops.

```go
package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}
```

Output:

```text
1
2
3
4
5
```

### While-style loop

Go doesn't have a separate `while` keyword.

```go
count := 1

for count <= 5 {
	fmt.Println(count)
	count++
}
```

---

## 9. Arrays

```go
package main

import "fmt"

func main() {
	numbers := [3]int{10, 20, 30}

	fmt.Println(numbers)
	fmt.Println(numbers[0])
}
```

Output:

```text
[10 20 30]
10
```

Arrays have a fixed size.

---

## 10. Slices

Slices are more commonly used than arrays.

```go
package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30}

	numbers = append(numbers, 40)

	fmt.Println(numbers)
}
```

Output:

```text
[10 20 30 40]
```

### Explanation

```go
[]int
```

means a slice of integers.

```go
append(numbers, 40)
```

adds a new value.

---

## 11. Maps

Maps store key-value pairs.

```go
package main

import "fmt"

func main() {
	user := map[string]string{
		"name":  "Bibek",
		"role":  "Developer",
		"city":  "Kathmandu",
	}

	fmt.Println(user["name"])
	fmt.Println(user["role"])
}
```

Output:

```text
Bibek
Developer
```

---

## 12. Structs

Structs are used to create custom data types.

```go
package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	user := User{
		Name:  "Bibek",
		Email: "bibek@example.com",
		Age:   25,
	}

	fmt.Println(user.Name)
	fmt.Println(user.Email)
}
```

### Explanation

```go
type User struct
```

creates a custom `User` type.

The struct contains:

* `Name`
* `Email`
* `Age`

---

## 13. Methods

A method is a function associated with a type.

```go
package main

import "fmt"

type User struct {
	Name string
}

func (u User) SayHello() {
	fmt.Println("Hello", u.Name)
}

func main() {
	user := User{Name: "Bibek"}

	user.SayHello()
}
```

Output:

```text
Hello Bibek
```

---

## 14. Pointers

Pointers store the memory address of a value.

```go
package main

import "fmt"

func changeName(name *string) {
	*name = "John"
}

func main() {
	name := "Bibek"

	changeName(&name)

	fmt.Println(name)
}
```

Output:

```text
John
```

### Explanation

```go
&name
```

gets the address of `name`.

```go
*name
```

accesses the value stored at that address.

---

## 15. Interfaces

Interfaces define behavior.

```go
package main

import "fmt"

type Speaker interface {
	Speak()
}

type Person struct {
	Name string
}

func (p Person) Speak() {
	fmt.Println("Hello, I am", p.Name)
}

func main() {
	var speaker Speaker

	speaker = Person{Name: "Bibek"}

	speaker.Speak()
}
```

### Explanation

Any type that implements the required methods satisfies the interface.

---

## 16. Error Handling

Go commonly handles errors using return values.

```go
package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func main() {
	result, err := divide(10, 0)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(result)
}
```

### Important pattern

```go
result, err := someFunction()

if err != nil {
	// handle error
}
```

This pattern is extremely common in Go.

---

## 17. Packages

You can organize code into packages.

Example:

```text
my-app/
├── go.mod
├── main.go
└── calculator/
    └── calculator.go
```

`calculator/calculator.go`:

```go
package calculator

func Add(a int, b int) int {
	return a + b
}
```

`main.go`:

```go
package main

import (
	"fmt"
	"my-app/calculator"
)

func main() {
	result := calculator.Add(10, 20)

	fmt.Println(result)
}
```

---

## 18. Goroutines

Goroutines allow functions to run concurrently.

```go
package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine")
}

func main() {
	go sayHello()

	time.Sleep(time.Second)
}
```

### Explanation

```go
go sayHello()
```

starts `sayHello()` as a goroutine.

Goroutines are lightweight and are heavily used for concurrent applications.

---

## 19. Channels

Channels allow goroutines to communicate.

```go
package main

import "fmt"

func sendMessage(ch chan string) {
	ch <- "Hello from goroutine"
}

func main() {
	ch := make(chan string)

	go sendMessage(ch)

	message := <-ch

	fmt.Println(message)
}
```

### Explanation

```go
ch := make(chan string)
```

creates a string channel.

```go
ch <- "Hello"
```

sends data to the channel.

```go
message := <-ch
```

receives data from the channel.

---

## 20. HTTP Server

Go has a built-in HTTP package.

```go
package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go API")
}

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", nil)
}
```

Run:

```bash
go run .
```

Open:

```text
http://localhost:8080
```

---

## 21. JSON

Go provides the `encoding/json` package.

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user := User{
		Name: "Bibek",
		Age:  25,
	}

	data, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
```

Output:

```json
{"name":"Bibek","age":25}
```

---

## 22. Go Modules

Create a new project:

```bash
mkdir my-go-app
cd my-go-app
go mod init my-go-app
```

This creates:

```text
go.mod
```

Install a dependency:

```bash
go get github.com/example/package
```

Clean dependencies:

```bash
go mod tidy
```

---

## 23. Testing

Create:

```text
calculator.go
calculator_test.go
```

`calculator.go`:

```go
package calculator

func Add(a int, b int) int {
	return a + b
}
```

`calculator_test.go`:

```go
package calculator

import "testing"

func TestAdd(t *testing.T) {
	result := Add(10, 20)

	if result != 30 {
		t.Errorf("expected 30, got %d", result)
	}
}
```

Run:

```bash
go test ./...
```

---

## 24. Useful Go Commands

```bash
# Check Go version
go version

# Run project
go run .

# Build project
go build

# Run tests
go test ./...

# Format code
go fmt ./...

# Check code
go vet ./...

# Manage dependencies
go mod tidy

# Download dependencies
go mod download

# Show Go environment
go env
```

---

## 25. Recommended Go Project Structure

For a backend/API project:

```text
my-api/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── model/
│   └── middleware/
├── config/
├── migrations/
├── tests/
├── go.mod
├── go.sum
└── README.md
```

### Typical flow

```text
HTTP Request
     ↓
Handler
     ↓
Service
     ↓
Repository
     ↓
Database
```

This keeps HTTP handling, business logic, and database access separated.

---

## 26. Go Development Workflow

```bash
# Create project
mkdir my-api
cd my-api

# Initialize module
go mod init my-api

# Write code

# Format
go fmt ./...

# Test
go test ./...

# Check
go vet ./...

# Run
go run .

# Build
go build
```

---

## Quick Cheat Sheet

```text
Variables       → var / :=
Constants       → const
Function        → func
Struct          → type ... struct
Method          → func (receiver)
Interface       → type ... interface
Slice           → []type
Map             → map[key]value
Pointer         → * / &
Error           → error
Concurrency     → go
Communication   → chan
HTTP            → net/http
JSON            → encoding/json
Testing         → testing
Dependencies    → go.mod
```
