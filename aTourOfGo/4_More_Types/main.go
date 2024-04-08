package main

import (
	"fmt"
)

func pointerFunc() {
	i, j := 42, 2701

	p := &i
	fmt.Println("pointer addrress: ", &i)
	// dereferencing for pointer ポインターの逆参照
	fmt.Println("dereferencing for pointer: ", *p)
	fmt.Println("p(pointer address): ", p)

	*p = 21
	fmt.Println("i: ", i)

	p = &j
	*p = *p / 37
	fmt.Println("j: ", j)
}

type Vertex struct {
	X int
	Y int
}

func structFunc() {
	v := Vertex{1, 2}
	fmt.Println("struct X: ", v.X)
	p := &v
	(*p).X = 10
	p.X = 30
	fmt.Println("struct pointer address: ", &p)
	fmt.Println("struct change X: ", v.X)
	fmt.Println("struct Y: ", v.Y)
}

func structLiteral() {
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p  = &Vertex{1, 2}
	)
	fmt.Println("structLiteral v1: ", v1)
	fmt.Println("structLiteral v2: ", v2)
	fmt.Println("structLiteral v3: ", v3)
	fmt.Println("structLiteral p: ", p)
}

func arrayFunc() {
	var arr1 [2]string
	arr2 := [2]string{"Golang", "Ruby"}
	arr3 := [...]string{"Golang", "Ruby"}

	arr1[0] = "Golang"
	arr1[1] = "Ruby"
	fmt.Println("arrayFunc arr1: ", arr1)
	fmt.Println("arrayFunc arr2: ", arr2)
	fmt.Println("arrayFunc arr3: ", arr3)
}

func sliceFunc() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	s := primes[1:4]
	fmt.Println("sliceFunc s: ", s)
}

func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("sliceLiterals q: ", q)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("sliceLiterals s: ", s)
}

func sliceDefault() {
	s := []int{2, 3, 5, 7, 11, 13}

	ss := s[1:4]
	fmt.Println("sliceDefault ss: ", ss)

	sss := s[:2]
	fmt.Println("sliceDefault sss: ", sss)

	ssss := s[1:]
	fmt.Println("sliceDefault ssss: ", ssss)

	sssss := s[:]
	fmt.Println("sliceDefault sssss: ", sssss)
}

func sliceLengthCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("sliceLengthCapacity len: ", len(s))
	fmt.Println("sliceLengthCapacity cap: ", cap(s))

	c := s[0:2]
	fmt.Println("sliceLengthCapacity c: ", cap(c))
}

func makeSlice() {
	a := make([]int, 5)
	fmt.Println("makeSlice a", a, "len: ", len(a), "cap: ", cap(a))
	b := make([]int, 0, 5)
	fmt.Println("makeSlice b", b, "len: ", len(b), "cap: ", cap(b))
	capp := 10
	c := make([]int, 5, capp)
	fmt.Println("makeSlice c", c, "len: ", len(c), "cap: ", cap(c))
	d := append(a, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("makeSlice d", d, "len: ", len(d), "cap: ", cap(d))
}

func rangeFunc() {
	iterable := []int{1, 2, 3, 4, 5}

	for i, v := range iterable {
		fmt.Printf("rangeFunc iterable[%d]: %d\n", i, v)
	}
}

func mapFunc() {
	type Vertex struct {
		Lat, Long float64
	}

	m := map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println("mapFunc m: ", m)
	elem, ok := m["Bell Lab"]
	fmt.Println("mapFunc elem: ", elem, "ok: ", ok)
}

func main() {
	pointerFunc()
	structFunc()
	structLiteral()
	arrayFunc()
	sliceFunc()
	sliceLiterals()
	sliceDefault()
	sliceLengthCapacity()
	var nilSlice []int
	fmt.Println("nilSlice: ", nilSlice)
	fmt.Println("nil?: ", nilSlice == nil)
	makeSlice()
	rangeFunc()
	mapFunc()
}
