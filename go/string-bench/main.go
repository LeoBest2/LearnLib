package main

import (
	"bytes"
	"fmt"
	"strings"
)

func builder() {
	var sb strings.Builder
	sb.WriteString("hello")
	sb.WriteString("world")
	_ = sb.String()
}

func buffer() {
	var buff bytes.Buffer
	buff.WriteString("hello")
	buff.WriteString("world")
	_ = buff.String()
}

func sprinf() {
	_ = fmt.Sprintf("%s%s", "hello", "world")
}

func join() {
	s := []string{"hello", "world"}
	_ = strings.Join(s, "")
}

func add() {
	_ = "hello" + "world"
}
