package crawler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

    xurls "mvdan.cc/xurls"
)

type Crawler int

type Reply struct {
	// This is a struct to be passed when replying messages from RPC functions in nodes to the master.
	Status string
	Data string
}

type CrawlReply struct {
	//The crawl reply passes an array of multiple urls, whereas reply alone passes messages only.
	Status string
	Data []string
}

func (c *Crawler) Respond(seed string, reply *Reply) error {
	//Returns a message to the master that node is alive
	reply = &Reply{Status: fmt.Sprintf("Node Crawler is listening and ready to crawl its initial seeds like: %s", seed), Data: ""}
	log.Println("Master called initial response.")
	return nil
}

func (c *Crawler) Request(url string, reply *Reply) error {
	//Node makes a request for URL in frontier to grab data. It returns the response's body.
	log.Printf("Master requested to crawl %s frontier...", url)
	res, err := http.Get(url)
	if err != nil {
	    goto ERROR
	}

	if res.StatusCode < http.StatusBadRequest {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
		    goto ERROR
        }
		bodyString := string(bodyBytes)
		reply.Status = "OK"
        reply.Data = bodyString

	} else {
		log.Printf("Crawler.Request: when processing %s: HTTP status code was different than OK (200)", url)
		return errors.New("dist-crawler: HTTP Code was different than 200")
	}
	
    err = res.Body.Close()
	if err != nil {
		log.Fatalf("dist-crawler: could not close body reader")
		return err
	}
    return nil

ERROR:
    log.Fatalf("dist-crawler: when processing %s: %v", url, err)
    return err
    
}

func (c *Crawler) Crawl(data string, reply *CrawlReply) error {
    strict := xurls.Strict
    reply.Data = strict.FindAllString(data, -1)
    log.Println(reply.Data)
    return nil
}
