package concurrent

import "fmt"

type TreeNode struct {
	value int
	leftNode,rightNode *TreeNode
}

//结构体创建 工厂函数
func createNode(value int) *TreeNode {
	return &TreeNode{value:value}
}

func (node *TreeNode) traverseFun(f func(*TreeNode))  {
	if nil == node {
		return
	}
	f(node)
	node.leftNode.traverseFun(f)
	node.rightNode.traverseFun(f)
}

func (node *TreeNode) traverse() chan *TreeNode{
	c := make(chan *TreeNode)
	go func() {
		node.traverseFun(func(node *TreeNode) {
			c <- node
			fmt.Println(node.value)
		})
		close(c)
	}()
	return c
}

func ChannelTraverseTree()  {
	//创建结构体 方式二
	root := TreeNode{value:-1}

	//创建结构体 方式三
	//root := new(treeNode)

	//赋值
	root.leftNode = &TreeNode{value:1}
	root.rightNode = &TreeNode{value:2}

	root.leftNode.leftNode = createNode(11)
	root.leftNode.rightNode = createNode(12)

	root.rightNode.leftNode = createNode(21)
	root.rightNode.rightNode = createNode(22)

	fmt.Println(root)

	c := root.traverse()
	maxValue := 0
	for node := range c{
		if node.value > maxValue {
			maxValue = node.value
		}
	}

	fmt.Printf("最大的值是%d \n",maxValue)//22
}