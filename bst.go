package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Node represents a node in the binary search tree.
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert inserts a new node with the given key into the BST.
func (n *Node) Insert(key int) {
	if key < n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	} else if key > n.Key {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	}
}

// Search searches for a node with the given key in the BST.
func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}

	if key < n.Key {
		return n.Left.Search(key)
	} else if key > n.Key {
		return n.Right.Search(key)
	}

	return true
}

func bst() {
	var root *Node

	fmt.Println("Binary Search Tree Operations")
	fmt.Println("1. Insert")
	fmt.Println("2. Search")
	fmt.Println("3. Exit")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter your choice: ")
		choiceInput, _ := reader.ReadString('\n')
		choiceInput = strings.TrimSpace(choiceInput)

		choice, err := strconv.Atoi(choiceInput)
		if err != nil {
			fmt.Println("Invalid choice. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter the value to insert: ")
			valueInput, _ := reader.ReadString('\n')
			valueInput = strings.TrimSpace(valueInput)
			value, err := strconv.Atoi(valueInput)
			if err != nil {
				fmt.Println("Invalid value. Please enter a number.")
				continue
			}
			if root == nil {
				root = &Node{Key: value}
			} else {
				root.Insert(value)
			}
			fmt.Println("Value inserted successfully.")
		case 2:
			fmt.Print("Enter the value to search: ")
			searchInput, _ := reader.ReadString('\n')
			searchInput = strings.TrimSpace(searchInput)
			searchValue, err := strconv.Atoi(searchInput)
			if err != nil {
				fmt.Println("Invalid value. Please enter a number.")
				continue
			}
			if root == nil {
				fmt.Println("Tree is empty.")
			} else {
				if root.Search(searchValue) {
					fmt.Println("Value found.")
				} else {
					fmt.Println("Value not found.")
				}
			}
		case 3:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
