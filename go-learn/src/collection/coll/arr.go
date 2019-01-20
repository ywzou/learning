package coll

import "fmt"

//数组
func ArrayDemo()  {
	//声明一个长度为2的数组
	var arr [2]string
	arr[0] = "Go"
	arr[1] = "Language"
	fmt.Println(arr) //[Go Language]
	fmt.Println(arr[0],arr[1]) //Go Language

	//遍历
	length := len(arr)
	for i := 0; i < length; i++ {
		fmt.Printf("遍历 index[%d], value = %s \n", i, arr[i])
	}

	//声明并赋值
	arr2 := [3]int{3, 7, 8}
	fmt.Println(arr2) //[3 7 8]

	//遍历
	for v := range arr2{
		fmt.Printf("range 遍历 value = %d \n", v)
	}

	//声明一个不定长度的数组
	arr3 := [...]int{3, 7, 8, 5, 6}
	fmt.Println(arr3) //[3 7 8]
}