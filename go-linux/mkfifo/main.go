/* ~~~~~~~~~
 学习了解linux 命名管道
 Author: Leo
 Usage: go run main.go
	& echo "some text" > /tmp/go-fifo-demo
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"
)

const PATH = "/tmp/go-fifo-demo"

func main() {
	err := syscall.Mkfifo(PATH, 0644)
	if err != nil && !os.IsExist(err) {
		panic("Mkfifo failed: " + err.Error())
	}

	log.Printf("fifo path is : %s\n", PATH)

	f, err := os.Open(PATH)
	if err != nil {
		panic(fmt.Sprintf("open %s failed: %v", PATH, err))
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for {
		s.Scan()
		msg := s.Text()
		if len(msg) != 0 {
			log.Printf("recived: %s %d\n", msg, len(msg))
		}
	}
}
