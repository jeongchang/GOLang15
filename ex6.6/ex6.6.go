package main

import "fmt"

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c)

	k := 1
	t := 2

	k, t = t, k

	fmt.Println("k : ", k, ", t : ", t)

}
