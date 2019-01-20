package coll

import "fmt"

type User struct {
	name string
	age int
}

func MapDemo()  {
	//声明一个映射 键类型为string 值类型为int
	var m = map[string]int{
		"age": 20,
		"height" : 180,
	}

	fmt.Println(m)//map[age:20 height:180]
	fmt.Println(m["age"], m["height"], m["width"])//20 180 0

	//遍历
	for k,v := range m {
		fmt.Printf("key = %s , val = % d \n" ,k , v)
	}

	m2 := map[string]User{
		"A": {name:"张三", age: 20},
		"B": {name:"李四", age: 80},
	}
	fmt.Println(m2)//map[A:{张三 20} B:{李四 80}]

	//获取元素
	elem := m2["A"]
	fmt.Println(elem)//{张三 20}
	fmt.Println(elem.age,elem.name)//20 张三

	//有值才取出
	if elem,ok := m2["C"]; ok {
		fmt.Println(elem)
	}

	//修改值
	m2["A"] = User{name:"王麻子",age:34}
	fmt.Println(m2)//map[A:{王麻子 34} B:{李四 80}]

	//删除元素
	delete(m2,"A")
	fmt.Println(m2)//map[B:{李四 80}]
}