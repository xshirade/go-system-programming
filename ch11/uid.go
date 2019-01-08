package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("User ID: %d\n", os.Getuid())
	fmt.Printf("Group ID: %d\n", os.Getgid())
	groups, _ := os.Getgroups()
	fmt.Printf("Subgroup ID: %v\n", groups)
	fmt.Printf("User ID: %d\n", os.Geteuid())
	fmt.Printf("Group ID: %d\n", os.Getegid())
}
