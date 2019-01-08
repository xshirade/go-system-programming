package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

func tree(path string, prefix string, currentDepth int, maxDepth int) {
	if currentDepth == maxDepth {
		return
	}
	fileInfo, err := os.Lstat(path)
	if err != nil {
		panic(err)
	}
	if fileInfo.Mode().IsDir() == false {
		return
	}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(fileInfo.Mode())
		panic(err)
	}
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
	sort.SliceStable(_names, func(i, j int) bool { return _names[i] < _names[j] })
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

func main() {
	help := flag.Bool("help", false, "Print usage and this help message and exit.")
	maxDepth := flag.Int("L", -1, "Descend only level directories deep.")
	flag.Parse()
	if *help == true {
		fmt.Println("usage: tree [-L level] [<directory list>]")
		flag.PrintDefaults()
		return
	}
	paths := []string{}
	if flag.NArg() == 0 {
		paths = append(paths, ".")
	} else {
		paths = append(paths, flag.Args()...)
	}
	for _, path := range paths {
		fileInfo, err := os.Stat(path)
		if os.IsNotExist(err) == true {
			fmt.Printf("%s [error opening dir]\n", path)
		} else if err != nil {
			panic(err)
		} else {
			if fileInfo.IsDir() == true {
				fmt.Printf("%s\n", path)
				tree(path, "", 0, *maxDepth)
			}
		}
	}
}
