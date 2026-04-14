package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"morego/algorithms/parsing"
)

func main() {
	filePath := flag.String("file", "algorithms/parsing/ast.go", "Go source file to parse")
	flag.Parse()

	if err := run(*filePath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filePath string) error {
	src, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("read %s: %w", filePath, err)
	}

	functions, err := parsing.ParseFunctionInfo(string(src))
	if err != nil {
		return fmt.Errorf("parse %s: %w", filePath, err)
	}

	fmt.Printf("File: %s\n", filePath)
	for _, fn := range functions {
		fmt.Printf("Function: %s\n", fn.Name)
		if fn.Receiver != "" {
			fmt.Printf("Receiver: %s\n", fn.Receiver)
		}
		fmt.Printf("Parameters: %d\n", fn.ParameterCount)
		fmt.Printf("Results: %d\n", fn.ResultCount)
		if len(fn.Calls) == 0 {
			fmt.Println("Calls: none")
		} else {
			fmt.Printf("Calls: %s\n", strings.Join(fn.Calls, ", "))
		}
		fmt.Println()
	}

	return nil
}
