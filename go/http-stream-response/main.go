// ~~~~~~~~~
// http stream response
// Author: Leo
// Usage: go run main.go
// 		  curl localhost:8088
// 可看到每三秒输出一次Hello World!
package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		for i := 0; i < 3; i++ {
			w.Write([]byte("Hello World!\r\n"))
			w.(http.Flusher).Flush()
			time.Sleep(time.Second * 3)
		}
	})

	http.ListenAndServe(":8088", nil)
}
