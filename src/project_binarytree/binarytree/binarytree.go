package binarytree

import(
	f "fmt"
	"errors"
)


// Node of a binary tree
type Node struct {
	Value int
	Parent *Node
	Left  *Node
	Right *Node
}

// InitBinaryTree makes a binary tree
func InitBinaryTree(value int) *Node {
	var newNode Node
	newNode.Value = value
	return &newNode
}

// FindNodeWithKey find the node with the key value
func FindNodeWithKey(root *Node, key int) *Node {
	if root == nil {
		return nil
	}

	if (*root).Value == key {
		return root
	} else if key < (*root).Value {
		return FindNodeWithKey((*root).Left, key)
	} else {
		return FindNodeWithKey((*root).Right, key)
	}
}

//DeleteNodeWithKey returns error if key is not present
func DeleteNodeWithKey(root *Node, key int) (*Node ,error){
	nodeToDelete := FindNodeWithKey(root, key)
	if nodeToDelete == nil{
		return nil, errors.New(f.Sprintf("Trying to delete a key that does not exist in the tree: %d\n", key))
	}

	if nodeToDelete == root{
		toTheLeft := (*root).Left
		toTheRight := (*root).Right
		if toTheLeft == nil && toTheRight != nil{
			(*toTheRight).Parent = nil
			return toTheRight, nil
			
		} else if toTheLeft != nil && toTheRight != nil {
			(*toTheLeft).Parent = nil
			return toTheLeft, nil

		} else{
			itt := toTheRight
			for i:=0;; i++ {
				if (*itt).Left == nil {break}
				itt = (*itt).Left
			}

			(*(*itt).Parent).Left = nil
			(*itt).Parent = nil

			(*itt).Left = toTheLeft
			(*toTheLeft).Parent = itt

			(*itt).Right = toTheRight
			(*toTheRight).Parent = itt
			return itt, nil
		}

	}

	isLeftChild := false
	if (*(*nodeToDelete).Parent).Left == nodeToDelete {
		isLeftChild = true
	}


	parent := (*nodeToDelete).Parent
	toTheLeft := (*nodeToDelete).Left
	toTheRight := (*nodeToDelete).Right


	if toTheLeft == nil && toTheRight == nil {
		if isLeftChild {
			(*parent).Left = nil
		
		}else{
			(*parent).Right = nil
		}

	}else if (*nodeToDelete).Left == nil && (*nodeToDelete).Right != nil {
		(*toTheRight).Parent = parent
		if isLeftChild {
			(*parent).Left = toTheRight
		
		}else {
			(*parent).Right = toTheRight

		}

	}else if toTheLeft != nil && toTheRight == nil {
		(*toTheLeft).Parent = parent
		if isLeftChild {
			(*parent).Left = toTheLeft
		}else {
			(*parent).Right = toTheLeft
		}
	}else {

		itt := toTheRight
		for i:=0;;i++ {
			if (*itt).Left == nil {
				break
			}
			itt = (*itt).Left
		}

		(*(*itt).Parent).Left = nil

		(*itt).Left = toTheLeft
		if toTheLeft != nil { (*toTheLeft).Parent = itt }
		
		(*itt).Right = toTheRight
		if toTheRight != nil { (*toTheRight).Parent = itt }

		(*itt).Parent = parent
		if isLeftChild {
			(*parent).Left = itt
		} else {
			(*parent).Right = itt
		}


	}

	return root, nil

}


// InsertIntoBinaryTree with an integer 
func InsertIntoBinaryTree(root *Node, value int) {
	itter := root
	for i:=0; ; i++{
		if value < (*itter).Value{
			if (*itter).Left == nil{
				newNode := Node{Value: value, 
								Parent: itter, 
								Left: nil,
								Right: nil}
				(*itter).Left = &newNode
				break
			} else{
				itter = (*itter).Left
			}
		} else{
			if (*itter).Right == nil{
				newNode := Node{Value: value,
								Parent: itter, 
								Left: nil, 
								Right: nil}
				(*itter).Right = &newNode
				break
			} else{
				itter = (*itter).Right
			}
		}
	}
}