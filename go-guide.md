# Go Guide
In here I will *try* to save some of the most important lessons I learn while learning Golang.

### Table of Contents
1. [Install](#Install)
2. [Packages](#packages)
3. [Imports](#imports)
4. [Functions](#functions)
5. [Variables](#packages)
6. [For](#for)
7. [If](#if)
8. [Switch](#switch)
9. [Defer](#defer)

#### Install

Installing Go is rather easy since you have multiple options, allowing you to install it your way. I prefer to install it through my package manager, however it can be installed through [Go Intallation Guide](https://golang.org/doc/install).

One thing to note is that the `GOPATH` environment variable is very important since it specifies the location of your workspace. If no `GOPATH` is set, it is assumed to be `$HOME/go` on Unix systems and `%USERPROFILE%\go` on Windows. If you want to use a custom location as your workspace, you can set the GOPATH environment variable by running the following command.

```bash
# example of defining /home as the GOPATH
export GOPATH=$HOME
```

If you're using [VSCode's Go extension](https://github.com/Microsoft/vscode-go), it will also be useful to set the new `GOPATH` in the workspace settings, so that VSCode knows the new path to display import errors and such.

```bash
"go.gopath" = "$GOPATH"
```

#### Packages

Every Go program is made up of packages. When importing a package, only exported names can be referred to. Any "unexported" names are not accessible from outside the package.

A name is exported if it begins with a capital letter. For example, `Pi` is an exported name, which is exported from the `math` package.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Give me a slice of", math.Pi)
}
```

[⬆ Back to the top!](#go-guide)

#### Imports

Imports can be either multi-line or factored, but it is a good practice to use the latter. 
```go
// multi-line import
import "fmt"
import "math"
// factored import
import (
	"fmt"
	"math"
)
```

[⬆ Back to the top!](#go-guide)

#### Functions

Functions have their types defined after the declaration of parameters. Functions can also have multiple returns.

```go
func sub(x int, y int) int {
    return x - y
}

func add(x int, y int) (string, int) {
    return "Sum is equal to", y + x
}
```

Returns can also be naked, if return values are named at the top of the function, however, they harm readability in long functions.
```go
func moduleOf2(num int) (x int) {
	x = num % 2
	return
}
```

[⬆ Back to the top!](#go-guide)

#### Variables

A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted, taking the type of the initializer. The same can be done for function arguments.

```go
var a, b = 10, "ten"
```

When listing several variables of the same type, the type only needs to be defined last.

```go
var this, is, pretty, cool bool
```

Inside a function, the `:=` **short assignment statement** can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

```go
var a int = 2

func main() {
	b := 1
	this, is := "even nicer!", true

	fmt.Println("My name is Afonso, I'm", a, b, "and this is", this, is)
}

```

Furthermore, variables can be "factored" into blocks, as with import statements.

```go
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```

Go's **basic variable types** are

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

Variables declared without an explicit initial value are given their *zero value*.

- 0 for numeric types

- false for the boolean type, and

- "" (the empty string) for strings.

You can have **type conversion** by simply using expression `T(v)` converts the value `v` to the type `T`.

```go
i := 42
f := float64(i)
u := uint(f)
```

When there are numeric values involved, type inference is, again, dependant on the right hand side. 

```go
var i int         // int 
j := i            // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
fmt.Printf("These are our types\nj -> %T\nf -> %T\ng -> %T", j, f, g)

/* OUTPUT
These are our types
i -> int
f -> float64
g -> complex128 
*/
```

**Constants** cannot be declared using the := syntax and can be character, string, boolean, or numeric values.

`const Pi = 3.14`

**Constants** can be **untyped** taking the type by its context, and can be high precision values.

```go
const (
	// Shift 1 bit left 100 places.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1 aka 2.
	Small = Big >> 99
)
```

[⬆ Back to the top!](#go-guide)

#### For

Like many other languages, the basic for loop has three components separated by semicolons, the init statement, the condition expression and the post statement; and it will stop iterating once the boolean condition evaluates to `false`. As you can see the expression does not need to be surrounded by parentheses `( )` but the braces `{ }` are required. 
```go
sum := 0
for i := 0; i < 10; i++ {
	sum += i
}
```

The init and post statements are optional.

```go
sum := 1
for ; sum < 1000; {
	sum += sum
}
```

And so are the semicolons, the **For is Go's "while"**

```go
sum := 1
for sum < 1000 {
	sum += sum
}
```

The so called "while (true)" can be just written like this: **`for {		}`**

[⬆ Back to the top!](#go-guide)

#### If

Go's `if`, similar to `for` loops, do not need be surrounded by parentheses ( ) but the braces { } are required.

```go
if grade > 10 {
	return "good enough"
}
```

Like `for`, the `if` statement can start with a **short statement** to execute before the condition. Obviously, variables declared by the statement are only in scope until the end of the if.

```go
if grade := getGrade(); grade > 10 {
	return grade
} else {
	grade += 10
}
```

As you can see, the short statement is also available inside any of the `else` blocks, making it extremely useful to reduce the number of calls of a certain function. This is also really good in terms of performance, since the variable is only available inside these blocks and does not occupy a place in the stack.

[⬆ Back to the top!](#go-guide)

#### Switch

Go's switch is like the one found in C, C++, Java, JavaScript, etc. however Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases do not need to be constants. And, of course, cases are evaluated from top to bottom, stopping when a case succeeds.

```go
switch state := getGameState(); state {
	case "menu":
		fmt.Println("Do you want to play?")
	case "boss":
		fmt.Println("Are you dead yet?")
	case "gameover":
		fmt.Println("Yep, guess so...")
	default:
		fmt.Println("Playing")
}
```

A `switch true { }` can also be used, either with true or not (`switch { }`), effectively creating a clean if-then-else chain.

[⬆ Back to the top!](#go-guide)

#### Defer

A defer statement defers the execution of a function until the surrounding function returns, however, note that the deferred call's arguments are evaluated immediately.

```go
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
// output -> hello world
```

When stacking defers calls are pushed onto a stack, and, when the function returns, its deferred calls are executed in last-in-first-out order.

```go
for i := 0; i < 10; i++ {
	defer fmt.Print(i)
}
// output -> 9 8 7 6 5 4 3 2 1 0
```

[⬆ Back to the top!](#go-guide)