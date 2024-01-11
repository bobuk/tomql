package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
	"strings"
)

func getValueFromToml(config *toml.Tree, keys []string) (interface{}, error) {
	current := config
	for _, key := range keys {
		next := current.Get(key)
		if next == nil {
			return nil, fmt.Errorf("Section %s not found", strings.Join(keys, "."))
		}
		current = next.(*toml.Tree)
	}
	return current, nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: tomql <filename.toml> <section.param>")
		os.Exit(1)
	}

	filename := os.Args[1]
	key := os.Args[2]

	config, err := toml.LoadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error of reading file or file not found: %v\n", err)
		os.Exit(1)
	}

	parts := strings.Split(key, ".")

	if len(parts) < 2 {
		fmt.Fprintln(os.Stderr, "Incorrect format of section.param. Use .param for default section.")
		os.Exit(1)
	}

	sectionKeys := parts[:len(parts)-1]
	param := parts[len(parts)-1]

	value, err := getValueFromToml(config, sectionKeys)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	paramValue := value.(*toml.Tree).Get(param)

	if paramValue == nil {
		fmt.Fprintln(os.Stderr, "section or param not found")
		os.Exit(1)
	}

	fmt.Println(paramValue)
}
