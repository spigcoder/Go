package main

import "fmt"

type f int
type m int

type Data struct {
	value int
}

func escapeExample() *Data {
    // 这里的局部变量 d 将逃逸到堆上
    d := Data{value: 42}
	fmt.Printf("Address of d in escapeExample: %p\n", &d)  // 打印 d 的地址
    return &d
}

func main() {
	var x f = 1
	var y m = 2

	x = f(y)
	if x == 2{
		fmt.Println("x == 2")
	} 
}