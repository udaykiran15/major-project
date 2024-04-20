package main

import (
	"testing"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/v1/openai"
	"reflect"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"context"
	"regexp"
	"fmt"
	"io/ioutil"
)

// Test case for function createTestFile
type testCase struct {
	Name    string
	Imports []string
	Code    string
}

func Test_createTestFile(t *testing.T) {
	// Test case with existing test file
	path := "example.go"
	packageName := "test"
	testCases := []testCase{
		{
			Name: "TestFunction",
			Imports: []string{
				"fmt",
				"testing",
			},
			Code: "// Test case code goes here",
		},
	}
	err := createTestFile(path, packageName, testCases)
	if err != nil {
		t.Errorf("Error creating test file: %v", err)
	}

	// Test case without existing test file
	path = "example2.go"
	err = createTestFile(path, packageName, testCases)
	if err != nil {
		t.Errorf("Error creating test file: %v", err)
	}

	// Test case with error when reading existing test file
	path = "nonexistent.go"
	err = createTestFile(path, packageName, testCases)
	if err == nil {
		t.Error("Expected error when reading nonexistent test file, but got nil")
	}
}

// Test case for function generateTestCases
func TestGenerateTestCases(t *testing.T) {
	ctx := context.Background()
	client := openai.NewClient("your_api_key")
	path := "path/to/your/go/file.go"
	model := "gpt-3.5-turbo"

	generateTestCases(ctx, client, path, model)

	// Add assertions or further testing logic here if needed
}

// Test case for function chatGPTTestCases
func TestChatGPTTestCases(t *testing.T) {
	ctx := context.Background()
	client := openai.NewClient("your-api-key")

	packageName := "testpackage"
	functionName := "testFunction"
	functionCode := "func testFunction() { return \"test\" }"
	model := openai.GPT3Dot5Turbo

	sanitizedCode, importContent, err := chatGPTTestCases(ctx, client, packageName, functionName, functionCode, model)

	if err != nil {
		t.Fatalf("Failed to generate test case: %v", err)
	}

	// Add your assertions here based on the expected output of the sanitizeCode function
	if len(sanitizedCode) == 0 {
		t.Errorf("Sanitized code is empty")
	}

	if len(importContent) == 0 {
		t.Errorf("Import content is empty")
	}
}

// Test case for function sanitizeCode
func TestSanitizeCode(t *testing.T) {
	testCases := []struct {
		name          string
		rawCode       string
		expectedCode  string
		expectedImports []string
	}{
		{
			name: "No code block",
			rawCode: "This is a test code",
			expectedCode: "This is a test code",
			expectedImports: []string{},
		},
		{
			name: "Code block without package and imports",
			rawCode: "```go\nfunc test() {\nfmt.Println(\"Test\")\n}\n```",
			expectedCode: "func test() {\nfmt.Println(\"Test\")\n}",
			expectedImports: []string{},
		},
		{
			name: "Code block with package and imports",
			rawCode: "```go\npackage main\nimport (\n\"fmt\"\n)\nfunc test() {\nfmt.Println(\"Test\")\n}\n```",
			expectedCode: "func test() {\nfmt.Println(\"Test\")\n}",
			expectedImports: []string{"fmt"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			code, imports, err := sanitizeCode(tc.rawCode)
			
			if err != nil {
				t.Errorf("sanitizeCode returned an error: %v", err)
			}

			if code != tc.expectedCode {
				t.Errorf("Expected code: %s, got: %s", tc.expectedCode, code)
			}

			if len(imports) != len(tc.expectedImports) {
				t.Errorf("Expected number of imports: %d, got: %d", len(tc.expectedImports), len(imports))
			}

			for i, imp := range imports {
				if imp != tc.expectedImports[i] {
					t.Errorf("Expected import: %s, got: %s", tc.expectedImports[i], imp)
				}
			}
		})
	}
}

// Test case for function findImportBlock
func TestFindImportBlock(t *testing.T) {
	testCases := []struct {
		input string
		expectedOutput string
	}{
		{
			input: "import \"fmt\"",
			expectedOutput: "import \"fmt\"",
		},
		{
			input: "import (\n    \"fmt\"\n    \"os\"\n)",
			expectedOutput: "import (\n    \"fmt\"\n    \"os\"\n)",
		},
		{
			input: "package main\nfunc main() {}",
			expectedOutput: "",
		},
		{
			input: "import (\n    \"fmt\"\n)",
			expectedOutput: "import (\n    \"fmt\"\n)",
		},
		{
			input: "import (\"fmt\")",
			expectedOutput: "import (\"fmt\")",
		},
	}

	for _, tc := range testCases {
		output := findImportBlock(tc.input)
		if output != tc.expectedOutput {
			t.Errorf("findImportBlock(%s) = %s, expected %s", tc.input, output, tc.expectedOutput)
		}
	}
}

// Test case for function extractImports
func TestExtractImports(t *testing.T) {
	tests := []struct {
		name          string
		importBlock   string
		expectedSlice []string
	}{
		{
			name:          "Single line import with package alias",
			importBlock:   `import openai "github.com/sashabaranov/go-openai"`,
			expectedSlice: []string{"openai \"github.com/sashabaranov/go-openai\""},
		},
		{
			name:          "Single line import without package alias",
			importBlock:   `import "fmt"`,
			expectedSlice: []string{"\"fmt\""},
		},
		{
			name:          "Multi-line imports with comments",
			importBlock:   `import (
				"fmt"
				"log"
				"github.com/some/package"
				// this is a comment, it should be ignored
				)`,
			expectedSlice: []string{"\"fmt\"", "\"log\"", "\"github.com/some/package\""},
		},
		{
			name:          "Empty import block",
			importBlock:   "",
			expectedSlice: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractImports(tt.importBlock)

			if !reflect.DeepEqual(got, tt.expectedSlice) {
				t.Errorf("ExtractImports() = %v, want %v", got, tt.expectedSlice)
			}
		})
	}
}

// Test case for function createTestFile
func TestCreateTestFile(t *testing.T) {
	// Test case with existing test file and new imports
	path := "example.go"
	packageName := "main"
	testCases := []testCase{
		{
			Name:    "TestFunc1",
			Imports: []string{"fmt"},
			Code:    "// Test Func1 code here",
		},
		{
			Name:    "TestFunc2",
			Imports: []string{"math/rand"},
			Code:    "// Test Func2 code here",
		},
	}
	err := createTestFile(path, packageName, testCases)
	if err != nil {
		t.Errorf("Error creating test file: %v", err)
	}

	// Test case with new test file creation
	newPath := "new_example.go"
	newTestCases := []testCase{
		{
			Name:    "TestNewFunc1",
			Imports: []string{"time"},
			Code:    "// Test NewFunc1 code here",
		},
		{
			Name:    "TestNewFunc2",
			Imports: []string{"strings"},
			Code:    "// Test NewFunc2 code here",
		},
	}
	err = createTestFile(newPath, packageName, newTestCases)
	if err != nil {
		t.Errorf("Error creating test file: %v", err)
	}

	// Test case with no test cases
	emptyPath := "empty.go"
	emptyCases := []testCase{}
	err = createTestFile(emptyPath, packageName, emptyCases)
	if err != nil {
		t.Errorf("Error creating test file: %v", err)
	}

	// Clean up test files
	os.Remove(testFileName)
	os.Remove(filepath.Join(filepath.Dir(newPath), strings.TrimSuffix(filepath.Base(newPath), ".go")+"_test.go")
	os.Remove(filepath.Join(filepath.Dir(emptyPath), strings.TrimSuffix(filepath.Base(emptyPath), ".go")+"_test.go")
}

