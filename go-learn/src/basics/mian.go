package main

import (
	"basics/base"
	"fmt"
	"math"
)

//x % y
// x y 均是 int类型
func mol(x int,y int) int {
	return x % y
}

//声明变量 全局变量不能使用 := 声明
var name = "GoLang"
var c, python, java bool

//声明常量
const Pi = 3.14

func main() {
	//声明变量 并初始化
	var index = 20

	fmt.Println(index)
	fmt.Println(name,c,python,java)
	fmt.Println(Pi)

	fmt.Printf("9 的平方根是 %g \n", math.Sqrt(9))

	//声明变量
	x := 10
	y := 3
	f := mol

	fmt.Printf("x + y = %d \n",basics.Add(x,y))
	fmt.Printf("x - y = %d \n",basics.Sub(x,y))
	fmt.Printf("x * y = %d \n",basics.Multi(x,y))
	fmt.Printf("x / y = %d \n",basics.Div(x,y))
	fmt.Printf("x 余 y = %d \n",basics.Customize(x,y,f))

	a, b := basics.Swap(x, y)
	fmt.Printf("a = %d ,b = %d \n",a , b)


}