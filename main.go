package main

import (
	_ "embed"
	"fmt"
)

//go:embed example.env
var content string

func main() {
	tokenizer := NewTokenizer(content)

	tokens := tokenizer.ReadAll()
	fmt.Println(tokens)
}
