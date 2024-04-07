package main

import (
	"fmt"
	"math"
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

func main() {
	forFunc()
	forContinued()
	ifFunc(2)
	PowPrint(3, 2, 10)
	PowPrint(3, 3, 20)
}
