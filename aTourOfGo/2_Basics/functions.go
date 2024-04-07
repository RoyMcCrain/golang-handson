package main

import (
	"fmt"
	"math"
	"math/cmplx"
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

func split(num int) (x, y int) {
	x = num * 4 / 9
	y = num - x
	return
}

var c, python, java bool

var (
	Tobe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func convertTypes() {
	x, y := 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	fmt.Println("convertTypes: ", x, y, z)
}

func typeInference() {
	v := 42
	fmt.Printf("v is of type %T\n", v)
	vv := 42.0
	fmt.Printf("vv is of type %T\n", vv)
	vvv := 42i
	fmt.Printf("vvv is of type %T\n", vvv)
	vvvv := "42"
	fmt.Printf("vvvv is of type %T\n", vvvv)
}

func constants() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", math.Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

func numericConstants() {
	const (
		Big   = 1 << 100
		Small = Big >> 99
	)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Println("add:", add(11, 22))
	fmt.Println("add1:", add2(11, 22))

	first, second := swap("hello", "world")
	fmt.Println("swap:", first, second)

	x, y := split(17)
	fmt.Println("split:", x, y)

	var i int
	fmt.Println("variables:", i, c, python, java)

	var i2 int = 42
	d, php, javascript := true, false, "no!"
	fmt.Println("variables2:", i2, d, php, javascript)

	var j int = 1
	k := 2
	e, ruby, typescript := true, false, "yes!"
	fmt.Println("variables3:", j, k, e, ruby, typescript)

	fmt.Printf(
		"basic-types:\nType: %T, Value: %v\nType: $T, Value: %v\nType: %T, Value: %v\n\n",
		Tobe, Tobe, MaxInt, MaxInt, z, z,
	)

	fmt.Println("\n--------------------")
	convertTypes()
	typeInference()
	constants()
	numericConstants()
}
