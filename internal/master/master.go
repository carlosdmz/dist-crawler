package master

import (
	"log"
	"net/rpc"
)

type Reply struct {
	Status string
	Data string
}

func InitMaster(addr *string, seed *string) {
	log.Println("Master started.")
	node := callNode(*addr)

	var reply Reply

	err := node.Call("Crawler.Respond", seed, &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply.Status)

	err = node.Call("Crawler.Request", seed, &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply.Status)
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