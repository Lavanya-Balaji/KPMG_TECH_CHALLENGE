//We have a nested object. We would like a function where you pass in the object and a key and get back the value.
//The choice of language and implementation is up to you.
//Example Inputs
//object = {“a”:{“b”:{“c”:”d”}}}
//key = a/b/c
//object = {“x”:{“y”:{“z”:”a”}}}
//key = x/y/z
//value = a


// Step 1 : Open https://go.dev/play/
// Step 2 : Copy the Below Code and Run in the Above Go Playground

package main

import "fmt"

type x1 struct {
	x string
}

type y1 struct {
	y  string
	a  string
	x1 x1
}
type z1 struct {
	z  string
	y1 y1
}
type a1 struct {
	a  string
	z1 z1
}

func main() {

	a1 := a1{
		a: "x",

		z1: z1{
			z: "y",
			y1: y1{
				y: "z",
				a: "a",
			},
		},
	}
	fmt.Println("For Input Object:", a1)
	fmt.Println("key:", a1.a, "/", a1.z1.z, "/", a1.z1.y1.y)
	fmt.Println("Value:", a1.z1.y1.a)
}
