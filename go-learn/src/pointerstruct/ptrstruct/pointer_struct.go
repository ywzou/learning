package ptrstruct

import "fmt"

//定义一个结构体 定义一个点
type Point struct {
	x int
	y int
}

//指针
func PointerDemo()  {
	//声明两个变量
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值 42

	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值 21

	p = &j         // 指向 j 2701
	*p = *p / 37   // 通过指针对 j 进行除法运算 2701/37
	fmt.Println(j) // 查看 i 的值 73
}

//结构体的测试
func PointDemo()  {
	//声明一个结构体
	var point = Point{2,3}
	fmt.Println(point) //{2 3}

	//声明一个结构体
	var point2 = Point{}
	fmt.Println(point2) //{0 0}

	//结构体字段使用.来访问
	point2.x = 10 //给x赋值
	point2.y = 5
	fmt.Println(point2, point2.x, point2.y) //{10 5} 10 5
}

//结构体指针
func StructPointer() {
	fmt.Println("结构体指针...")

	//声明结构体
	point := Point{2,3}
	fmt.Println(point) //{2 3}

	p := &point //指针 取得point内存
	p.x = 1024
	fmt.Println(point) //{1024 3}

	point3 := &Point{1, 2}
	fmt.Println(point3) //&{1 2}
}