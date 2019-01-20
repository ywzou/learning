package main

import (
	"fmt"
	"thread/concurrent"
)

func main()  {
	fmt.Println("VVV")
	//concurrent.RoutineDemo()
	//concurrent.ChannelDemo()
	//concurrent.SelectDemo()
	//concurrent.WaitGroupDemo()
	concurrent.ChannelTraverseTree()
}