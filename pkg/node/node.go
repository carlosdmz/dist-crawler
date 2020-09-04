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

func resolveAddr(addr string) (*net.TCPAddr, error) {
	resolve, err := net.ResolveTCPAddr("tcp", addr)
	return resolve, err
}

func TCPListen(network string, resolver *net.TCPAddr) (*net.TCPListener, error){
	server, err := net.ListenTCP(network, resolver)
	return server, err
}

func registerCrawler() error {
	crawlerType := new(crawler.Crawler)
	err := rpc.Register(crawlerType)
	return err
}

func startServer(addr string) {
	resolve, err := resolveAddr(addr)
	if err != nil {
		log.Panic("startServer:", err)
	}

	server, err := TCPListen("tcp", resolve)
	if err != nil {
		log.Panic("startServer:", err)
	}

	err = registerCrawler()
	if err != nil {
		log.Panic("startServer:", err)
	}

	log.Println("Node is started.")
	rpc.Accept(server)
}