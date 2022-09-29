package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	hostname, _ := os.Hostname()
	fmt.Println("程序看到的主机名是: ", hostname)
	fmt.Println("程序看到的文件系统是:")
	files, _ := ioutil.ReadDir("/")
	for _, f := range files {
		fmt.Printf("/%-10s\t%s\t%d字节\n", f.Name(), f.ModTime().Format("2006-01-02 15:04:05"), f.Size())
	}
	out, _ := exec.Command("ps", "-e", "-o", "pid,ppid,user,comm").Output()
	fmt.Println("程序看到的进程是:")
	fmt.Println(string(out))
	out, _ = exec.Command("ip", "addr").Output()
	fmt.Println("程序看到的网络是:")
	fmt.Println(string(out))
}
