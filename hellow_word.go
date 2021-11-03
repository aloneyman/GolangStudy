package main

import (
	"fmt"
	_ "fmt"
	_ "os"
)

func getChh() int {
	return 1024
}

func main() {
	power := 1000
	fmt.Printf("default power is %d\n", power)

	/*
		尽管变量power使用了:=，但是编译器不会在第2次使用:=时报错，因为这里有一个变量name，这是一个新的变量，
		允许使用:=。但是你不能改变power的类型。它已经被声明成一个整型，只能赋值整数。
	*/
	name, power := "Goku", 9000
	fmt.Printf("%s's power is over %d\n", name, power)

	/*
			现在，你只需要记住使用var NAME TYPE声明一个变量，并且变量的初始值为它相应类型的零值，使用NAME := VALUE
		声明一个变量并赋值，使用NAME = VALUE去给已经声明过的变量赋值。
	*/

}
