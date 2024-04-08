package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func forFunc() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("forFunc: ", sum)
}

func forContinued() {
	sum := 1
	for sum < 10000 { // go-fmtでフォーマットされる
		sum += sum
	}

	fmt.Println("forContinued: ", sum)
}

// 平方根の処理はスキップ
func ifFunc(x float64) {
	if x < 0 {
		ifFunc(-x)
	}
	fmt.Println("ifFunc: ", fmt.Sprint(math.Sqrt(x)))
}

func PowPrint(x, n, lim float64) {
	var result float64
	if v := math.Pow(x, n); v < lim {
		result = v
	} else {
		result = lim
	}
	fmt.Println("PowPrint: ", result)
}

func switchFunc() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func switchFunc2() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println("today: ", today)
	fmt.Printf("today type: %T\n", today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func switchFunc3() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferFunc() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println("1 ----", i)
	}
	for i := 0; i < 10; i++ {
		defer fmt.Println("2 ----", i)
	}
	fmt.Println("done")
}

func main() {
	forFunc()
	forContinued()
	ifFunc(2)
	PowPrint(3, 2, 10)
	PowPrint(3, 3, 20)
	switchFunc()
	switchFunc2()
	switchFunc3()
	deferFunc()
}
