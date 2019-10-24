package main

import (
	f "fmt"
	bt "project1/binarytree"
	"project1/utils"
	// "bufio"
	// "io"
	// "os"
	// "io/ioutil"
)

func main() {
	f.Println("Hello world!")

	root := bt.InitBinaryTree(5)
	bt.InsertIntoBinaryTree(root, 2)
	bt.InsertIntoBinaryTree(root, 18)
	bt.InsertIntoBinaryTree(root, -4)
	bt.InsertIntoBinaryTree(root, 3)

	// utils.PrintTree(&root)

	f.Println(bt.FindNodeWithKey(root, 2))
	root, err := bt.DeleteNodeWithKey(root, 2)
	utils.CheckError(err)
	f.Println(bt.FindNodeWithKey(root, 2))



	f.Println("main terminated")

}
