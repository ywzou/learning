package concurrent

import (
	"fmt"
	"time"
)

//工作者
func worker(c chan int,id int)  {
	for  {
		//从channel接收值
		n := <-c;
		fmt.Printf("第 %d 个worker 从channel接收数据 %c \n", id, n)
	}
}

//创建worker
func createWorker(id int) chan int {
	c := make(chan int)
	// 去工作接收信息
	go worker(c, id)
	return c
}

//创建worker 返回的channel只能接收数据 只能向channel发送数据
func createWorkeSend(id int) chan<- int {
	c := make(chan int)
	// 去工作接收信息
	go worker(c, id)
	return c
}

//创建worker 返回的channel只能接收receive数据 不能像返回值发送数据
func createWorkeReceive(id int) <-chan int {
	c := make(chan int)
	// 去工作接收信息
	go func() {
		//n := <-c;
		c <- id
		fmt.Printf("第 %d 个worker 从channel接收数据 %c \n", id, <-c)
	}()
	return c
}

//带缓冲的信道
func bufChannel()  {
	c := make(chan int,2)
	c <- 1
	c <- 2
	fmt.Println(<-c)

}

//斐波那契
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	//关闭一个信道来表示没有需要发送的值了
	close(c)
}

//select 语句使一个 Go 程可以等待多个通信操作。
func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {

		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:// quit close 则退出
			fmt.Println("quit")
			return
		}
	}
}

func ChannelDemo()  {
	//创建一个信道
	c := make(chan int)
	//开启协程接收channel
	go func() {
		for {
			n := <- c
			fmt.Println(n)
		}
	}()

	//向channel传递值
	c <- 2
	c <- 4
	time.Sleep(time.Millisecond)

	fmt.Println("=====接收channel的函数=====")
	var channels [5]chan int
	for i:= 0; i<5; i++ {
		channels[i] = make(chan int)
		go worker(channels[i],i)
	}

	//向channel发送值
	for i:= 0; i<5; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)

	fmt.Println("======返回channel的函数======")
	var channels2 [5]chan int
	for i:= 0; i<5; i++ {
		channels2[i] = createWorker(i)
	}

	for i:= 0; i<5; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)

	fmt.Println("======只能向返回的channel的发送数据 不能向返回channel接收数据======")
	var channels3 [5]chan<- int
	for i:= 0; i<5; i++ {
		channels3[i] = createWorkeSend(i)
	}

	for i:= 0; i<5; i++ {
		channels3[i] <- 'A' + i

		//不能取数据 不能从channel取出数据 send-only
		//n := <- channels3[i]
	}

	time.Sleep(time.Millisecond)


	fmt.Println("======只能向返回channel接收数据  不能像返回的channel发送数据======")
	var channels4 [5]<-chan int
	for i:= 0; i<5; i++ {
		channels4[i] = createWorkeReceive(i)
	}

	for i:= 0; i<5; i++ {
		//不能像channel发送数据 只能接收数据 receive-only
		//channels4[i] <- 'A' + i

		n := <- channels4[i]
		fmt.Println(n)
	}

	time.Sleep(time.Millisecond)

	fmt.Println("======带缓冲的信道======")
	bufChannel()

	fmt.Println("======斐波那契======")
	c = make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("======斐波那契select======")
	c = make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciSelect(c, quit)
}