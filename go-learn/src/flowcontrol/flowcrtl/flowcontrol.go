package flowcontrol

import (
	"fmt"
	"math"
	"time"
)

func FlowControl()  {
	step := 10
	loop(step)
	loop2(step)
	while(step)
	unlimitedLoop()
	compare(12)
	// 3*3 = 9  9 < 10  return 10
	// 3*3*3 = 27  27 > 20  return 20
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println()
	multiplicationTable()
	grade("M")
	unCondition()
	deferFun()
}

//基础语法
func loop(step int)  {
	sum := 0
	for i:= 0; i< step; i++  {
		sum += i
	}
	fmt.Println(sum)
}

//初始化语句和后置语句是可选的。
//将loop() 可简化为
func loop2(step int)  {
	sum := 1
	for ; sum < step; {
		sum += sum
	}
	fmt.Println(sum)
}

//for 是 Go 中的 “while”
func while(step int)  {
	sum := 1
	for sum < step {
		sum += sum
	}
	fmt.Println(sum)
}

//死循环
func unlimitedLoop()  {
	step := 1
	for {
		//循环20次结束
		if step > 20 {
			break
		}
		fmt.Printf("无限循环 step = %d \n",step)
		step += 1
	}
}

//if
func compare(x int)  {
	if x > 10 {
		fmt.Printf("x = %d 大于10 \n",x)
	}
}

//语句可以在条件表达式前执行一个简单的语句。该语句声明的变量作用域仅在 if 之内。
func pow(x, n, lim float64) float64 {
	//计算x的n次幂 如果结果小于lim 则返回结果，否则返回lim
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

// 九九乘法表
func multiplicationTable()  {
	for i := 1; i <= 9; i++  {
		for j := 1; j <= i; j++  {
			fmt.Printf("%d * %d = %d \t", j , i , i * j)
		}
		fmt.Println()
	}
}

//等级
func grade(gd string)  {
	switch gd {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	case "C":
		fmt.Println("中等")
	case "D":
		fmt.Println("及格")
	case "E":
		fmt.Println("差")
	default:
		fmt.Println("未知")
	}
}

//没有条件的switch
func unCondition()  {
	//当前时间
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("现在是早上")
	case t.Hour() < 18:
		fmt.Println("现在是下午")
	default:
		fmt.Println("现在是晚上")
	}
}

// defer
func deferFun()  {
	fmt.Println("方法开始")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("方法结束")
}