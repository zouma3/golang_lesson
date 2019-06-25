package main

import (
	"fmt"
	"os/user"
	"strconv"
	"time"
)

func init() {
	fmt.Println("Hallo World", time.Now())
	fmt.Println(user.Current())
	fmt.Println()
}

/*
	変数定義練習
*/
func VariableDefinition() {
	var (
		i   int     = 1
		f64 float64 = 1.2
		s   string  = "test"
		t   bool    = true
		f   bool    = false
	)
	fmt.Println(i, f64, s, t, f)

	i2 := 1
	f642 := 1.2
	s2 := "test"
	t2 := true
	f2 := false
	fmt.Println(i2, f642, s2, t2, f2)
	fmt.Printf("%T %T %T %T %T\n", i2, f642, s2, t2, f2)

	var s3 string = "14"
	i3, _ := strconv.Atoi(s3)
	fmt.Println(i3)

	const Pi = 3.14
	fmt.Println(Pi)
}

func array() {
	var a [2]int = [2]int{100, 200}
	//a[0] = 100
	//a[1] = 200
	fmt.Println(a)

	n := []int{1, 2, 3, 4, 5}
	fmt.Println(n)
}

func add(x int, y int) (int, int) {
	p := func() {
		fmt.Println("add")
	}
	p()
	return x + y, x - y
}

func main() {
	//VariableDefinition()
	//array()
	//x, y := add(1, 2)
	//fmt.Println(x, y)

}
