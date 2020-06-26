package master

import (
	"log"
	"net/rpc"
)

func InitMaster(addr string) {
	log.Println("Master started.")
	node := callNode(addr)

	var reply bool
	err := node.Call("Crawler.Respond","", &reply)
	if err != nil {
		log.Fatal(err)
	}
}

func electNode() {
	log.Println("Electing a node...")
}

func callNode(addr string) *rpc.Client {
	// Might elect a node to seed...
	electNode()

	log.Println("Checking if a node responds...")
	node, err := rpc.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	return node
}