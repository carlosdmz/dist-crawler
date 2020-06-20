package node

import "fmt"

func InitNode(master string) {
	fmt.Printf("Node started, pointing to master at %s", master)
}