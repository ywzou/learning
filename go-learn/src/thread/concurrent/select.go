package concurrent

import (
	"fmt"
	"math/rand"
	"time"
)

func genChannel() chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			c <- i
			i++
		}
	}()
	return c
}

func SelectDemo()  {
	var c1, c2 = genChannel(),genChannel()
	//非阻塞是的接收值
	//三秒结束
	tm := time.After(time.Second * 3);
	for {
		select {
		case n := <-c1 :
			fmt.Printf("从c1接收到值%d \n",n)
		case n := <-c2 :
			fmt.Printf("从c2接收到值%d \n",n)
		case <-tm:
			fmt.Printf("程序执行结束....\n")
			return
		//default:
		//	fmt.Printf("没有值\n")
		}
	}
}