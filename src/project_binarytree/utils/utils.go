package utils

import (
	f "fmt"
	bt "project1/binarytree"
)


// CheckError halts the goroutine and prints the error
func CheckError(e error){
	if e != nil{
		f.Println(e)
		panic(e)
	}
}

// PrintTree by iterating the tree in the BFS style
func PrintTree(root *bt.Node) {
	queue := make([]*bt.Node, 1)
	queue[0] = root

	for i:=0;;i++{
		childrenCount := 0
		for j:=0; j < len(queue); j++{
			if (*queue[j]).Left != nil { childrenCount ++ }
			if (*queue[j]).Right != nil { childrenCount ++}
		}

		newQueue := make([]*bt.Node, childrenCount)
		newQueueIndex := 0
		for j:=0; j < len(queue); j++{
			if queue[j].Left != nil { 
				newQueue[newQueueIndex] = (*queue[j]).Left
				newQueueIndex ++ 
			}
			if queue[j].Right != nil { 
				newQueue[newQueueIndex] = (*queue[j]).Right
				newQueueIndex ++
			}
			
			f.Print((*queue[j]).Value, "\t")
		}

		f.Println()
		if childrenCount == 0{
			break
		}
		queue = newQueue

	}
}
