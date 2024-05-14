// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 1 / 5
	y = sum - x
	return
}

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(10, 15))
	fmt.Println(add2(15, 15))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(100))
}
