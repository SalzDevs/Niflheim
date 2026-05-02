package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_go "github.com/tree-sitter/tree-sitter-go/bindings/go"
	tree_sitter_java "github.com/tree-sitter/tree-sitter-java/bindings/go"
	tree_sitter_javascript "github.com/tree-sitter/tree-sitter-javascript/bindings/go"
	tree_sitter_python "github.com/tree-sitter/tree-sitter-python/bindings/go"
)

type ParseResult struct {
	Tree *tree_sitter.Tree
	Code []byte
	Language *tree_sitter.Language
}


func LangDetector(filePath string) (string, error) {
	ext := filepath.Ext(filePath)
	
	switch ext {
	case ".js":
		return "javascript", nil
	case ".ts":
		return "typescript", nil
	case ".py":
		return "python", nil
	case ".go":
		return "go", nil
	case ".java":
		return "java", nil
	default:
		return "", fmt.Errorf("unsupported file type: %s", ext)
	}
}

func ParserWrapper(filePath string) (*ParseResult, error) {
	langName , err := LangDetector(filePath) 
	if err != nil {
		fmt.Printf("Skipping unsupported file type: %s\n", filePath)
		return nil, nil
	}

	code, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	parser := tree_sitter.NewParser()
	defer parser.Close()

	var language *tree_sitter.Language
	switch langName {
	case "javascript":
		language = tree_sitter.NewLanguage(tree_sitter_javascript.Language())
	case "python":
		language = tree_sitter.NewLanguage(tree_sitter_python.Language())
	case "go":
		language = tree_sitter.NewLanguage(tree_sitter_go.Language())
	case "java":
	 	language = tree_sitter.NewLanguage(tree_sitter_java.Language())
	default:
		return nil, fmt.Errorf("unsupported language: %s", langName)
	}

	if err := parser.SetLanguage(language); err != nil {
		return nil, fmt.Errorf("failed to set language: %w", err)
	}

	tree := parser.Parse(code, nil)
	if tree == nil {
		return nil, fmt.Errorf("failed to parse file: %s", filePath)
	}

	return &ParseResult{
		Tree: tree,
		Code: code,
		Language: language,		
	}, nil
}

func FileWalker(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && !strings.HasPrefix(info.Name(), ".") {	
			return filepath.SkipDir
		}

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil 
}

