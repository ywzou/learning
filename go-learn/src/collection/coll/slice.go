package coll

import "fmt"

//切片
func SliceDemo()  {
	fmt.Println("切片 slice ...")
	//声明一个切片
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)//[1 2 3 4 5 6]

	// 切片遍历
	for i, v := range arr {
		fmt.Printf("index = %d, value = %d\n", i, v)
	}

	// 切片遍历 下标或值赋予 '_' 来忽略它
	for _, v := range arr {
		fmt.Printf("value = %d\n", v)
	}

	// 获取切片
	var slice []int = arr[:] //全部元素
	fmt.Println(slice)//[1 2 3 4 5 6]

	slice = arr[:4] // [0,4)元素
	fmt.Println(slice)//[1 2 3 4]

	slice = arr[3:] // [3:5] 元素
	fmt.Println(slice) //[4 5 6]

	slice = arr[1:4] // 索引 (1,4] 元素
	fmt.Println(slice) //[2 3 4]

	//声明一个slice
	var slice2 []int

	//像slice其添加值
	for i := 0; i < 10; i++ {
		printSlice(slice2)
		slice2 = append(slice2, i*2);
	}

	//声明一个长度len为3 类型为int的 slice
	slice3 := make([]int, 3)
	printSlice(slice3)

	//声明一个长度len为3 cap为8 类型为int的 slice
	slice4 := make([]int, 3, 8)
	printSlice(slice4)

	//将slice copy到slice3
	copy(slice3, slice)
	fmt.Println(slice3)//[2 3 4]

	//删除下标为2的值
	slice3 = deleteElement(slice3,2)
	fmt.Println(slice3)//[2 4]

	//修改下标为1的值
	updateElement(slice3,1, 100)
	fmt.Println(slice3)//[2 100]
}

// 修改slice的值
func updateElement(s []int, index int, val int){
	s[index] = val
}

//删除元素 s 要删除的原始数据 index为要删除的索引下标
func deleteElement(s []int, index int) []int {
	low := index - 1
	return append(s[:low],s[index:]...)
}

//注意 这个类型 []int 并不是数组 而是 切片slice
// len 长度
// cap slice的内部机制 可理解为预留长度 容量
func printSlice(s []int)  {
	fmt.Printf("value = %v , len = %d , cap = %d \n" , s, len(s), cap(s))
}