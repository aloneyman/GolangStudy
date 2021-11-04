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

func main() {
	chh := &user{ // 传递地址
		name: "chh",
		id:   1,
	}
	fmt.Println(chh)
	add_id(chh)
	fmt.Println(chh)
}
