package concurrent

import (
	"fmt"
	"runtime"
	"time"
)

func threadFun(index int)  {
	for  {
		//Printf是io操作 或有协程测切换 主动交出资源
		fmt.Printf("第 %d 个线程....\n",index)
	}
}


func RoutineDemo()  {
	//开启10个线程
	for i := 0; i < 10; i ++ {
		go threadFun(i)
	}
	time.Sleep(time.Millisecond) //休眠一毫秒

	var arr [5]int
	for i := 0; i < 5; i ++ {
		go func(i int) {
			for {
				arr[i] ++
				//手动交出协程
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond) //休眠一毫秒
	fmt.Println(arr)//[632 791 755 545 785] 并不是平均的
}