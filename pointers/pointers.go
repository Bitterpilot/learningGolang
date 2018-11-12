package main

import "fmt"

func main() {
	x := 15
	a := &x         // point to x  (memory address)
	fmt.Println(a)  // prints out the mem addr.
	fmt.Println(*a) // read a through the pointer, so this will print out a value (15 in this case)
	
	*a = 5          // sets the value pointed at to 5, which means x is modified (since x is stored at the mem addr)
	fmt.Println(x)  // see the new value of x
	*a = *a**a      // what is the value of x now? 
	fmt.Println(x)  // prints a value
	fmt.Println(*a) // prints a value

	fmt.Println(&x) // prints a memory address
	fmt.Println(a)  // prints a memory address
}