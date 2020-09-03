package node

import (
	"github.com/carlosdamazio/dist-crawler/pkg/crawler"
	"log"
	"net"
	"net/rpc"
)

func InitNode() {
	log.Println("Node server starting...")
	startServer("0.0.0.0:13372")
}

func startServer(addr string) {
	resolve, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatal("startServer:", err)
	}

	server, err := net.ListenTCP("tcp", resolve)
	if err != nil {
		log.Fatal("startServer:", err)
	}

	crawlerType := new(crawler.Crawler)
	err = rpc.Register(crawlerType)

	if err != nil {
		log.Fatal("startServer:", err)
	}

	log.Println("Node is started.")
	rpc.Accept(server)
}