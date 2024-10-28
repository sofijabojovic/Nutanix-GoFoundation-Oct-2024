# Go Foundation

## Magesh Kuppan
- tkmagesh77@gmail.com

## Prerequisites
- Go Tool (https://go.dev/dl)
- Visual Studio Code
- Go extension for VS Code (https://marketplace.visualstudio.com/items?itemName=golang.Go)

## Schedule
- Session-01    : 01:30 hrs
- Tea Break     : 00:15
- Session-02    : 01:30 hrs
- Lunch Break   : 00:45
- Session-03    : 01:30 hrs
- Tea Break     : 00:15
- Session-04    : 01:30 hrs

## Repository
- https://github.com/tkmagesh/nutanix-gofoundation-oct-2024

## Methodology
- No powerpoint
- Discuss & Code

## Why Go?
- Simple Language
    - ONLY 25 keywords
    - No access modifiers (NO public/private/protected)
    - No classes (only structs)
    - No inheritance (only composition)
    - No reference types (everything is a value)
    - No pointer arithmatic
    - No exceptions (only errors)
    - No try-catch-finally construct
    - No implicit type conversion
- Compiled to machine code
    - Has tooling support for cross compilation
- Performance
    - Close to hardware
    - Compile to machine code
    - Equivalent to C++
- Concurrency Support
    - Managed Concurrency (vs OS Thread based concurrency)
    - Goroutines - concurrency model
        - Cheap (4 KB)
        - Efficient (minimizes context switch)
    - Concurrency support is built in the language itself
        - "go" keyword, channel ("chan") data type, channel ("<-") operator, for-range, select-case
        - Standard library
            - "sync" package
            - "sync/atomic" package

## Compilation
```shell
go build <filename.go>
```

### Compilation & Execute
```shell
go run <filename.go>
```

### Cross Compilation
#### List the Go tooling system environment variables
```
go env
```
#### Env variables for cross compilation
```
go env GOOS GOARCH
```
#### List of supported platforms
```
go tool dist list
```
#### How to cross compile
```
GOOS=windows GOARCH=amd64 go build 01-hello-world.go
```

## Data Types
- bool
- string
- integers
    - int8
    - int16
    - int32
    - int64
    - int
- unsigned integers
    - uint8
    - uint16
    - uint32
    - uint64
    - uint
- floating points
    - float32
    - float64
- complex numbers
    - complex64 (real[float32] + imaginary[float32])
    - complex128 (real[float64] + imaginary[float64])
- type alias
    - byte
    - rune (unicode code point)

### Zero values

| type | value |
| -------|-------- |
| int, uint, float | 0 |
| string | "" |
| bool | false |
| func | nil |
| struct | struct instance |
| pointer |nil |
| interface | nil | 

### Variables in function scope & package scope
#### function scope
    - Can use :=
    - Cannot have unused variables
#### package scope
    - Cannot use :=
    - Can have unsed variables

## Programming constructs
### if else
### switch-case
### for

## functions
- anonymous functions
- higher order functions
    - assign functions as values to variables
    - pass functions as arguments to other functions
    - return functions as return values
- variadic functions
- deferred functions

## Collections
### Array
    - fixed sized typed collection
### Slice
    - varying sized typed collection
### Map

