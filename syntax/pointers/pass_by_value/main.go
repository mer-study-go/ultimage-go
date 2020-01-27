// Sample program to show the basic concept of pass by value
package main

import "fmt"

func main() {
	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	fmt.Println("count:\tValue Of [", count, "]\tAddr Of[", &count, "]")

	// Pass the "value of" the count.
	increment(count)
	fmt.Println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

}

// increment declares count as a pointer variable whose value is always
// an address and points to values of type int.
func increment(inc int) {
	// Increment the "value of" inc.
	inc++
	fmt.Println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}
