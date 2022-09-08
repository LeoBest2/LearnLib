/*
	GOOS=linux GOARCH=386 go build -ldflags '-s -w' -o webserver
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	log.Fatalln(http.ListenAndServe(":80", nil))
}
