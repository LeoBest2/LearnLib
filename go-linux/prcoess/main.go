package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	var process []int
	var validId = regexp.MustCompile("^[0-9]+$")

	infoList, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Println(infoList)
	}

	for _, info := range infoList {
		if info.IsDir() && validId.MatchString(info.Name()) {
			p, _ := strconv.Atoi(info.Name())
			process = append(process, p)
		}
	}

	sort.Ints(process)

	statRe := regexp.MustCompile(`([0-9]+) \((.+?)\) [a-zA-Z]+ ([0-9]+)`)
	fmt.Printf("%6s\t%6s\t%s\n", "PID", "PPID", "NAME")
	for _, p := range process {
		b, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/stat", p))
		if err != nil {
			continue
		}
		matches := statRe.FindStringSubmatch(string(b))
		fmt.Printf("%6s\t%6s\t%s\n", matches[1], matches[3], matches[2])
	}
}
