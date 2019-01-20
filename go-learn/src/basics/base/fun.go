package basics

//加法 返回 x + y
func Add(x int,y int) int {
	return x + y
}

//减法 返回 x - y
func Sub(x int,y int) int {
	return x - y
}

//乘法 返回 x * y
func Multi(x int,y int) int {
	return x * y
}

//除法 返回 x / y
func Div(x int,y int) int {
	return x / y
}

//函数作为参数
func Customize(x, y int, cust func(int ,int) int) int{
	return cust(x,y)
}

//多个返回值
func Swap(x, y int) (int, int){
	return Div(x,y), x % y
}

//多个返回值 命名
func Swap2(x, y int) (a int, b int){
	a = Div(x,y)
	b = x % y
	return
}