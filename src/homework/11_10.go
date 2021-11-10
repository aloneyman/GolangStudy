package main

import (
	"bytes"
	"fmt"
	"strings"
)

func prefix(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
func command(s string, idx int) string {
	if len(s) <= 3 {
		return s
	}
	return s[:idx+3] + "." + command(s[idx+3:], idx+3)
	//return comma(s[:n-3]) + "," + s[n-3:]
}
func SorP(s string, sub string) bool {
	return strings.HasSuffix(s, sub) && strings.HasPrefix(s, sub)
}

func commandBuffer(val []int) string {
	var buf bytes.Buffer
	for i, value := range val {
		if i%3 == 0 && i > 0 {
			buf.WriteString(".")
		}
		fmt.Fprintf(&buf, "%d", value)
	}
	return buf.String()
}
func sub(s string, str string) bool {
	return strings.Contains(s, str)
}

func main() {
	//fmt.Println(prefix("abcd/chh.hahah"))
	//fmt.Println(command("123456",0))
	//fmt.Println(commandBuffer([]int{1,2,3,4,5}))
	fmt.Println(sub("abcde", "cd"))
}
