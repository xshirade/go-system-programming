package main

import "net"

func main() {
	conn, err := net.Dial("unix", "socketfile")
	if err != nil {
		panic(err)
	}
	// connを使った読み書き
}
