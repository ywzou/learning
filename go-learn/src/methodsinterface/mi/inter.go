package mi

import "fmt"

//接口interface
type iPerson interface {
	say() string
}

//结构体
type Person struct {
	name string
}

//实现接口方法
func (p Person) say() string {
	return p.name + " say hello!"
}

func describe(i iPerson)  {
	if nil == i {
		fmt.Println("iPerson is nil...")
		return
	}
	fmt.Printf("(%v, %T)\n", i, i)
}

func InterDemo()  {
	var i iPerson
	ps := Person{name:"张三"}
	i = ps
	fmt.Println(i.say())
	describe(i)//({张三}, mi.Person)

	//隐式实现
	var i2 iPerson = Person{name:"李四"}
	fmt.Println(i2.say())
	describe(i2)//({李四}, mi.Person)
}

//类型选择
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("int类型 值等于%v \n", v)
	case string:
		fmt.Printf("string类型 值等于%v \n", v)
	case iPerson:
		fmt.Printf("iPerson类型 值等于%v \n", v)
	default:
		fmt.Printf("未知的类型%T!\n", v)
	}
}

func InterDemo2()  {
	var i interface{} = "hello"

	// 类型断言
	s := i.(string)
	fmt.Println(s) //hello

	s, ok := i.(string)
	fmt.Println(s, ok)//hello true

	f, ok := i.(int)
	fmt.Println(f, ok)//0 false

	//f = i.(float64) // 报错(panic)

	do(21)
	do("hello")
	do(true)
	do(Person{name:"李四"})
}