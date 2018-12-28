package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func tree(path string, prefix string) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	file, err := os.Open(absolutePath)
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
				tree(path+"/"+name, prefix+"    ")
			} else {
				fmt.Printf("%s├── %s\n", prefix, name)
				tree(path+"/"+name, prefix+"│   ")
			}
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: tree path")
		os.Exit(0)
	}
	relativePath := os.Args[1]
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		panic(err)
	}
	fileInfo, err := os.Stat(absolutePath)
	if os.IsNotExist(err) == true {
		panic(err)
	} else if err != nil {
		panic(err)
	} else {
		if fileInfo.IsDir() == true {
			fmt.Printf("%s\n", relativePath)
			tree(os.Args[1], "")
		}
	}
}
