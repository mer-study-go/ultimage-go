// Sample program to show the basic concept of using a pointer to share data.
package user

import "fmt"

// user represents a user in the system.
type user struct {
	name   string
	email  string
	logins int
}

func main() {
	// Declare and initialize a variable named mer of type user.
	mer := user{
		name:  "Mer",
		email: "Mer@merjq.com",
	}

	//** We don't need to include all the fields when specifying field names with a struct literal.

	// Pass the "address of" the mer value.
	display(&mer)

	// Pass the "address of" the logins field from within the mer value.
	increment(&mer.logins)

	// Pass the "address of" the mer value.
	display(&mer)
}

// increment declares logins as a pointer variable whose value is
// always an address and points to value of type int.
func increment(logins *int) {
	*logins++
	fmt.Printf("&logins[%p] logins[%p] *logins[%d]\n\n", &logins, logins, *logins)
}

// display declares u as user pointer variable whose value is always an address
// and points to values of type user.
func display(u *user) {
	fmt.Printf("%p\t%+v\n", u, *u)
	fmt.Printf("Name: %q Email: %q Logins %d\n\n", u.name, u.email, u.logins)
}
