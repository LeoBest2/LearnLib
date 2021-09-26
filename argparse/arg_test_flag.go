package main

import (
	"flag"
	"fmt"
	"os"
)

type allowTypes struct {
	values []string
	count  int
}

func (t *allowTypes) String() string {
	return fmt.Sprint(t.values)
}

var (
	src      string
	dst      string
	types    allowTypes
	silent   bool
	verbose  bool
	maxCount int
)

func (t *allowTypes) Set(value string) error {
	t.values = append(t.values, value)
	t.count += 1
	return nil
}

func main() {
	flag.StringVar(&src, "src", "", "原始目录")
	flag.StringVar(&dst, "dst", "", "原始目录")
	flag.Var(&types, "type", "文件扩展名, 可多个参数如:-type txt -type md")
	flag.BoolVar(&silent, "s", false, "不打印移动详情, 与-v不能同时使用")
	flag.BoolVar(&verbose, "v", false, "打印移动详情, 与-s不能同时使用")
	flag.IntVar(&maxCount, "m", 10000, "最大移动的数量")
	flag.Parse()

	if src == "" || dst == "" || types.count == 0 {
		flag.Usage()
	} else if silent == false && verbose == false {
		// 默认为-v
		verbose = true
	} else if silent && verbose {
		_, _ = fmt.Fprintln(os.Stderr, "-v 和 -s 参数不能同时使用")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println(src, dst, types, silent, verbose, maxCount)
	if verbose {
		fmt.Printf("\n正在从文件夹: %s 移动类型: %s 文件到文件夹: %s , 移动时开启打印: %v ,最大移动数量: %d",
			src, types.String(), dst, verbose, maxCount)
	} else {
		fmt.Printf("\n正在从文件夹: %s 移动类型: %s 文件到文件夹: %s , 移动时关闭打印: %v ,最大移动数量: %d",
			src, types.String(), dst, silent, maxCount)
	}
}
