package main

import "fmt"

type bill struct {
	flag    bool
	counter int16
	pi      float32
}

type alice struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	// Declare a variable of an anonymous type set to its zero value.
	var e1 struct {
		flag    bool
		counter int16
		pi      float32
	}

	// Display the value
	fmt.Printf("%+v\n", e1)

	// Declare a variable of an anonymous type and init using a struct literal.
	e2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	var b bill
	var a alice
	var b2 bill
	// compiler will complain if we assign a to b as `b = a`
	b = bill(a)
	// implicit conversion works for literal type
	b2 = e2
	fmt.Println(b, a)
	fmt.Println(b2)

	// Display the values.
	fmt.Printf("%+v\n", e2)
	fmt.Println("Flag", e2.flag)
}
