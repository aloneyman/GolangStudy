package main

import (
	"fmt"
	_ "fmt"
	_ "os"
)

func getChh() int{
	return 1024
}

func main() {
	chh := getChh()
	fmt.Printf("%d",chh)
}