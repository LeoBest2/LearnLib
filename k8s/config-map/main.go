/*
	GOOS=linux GOARCH=386 go build -ldflags '-s -w' -o cm-demo
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("LEO-NAME: " + os.Getenv("LEO-NAME"))
	fmt.Println("LEO-AGE: " + os.Getenv("LEO-AGE"))
}
