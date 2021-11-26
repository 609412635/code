package main

import (
	"fmt"
)

type TreeNode struct {
	Item string
	Left *TreeNode
	Right *TreeNode
}
type bst struct {
	root *TreeNode
}


/*
        m
     k     l
  h    i     j
a  b  c  d  e  f


//先序遍历（根左右）：m k h a b i c d l j e f

//中序遍历（左根右）：a h b k c i d m l e j f

//后序遍历（左右根）：a b h c d i k e f j l m

*/
func (tree *bst) buildTree() {

	m := &TreeNode{Item: "m"}
	tree.root = m

	k := &TreeNode{Item: "k"}
	l := &TreeNode{Item: "l"}
	m.Left = k
	m.Right = l

	h := &TreeNode{Item: "h"}
	i := &TreeNode{Item: "i"}
	k.Left = h
	k.Right = i

	a := &TreeNode{Item: "a"}
	b := &TreeNode{Item: "b"}
	h.Left = a
	h.Right = b

	c := &TreeNode{Item: "c"}
	d := &TreeNode{Item: "d"}
	i.Left = c
	i.Right = d

	j := &TreeNode{Item: "j"}
	l.Right = j

	e := &TreeNode{Item: "e"}
	f := &TreeNode{Item: "f"}

	j.Left = e
	j.Right = f

}

//先序遍历

func first(tree *TreeNode)  {
	if tree != nil {
		fmt.Println(tree.Item)
		first(tree.Left)
		first(tree.Right)
	}
}

func mid(tree *TreeNode)  {
	if tree != nil {
		mid(tree.Left)
		fmt.Println(tree.Item)
		mid(tree.Right)
	}
}

func last(tree *TreeNode)  {
	if tree != nil {
		last(tree.Left)
		last(tree.Right)
		fmt.Println(tree.Item)
	}
}

func main() {

	tree := &bst{}

	tree.buildTree()

	last(tree.root)
	//a:=firstTraversal(tree.root)
	//fmt.Println(a)
	a:=lastTraversal(tree.root)
	fmt.Println(a)
	//tree.inOrder()
	//tree.lastOrder()

}

// 先序遍历栈实现
func firstTraversal(root *TreeNode)[]string  {
	result := make([]string, 0)
	for root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) >0 {
		if root != nil {
			result = append(result,root.Item)
			stack = append(stack,root)
			root = root.Left
		} else {
			val := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = val.Right
		}

	}
	return result
}

func midTraversal(root *TreeNode)[]string  {
	result := make([]string, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack,root)
			root = root.Left
		} else {
			val := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result,val.Item)
			root = val.Right
		}
	}

	return result

}

// 后序遍历，要入栈两次，第二次弹出时候，才输出节点输出节点时刻，右节点没有数据或者右节点已经弹出过了
func lastTraversal(root *TreeNode)[]string  {
	result := make([]string, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	var last *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}

		node :=stack[len(stack)-1]
		if node.Right == nil || node.Right == last {
			stack = stack[:len(stack)-1]
			result = append(result,node.Item)
			last = node
		} else {
			root = node.Right
		}

	}
	return result
}