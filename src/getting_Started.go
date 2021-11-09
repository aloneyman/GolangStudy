package main

import "fmt"

const boilingF = 212.0

func incr(p *int) int {
	*p++
	return *p

}
func fio(n int) int {
	x, y := 0, 1
	for i := 1; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
func main() {
	//x := 1
	//p := &x
	//fmt.Printf("%d\n",p)
	//fmt.Printf("%d\n",*p)
	//*p = 2
	//fmt.Printf("%d\n",*p)
	//v := 1
	//incr(&v)
	//fmt.Println(v)
	//fmt.Println(incr(&v))
	//var n = flag.Bool("n", false, "omit trailing newline")
	//var sep = flag.String("s", " ", "separator")
	//flag.Parse()
	//fmt.Print(strings.Join(flag.Args(), *sep))
	//fmt.Println()
	//fmt.Println(*n)
	//
	//fmt.Println(fio(4))
	//fmt.Println(byte(0&1))
	//x := "hello"
	//for _, x := range x {
	//	x := x + 'A' - 'a'
	//	fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
	//}
	//fmt.Println()

	y := 4
	x := 3
	if x := f(); x == 0 {
		fmt.Println("first")
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println("second")
		fmt.Println(x, y)
	} else {
		fmt.Println("third")
		fmt.Println(x, y)
	}
	fmt.Println(x, y)
}
func f() int {
	return 3
}
func g(x int) int {
	return x
}
