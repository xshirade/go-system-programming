package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func open() {
	file, err := os.Create("textfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "New file content\n")
}

func read() {
	file, err := os.Open("textfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("Read file:")
	io.Copy(os.Stdout, file)
}

func append() {
	file, err := os.OpenFile("textfile.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "Appended content\n")
}

func mkdir() {
	os.Mkdir("setting", 0755)
	os.MkdirAll("setting/myapp/networksettings", 0755)
}

func rm() {
	os.Remove("textfile.txt")
	os.RemoveAll("setting")
}

func truncate() {
	os.Truncate("textfile.txt", 20)
	// file, _ := os.Open("textfile.txt")
	// file.Truncate(20)
}

func main() {
	open()
	read()
	for i := 0; i < 5; i++ {
		append()
	}
	read()
	truncate()
	read()
	mkdir()
	// fmt.Println("info: sleep for 1 second")
	time.Sleep(1 * time.Second)
	rm()
}
