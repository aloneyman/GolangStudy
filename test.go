// 结构体

package main

import (
	"fmt"
	_ "fmt"
)

type user struct {
	name string
	id   int
}

func add_id(use *user) { // 显然 *user 和 user 是有关系但是不同种类类型的变量
	use.id += 10
}

func (use *user) add_id() { // 将user 和 add_id 函数进行绑定  *use是add_id方法的接收者
	use.id += use.id
}

func change_user(use *user) {
	//use = &user{"yyc", 1} // 不用取地址符
	use.name = "yyc" //是use拷贝过来的时候就是拷贝的指针
}

func main() {
	chh := &user{ // 传递地址
		name: "chh",
		id:   1,
	}
	//fmt.Println(chh)
	//add_id(chh)
	//fmt.Println(chh)
	//change_user(chh)
	//fmt.Println(chh)
	chh.add_id()
	fmt.Println(chh)
}
