package main

import (
	"fmt"
	"strings"
)

func main() {
	var name = `resume`
	fmt.Println(len(name))
	var profession = []rune("developer")
	fmt.Println(len(profession))
	var stringBuilder strings.Builder
	stringBuilder.WriteString("Hello ")
	stringBuilder.WriteString(name)
	stringBuilder.WriteString("! You are a ")
	stringBuilder.WriteString(string(profession))
	stringBuilder.WriteString(".")
	fmt.Println(stringBuilder.String())
}
