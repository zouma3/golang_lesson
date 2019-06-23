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
}

func main(){
	VariableDefinition()
}