# Dist-crawler
[![carlosdamazio](https://circleci.com/gh/carlosdamazio/dist-crawler.svg?style=svg)](https://circleci.com/gh/carlosdamazio/dist-crawler) [![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org) [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens)

A proposal for a distributed web crawler built upon RPC communication protocol.

## What is it, really?
Well, I'm learning Go. So, instead of building a CRUD application, I decided to build something that I had in mind. It's a dead simple Web Crawler that works on a Master -> Node architecture. It's currently in development, so I have a few remarks on it:

- It has a single node feature. In the future, the crawler must accept more than one node connection to make parallel processing and load balancing;
- It's not using the concurrency features that Go has to offer;
- It uses RPC communication. Maybe in the future I might use gRPC to make the most of the enhanced version of the protocol;
- It should sanitize the URLs extracted from the pages, it comes with lots of trash (assets, mail:to, etc);
- It might have a database to store indexed pages and it's contents, think about storing the whole page or slice it if you want to;
- Someday must have an UI application that can detect running instances and show their status.

## How to use
I know, this thing does not have tests since it's my little toy, a breakable one. But if you want to adventure yourself with it, be sure to have Go 1.14, since it's the version that I used to build it.

```bash
# Download the project
go get github.com/carlosdamazio/dist-crawler

# Start the node agent
$ go build -o dist-crawler ./cmd/dist-crawler/main.go
$ ./dist-crawler node

# On another instance, run the master
$ ./dist-crawler master --nodesAddr=localhost:13372 http://damazio.dev
