package main

import (
	"testing"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

// Test case for function bst
func TestBST(t *testing.T) {
	// Redirect stdin for testing input
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	r, w, _ := os.Pipe()
	os.Stdin = r

	// Write test input
	inputs := []string{
		"1", // Insert
		"5", // Value to insert
		"2", // Search
		"5", // Value to search
		"3", // Exit
	}
	go func() {
		for _, input := range inputs {
			fmt.Fprintln(w, input)
		}
	}()

	bst()

	// Reset stdin
	w.Close()
}

// Test case for function bst
func TestBST(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "Invalid choice", input: "a\n3", expected: "Invalid choice. Please enter a number."},
		{name: "Insert value successfully", input: "1\n5\n3", expected: "Value inserted successfully."},
		{name: "Search - value not found", input: "2\n5\n10\n3", expected: "Value not found."},
		{name: "Search - value found", input: "1\n5\n2\n2\n5\n3", expected: "Value found."},
		{name: "Exit program", input: "3", expected: "Exiting program."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r, w, _ := os.Pipe()
			os.Stdin = r
			w.WriteString(tt.input)

			buffer := new(strings.Builder)
			w2 := bufio.NewWriter(buffer)
			os.Stdout = w2

			bst()

			w.Close()
			w2.Flush()

			output := buffer.String()
			if !strings.Contains(output, tt.expected) {
				t.Errorf("Expected: %s, Got: %s", tt.expected, output)
			}

			os.Stdin = os.Stdin
			os.Stdout = os.Stdout
		})
	}
}

// Test case for function bst
func TestBST(t *testing.T) {
	// Redirect input to test the function
	input := "1\n5\n2\n5\n3\n"
	reader := bufio.NewReader(strings.NewReader(input))
	os.Stdin = reader

	bst()
}

func TestInvalidChoiceError(t *testing.T) {
	// Redirect input to test the function with an invalid choice
	input := "4\n"
	reader := bufio.NewReader(strings.NewReader(input))
	os.Stdin = reader

	bst()
}

func TestInsertValue(t *testing.T){
	// Redirect input to test inserting a value
	input := "1\n5\n3\n"
	reader := bufio.NewReader(strings.NewReader(input))
	os.Stdin = reader

	bst()
}

func TestSearchValue(t *testing.T){
	// Redirect input to test searching a value
	input := "2\n5\n3\n"
	reader := bufio.NewReader(strings.NewReader(input))
	os.Stdin = reader

	bst()
}

func TestSearchValueNotFound(t *testing.T){
	// Redirect input to test searching a value that doesn't exist
	input := "2\n10\n3\n"
	reader := bufio.NewReader(strings.NewReader(input))
	os.Stdin = reader

	bst()
}

func TestInvalidValue(t *testing.T){
	// Redirect input to test inserting an invalid value
	input := "1\nabc\n3\n"
	reader := bufio.NewReader(strings.NewReader(input))
	os.Stdin = reader

	bst()
}

func TestTreeEmpty(t *testing.T){
	// Redirect input to test searching when the tree is empty
	input := "2\n5\n3\n"
	reader := bufio.NewReader(strings.NewReader(input))
	os.Stdin = reader

	bst()
}

func TestExitProgram(t *testing.T){
	// Redirect input to test exiting the program
	input := "3\n"
	reader := bufio.NewReader(strings.NewReader(input))
	os.Stdin = reader

	bst()
}

// Test case for function bst
func TestBST(t *testing.T) {
	t.Run("Insert and Search", func(t *testing.T) {
		reader := bufio.NewReader(os.Stdin)
		value := 10

		root := &Node{Key: value}

		root.Insert(5)
		root.Insert(15)

		searchValue := 5
		found := root.Search(searchValue)
		if !found {
			t.Errorf("%d should be found in the tree", searchValue)
		}

		searchValue2 := 20
		notFound := root.Search(searchValue2)
		if notFound {
			t.Errorf("%d should not be found in the tree", searchValue2)
		}
	})

	t.Run("Invalid Input", func(t *testing.T) {
		reader := bufio.NewReader(os.Stdin)

		root := &Node{}

		valueInput := "invalid"
		insertInput := "invalid"
		forceSearchValue := 0

		_, err := strconv.Atoi(valueInput)
		if err == nil {
			t.Error("Validation for invalid value input failed")
		}

		root.Insert(10)
		_, err = strconv.Atoi(insertInput)
		if err == nil {
			t.Error("Validation for invalid insert input failed")
		}

		if root.Search(forceSearchValue) {
			t.Error("Should not find value in an empty tree")
		}
	})
}

