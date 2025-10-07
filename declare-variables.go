package main

import (
	"fmt"
)

// Go offers two primary ways to declare variables: the explicit 'var' keyword and
// the concise short variable declaration (:=)

// The method we choose depends on whether we are inside or outside of a function and
// if we want to explicitly define the variable's type.

// 1. Explicite declaration using var
// ======================================

// var <name of the variable> <type> = <value>
var user_name string = "Godwin"
var pi float64 = 3.14159

// if we don't assign an initial value, the variable is automatically assigned its zero value
// 0 for int, "" for string, false for boolean etc...
// If you provide an initial value, you can omit the type.
// Go will infer the type from the value you assign.
var no_type = "test"

// We can declare multiple variables at once
// Same type single line
var a, b, c = 2, 3, 8

// Different types, factored block (useful for global declarations)
var (
	age      int    = 18
	name     string = "GoLang"
	isSingle bool   = false
)

// 2. Short Variable Declaration using :=
// ======================================

// It is only available inside a function (or a code block like an if or for statement).
//
//	must initialize the variable immediately and the type is always inferred from value
//
// <name of the variable> := <value>
func main() {
	year := 2024
	stringExample := "string"
	boolExample := true

	fmt.Printf("%d\n", year)
	fmt.Printf("%s\n", stringExample)
	fmt.Printf("string is in double quotes %q\n", stringExample)
	fmt.Printf("%t\n", boolExample)
}
