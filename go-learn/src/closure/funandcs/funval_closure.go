package funandcs

import "fmt"

//值函数和闭包

//计算 传入一个需要两个int类型的参数并返回值为int的函数  没有返回值
func computeFn(x int, y int ,fn func(int, int) int) {
	fmt.Println(fn(x,y))
}

//传入一个需要两个int类型的参数并返回值为int的函数  放回一个int类型值
func computeFn2(x int, y int ,fn func(int, int) int) int {
	return fn(x,y);
}

//定义一个 需要两个int类型的参数并返回值为int的函数类型
type computeParam func(int, int) int

func computeFnSim(x int, y int ,fn computeParam) {
	fmt.Println(fn(x,y))
}

//定义一个需要连个int类型参数 且放回值为一个不需要参数且返回值为int的函数
func computeFn3(x int, y int) func() int{
	return func() int {
		return x * y
	}
}

func computeFn4(x int, y int) func(int) int{
	return func(z int) int {
		return x * y / z
	}
}

//闭包 斐波纳契闭包 (0, 1, 1, 2, 3, 5, ...)`。
//	1 	1 	2 	3 	5 	8 	13 	21 	34 	55
//  	a	b
//			a	b
func fibonacci() func() int {
	a,b := 0,1
	return func() int {
		a,b = b, a+b
		return a
	}
}

// i到n的和
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

//定义一个函数类型 需要一个int类型参数 返回一个int类型值和一个函数 递归
type IAdder func(int) (int,IAdder)

//相对标准的函数是编程 不能有过程变量
func adder2(base int) IAdder {
	return func(v int) (int, IAdder) {
		return base + v , adder2(base + v)
	}
}

func Compute()  {
	addFn := func(x int, y int) int {
		return x + y
	}

	computeFn(5,3,addFn) //8
	computeFnSim(11,3,addFn) //14

	subFn := func(x int, y int) int {
		return x - y
	}
	fmt.Println(computeFn2(5,3,subFn)) //2

	f :=computeFn3(5,3);
	fmt.Println(f()) //15

	f2 :=computeFn4(5,3);
	fmt.Println(f2(2)) //7


	pos := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("1+...+%d = %d \n",i,pos(i))
	}

	fmt.Println("========")
	pos2 := adder2(0)
	var sum int
	for i := 0; i < 10; i++ {
		sum,pos2 = pos2(i);
		fmt.Printf("0+...+%d = %d \n",i,sum)
	}

	//10个斐波纳契
	f3 := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d \t",f3())
	}
}