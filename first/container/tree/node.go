package main

import (
    "fmt"
)

type treeNode struct {
   value int
   left, right *treeNode
}

func createNode(value int) *treeNode {
    return &treeNode{value: value}
}

func (node *treeNode) print() {
    fmt.Print(node.value, " ")
}

//结构体方法
func (node *treeNode) setValue(value int) {
    if node == nil {
        fmt.Println("Setting value to nil node. Ignored")
        return
    }
    node.value = value
}

func (node *treeNode) traverse() {
    if node == nil {
        return
    }
    node.left.traverse()
    node.print()
    node.right.traverse()
}

func main() {
    var root treeNode

    root = treeNode{value: 3}
    root.left = &treeNode{}
    root.right = &treeNode{5, nil, nil}
    root.right.left = new(treeNode)
    root.left.right = createNode(2)
    fmt.Println(root)

    nodes := []treeNode {
        {value: 3},
        {},
        {6, nil, &root},
    }
    fmt.Println(nodes)
    root.right.left.setValue(4)
    root.right.left.print()
    fmt.Println()

    //var pRoot *treeNode
    //pRoot.setValue(200)
    //pRoot = &root
    //pRoot.setValue(300)
    //pRoot.print()

    root.traverse()
}
