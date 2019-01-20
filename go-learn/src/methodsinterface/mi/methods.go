package mi

import "fmt"

//定义一个简单的二叉树结构体
type treeNode struct {
	value int
	leftNode,rightNode *treeNode
}

//结构体的方法 接收者为值
func (node treeNode) print()  {
	fmt.Println(node.value)
}

//设置是 接收者为指针
func (node *treeNode) setValue(value int)  {
	node.value = value
}

//遍历二叉树
func (node *treeNode) traverse()  {
	if nil == node {
		return
	}
	node.print()
	node.leftNode.traverse()
	node.rightNode.traverse()
}

//结构体创建 工厂函数
func createNode(value int) *treeNode {
	return &treeNode{value:value}
}

//结构体创建演示
func StructDemo()  {
	//创建结构体 方式一
	//var root treeNode

	//创建结构体 方式二
	root := treeNode{value:-1}

	//创建结构体 方式三
	//root := new(treeNode)

	//赋值
	root.leftNode = &treeNode{value:1}
	root.rightNode = &treeNode{value:2}

	root.leftNode.leftNode = createNode(11)
	root.leftNode.rightNode = createNode(12)

	root.rightNode.leftNode = createNode(21)
	root.rightNode.rightNode = createNode(22)

	fmt.Println(root)

	//调用结构体方法
	root.setValue(0)
	root.print()
	fmt.Println("===== 二叉树遍历 ====")
	root.traverse()
}