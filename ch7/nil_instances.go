/* 
	Code your methods for nil instances 

	Implementation of a binary tree that takes advantage of nil values for the receiver
*/

package main

import (
	"fmt"
)

type IntTree struct {
	val         int
	left, right *IntTree
}

// Insert a value into the binary search tree
func (it *IntTree) Insert(val int) *IntTree {
	// Check if the tree is empty (nil)
	if it == nil {
		return &IntTree{val: val}
	}
	// Insert to the left if the value is smaller
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		// Insert to the right if the value is greater
		it.right = it.right.Insert(val)
	}
	// Return the current node
	return it
}

// Check if the tree contains a given value
func (it *IntTree) Contains(val int) bool {
	// Handle nil node
	switch {
	case it == nil:
		return false
	case val < it.val:
		// Recursively search the left subtree
		return it.left.Contains(val)
	case val > it.val:
		// Recursively search the right subtree
		return it.right.Contains(val)
	default:
		// Value found
		return true
	}
}

func main() {
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)

	// Check if the tree contains certain values
	fmt.Println(it.Contains(2))  // Output: true
	fmt.Println(it.Contains(12)) // Output: false
}
