package main

import (
	"fmt"
	"os/user"
	"time"
)

func init(){
	fmt.Println("Hallo World",time.Now())
	fmt.Println(user.Current())
	fmt.Println()
}

/*
	変数定義練習
 */
func VariableDefinition(){
	var (
		i int = 1
		f64 float64 = 1.2
		s string = "test"
		t bool = true
		f bool = false
	)
	fmt.Println(i,f64,s,t,f)

	i2 := 1
	f642 := 1.2
	s2 := "test"
	t2 := true
	f2 := false
	fmt.Println(i2,f642,s2,t2,f2)

	const Pi = 3.14
	fmt.Println(Pi)
}

func main(){
	VariableDefinition()
}