package concurrent

import (
	"fmt"
	"sync"
)

//等待
func doTaskWG(id int,c chan int, wg *sync.WaitGroup)  {
	for n := range c{
		fmt.Printf("任务接收第 %d 个任务，值为 %c \n", id, n)
		wg.Done()
	}
}

type taskWG struct {
	c chan int
	wg *sync.WaitGroup
}

func createTaskWG(id int,wg *sync.WaitGroup) taskWG{
	t := taskWG{
		c: make(chan int),
		wg: wg,
	}
	go doTaskWG(id, t.c, wg)
	return t
}

func WaitGroupDemo()  {
	var wg sync.WaitGroup
	var workers [5]taskWG
	for i:=0; i<5; i++  {
		workers[i] = createTaskWG(i, &wg);
	}

	wg.Add(10)//10个任务
	for i, worker := range workers  {
		worker.c <- 'A' + i
		//wg.Add(1) //循环一次加一个任务
	}

	//fmt.Println("===============")
	for i, worker := range workers  {
		worker.c <- 'a' + i
		//wg.Add(1) //循环一次加一个任务
	}

	//等待任务结束
	wg.Wait()
}