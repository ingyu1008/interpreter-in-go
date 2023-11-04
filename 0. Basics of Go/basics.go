package main

func main() {
	// How to declare variables in go
	// var <varName> <type> = <value>
	var a int
	var f float32 = 3.14
	var i, j, k int = 0, 1, 2
	var d = 5 // type can be ommitted if it has initial value
	a = 2
	println(a, f, i, j, k, d)

	// Constants
	// use keyword [const] instead of [var]
	const helloWorld = "Hello World!"
	const (
		x = iota
		y
		z
	)
	// Groupping is possible with parenthesis
	// iota will give values 0, 1, 2, ... in order

	// Strings are immutables in Go
	myString := "Immutable"
	println(myString)

	// Type Conversion
	// No implicit type conversion is done
	// every conversion has to be explicit as [Type(var)]
	var n int = 100
	var u uint = uint(n)
	var fl float32 = float32(n)
	println(fl, u)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)

	// Operators
	// Same with C++
	var p *int
	var X int = 10

	p = &X
	println(*p)

	// Conditionals
	// Ifs are simillar to C++
	a = 1
	if a == 1 {
		println("a is 1")
	} else {
		println("a is 0")
	}

	if temp := a + a; temp <= 1 {
		println("2a <= 1")
	} else {
		println("2a > 1")
	}

	// Switch is different
	// break is not needed, use fallthrough to make it work as c++
	switch a {
	case 1:
		println("[switch] a is 1")
		fallthrough
	default:
		println("fallthrough")
	}

	// switch can also be done with types
	switch t := interface{}(p); t.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case *int:
		println("int pointer")
	default:
		println("undefined")
	}

	// Loops
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	println(sum)

	arr := []int{2, 3, 5, 8}
	for idx, val := range arr {
		println(idx, val)
	}
}
