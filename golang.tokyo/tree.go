package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func tree(path string, prefix string, currentDepth int, maxDepth int) {
	if currentDepth == maxDepth {
		return
	}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	if fileInfo.IsDir() == true {
		names, err := file.Readdirnames(0)
		if err != nil {
			panic(err)
		}
		re := regexp.MustCompile(`^\.`)
		_names := []string{}
		for _, name := range names {
			if re.MatchString(name) == false {
				_names = append(_names, name)
			}
		}
		for index, name := range _names {
			if len(_names) == index+1 {
				fmt.Printf("%s└── %s\n", prefix, name)
				tree(filepath.Join(path, name), prefix+"    ", currentDepth+1, maxDepth)
			} else {
				fmt.Printf("%s├── %s\n", prefix, name)
				tree(filepath.Join(path, name), prefix+"│   ", currentDepth+1, maxDepth)
			}
		}
	}
}

func main() {
	maxDepth := flag.Int("L", -1, "Descend only level directories deep.")
	flag.Parse()
	relativePath := ""
	if flag.NArg() == 0 {
		relativePath = "."
	} else if flag.NArg() > 1 {
		fmt.Println("usage: tree [-L level] [directory]")
		return
	} else {
		relativePath = os.Args[len(os.Args)-1]
	}
	fileInfo, err := os.Stat(relativePath)
	if os.IsNotExist(err) == true {
		fmt.Println(err)
	} else if err != nil {
		panic(err)
	} else {
		if fileInfo.IsDir() == true {
			fmt.Printf("%s\n", relativePath)
			tree(relativePath, "", 0, *maxDepth)
		}
	}
}
