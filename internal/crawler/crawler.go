package crawler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Crawler int

type Reply struct {
	Status string
	Data string
}


func (c *Crawler) Respond(seed string, reply *Reply) error {
	*reply = Reply{Status: fmt.Sprintf("Node Crawler is listening and ready to crawl its initial seeds like: %s", seed), Data: ""}
	log.Println("Master called initial response.")
	return nil
}

func (c *Crawler) Request(url string, reply *Reply)  error {
	log.Printf("Master requested to crawl %s frontier...", url)
	res, err := http.Get(url)

	if err != nil {
		log.Fatalf("Crawler.Request: when processing %s: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("Crawler.Request: when processing %s: %v", err)
			return err
		}

		bodyString := string(bodyBytes)
		*reply = Reply{Status: "OK", Data: bodyString}
		log.Println(reply.Status)
		return nil
	} else {
		log.Fatalf("Crawler.Request: when processing %s: HTTP status code was different than OK (200)", err)
		return nil
	}

	return nil
}
/*
func (c *Crawler) Crawl() []string {

}
 */