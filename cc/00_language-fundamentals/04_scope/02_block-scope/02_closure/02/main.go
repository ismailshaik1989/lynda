package main

import "fmt"

var x = 0

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
