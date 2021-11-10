package main

import "fmt"

//type Currency int
//
//const (
//	USD Currency = iota // 美元
//	EUR                 // 欧元
//	GBP                 // 英镑
//	RMB                 // 人民币
//)
//
//symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
//
//fmt.Println(RMB, symbol[RMB]) // "3 ￥"
func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

//var s []int    // len(s) == 0, s == nil
//s = nil        // len(s) == 0, s == nil
//s = []int(nil) // len(s) == 0, s == nil
//s = []int{}    // len(s) == 0, s != nil
func rotate(arr []int, r int) []int {
	lens := len(arr)
	res := make([]int, lens)
	for i := 0; i < lens; i++ {
		idx := i + r
		if idx >= lens {
			idx -= lens
		}
		res[i] = arr[idx]
	}
	return res
}
func emptyString(str []string) []string {
	i := 0
	for _, val := range str {
		idx := i + 1
		if idx >= len(str) {
			break
		}
		if val != str[idx] {
			str[i] = val
			i++
		}
	}
	return str[:i]
}
func main() {
	//arr := [...]int{1,2,3,4}
	//fmt.Println(arr)
	//arr := [...]int{5:-1} // 定一个了一个容量为6的数组最后一个数是-1
	//fmt.Println(arr)
	//var arr = [...]int{1,2,3,4,5}
	//reverse(arr[2:])
	//fmt.Println(arr)
	//var x []int
	//x = append(x,1)
	//x = append(x , 2)
	//x = append(x,x...)
	//fmt.Println(x)
	//fmt.Println(nonempty([]string{"zrx","chh"}))
	//fmt.Println(len(nonempty([]string{"zrx","chh"})))
	//data := []string{"one", "", "three"}
	//fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	//fmt.Printf("%q\n", data)           // `["one" "three" "three"]` 第二个位置传递引用所以改变一起改变
	//fmt.Println(remove([]int{1,2,4,5,6},1))

}
func equ(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for xk, xv := range x {
		if yv, ok := x[xk]; !ok || xv == yv {
			return false
		}
	}
	return true
}
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	fmt.Println(slice)
	return slice[:len(slice)-1]
}
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
