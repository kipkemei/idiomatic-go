package main

import "fmt"

const x int64 = 10

const (
	idKey = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	// Ways to declare variables.
	// Use var keyword both within and outside function.
	var a int
	var b int = 1
	var c = 2
	var d, e int = 3, 4
	var f, g int
	var h, i = 10, "hello"

	var (
		j int
		k = 20
		l int = 21
		m, n = 22, "hey"
		o, p string
	)

	// Short declaration format operator := is used to declare variables
	// within functions. x := 10 is similar to var x = 10 that uses type inference.
	q := 30
	r, s := 31, "Ni hao"

	// := operator allows you to assign values to existing variables as long as
	// there is one new variable on the left hand side of the :=, then any of
	// the other variables can already exist.
	t := 40
	fmt.Println(t)
	
	t, u := 42, "Yoh!"

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u)

	// Print constants
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(idKey)
	fmt.Println(nameKey)
	fmt.Println(z)
}