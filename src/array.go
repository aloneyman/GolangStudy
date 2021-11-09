package main

import (
	"fmt"
	"os"
)

//在go中，数组是固定大小的。
func main() {
	//var score [30]int
	//// var score := [4]int{1,2,3,4}  直接赋值
	//score[0] = 1;
	//for index , val := range score{
	//	fmt.Printf("%d %d\n",index,val)
	//}

	score := make([]int, 0, 10)
	for i, arg := range os.Args[1:] {
		fmt.Printf("{} {}", i, arg)
	}
	score = score[0:5]
	score = append(score, 5)
	//add_score := []int{1, 2, 3, 4}
	//score = append(score, add_score...)
	fmt.Println(score)
}
