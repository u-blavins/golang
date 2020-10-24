# Golang Training

## Prerequisites

## Learning Resource
[golang-tutorial](https://www.educative.io/blog/golang-tutorial)

[GOLANGBOT](https://golangbot.com/golang-tutorial-part-1-introduction-and-installation/)

## Notes

**Advantages of Go for server side programming**

- Simple Syntax: syntax is simple and concise, not bloated with unnecessary features.
- Easy to write concurrent programs: Concurrency is an inherent part of the language - writing multithreaded progras is a piece of cake.

**Intermediate concepts of Go**

Go uses flexible `for` and `switch` loops. There are also new control structures, such as type switch and `select.

> `if-else`: this construct tests for a conditional statement - logical or boolean.
> `switch-case`: structure is used instead of long `if` statements that compare variables to values.

**Packages**

The approach of writing all source code in a single file is not scalable. It becomes impossible to reuse and maintain code written this way. This is where packages are helpful.

Packages are used to organise Go source code. Packages are a collection of Go source files that reside in the same directory.

Every executable Go application must contain the main function. The `main` function is the entry point for execution. The `main` function should reside in the main package.

> `package packagename` specifies that a particular source file belongs to `packagename`

**Go Modules**

A Go Module is nothing but a collection of Go packages. The import path for the custom package we create is derived from the name of the Go module.

All other third-party packages (such as code from GitHub) which our application uses will be present in the go.mod file, along with the version. The go.mod file is created when we create a new module.

`go mod init learnpackage`

**init function**

The `init` function must not have any return type and it must not have any parameters. It will be called automatically when the package is initalised.

> `func init() {...}`

The `init` function can be used to perform initialisation tasks and also be used to verify the correctness of the program.