package main

import (
	"github.com/carlosdamazio/dist-crawler/internal/master"
	"github.com/carlosdamazio/dist-crawler/internal/node"
	"github.com/carlosdamazio/dist-crawler/internal/standalone"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"os"
)


var(
	app         			= kingpin.New("dist-crawler", "Distributed Web Crawler.")

	standaloneMode  		= app.Command("standalone", "Standalone mode.")
	seedStandAloneMode		= standaloneMode.Arg("seed", "Initial URL to be called upon.").Required().String()

	masterMode  			= app.Command("master", "Master mode.")
	seedMasterMode			= masterMode.Arg("seed", "Initial URL to be called upon.").Required().String()

	nodeMode    			= app.Command("node", "Node mode.")
	masterAddr  			= nodeMode.Flag("masterAddr", "Master host and port.").Required().String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case masterMode.FullCommand():
		master.InitMaster()
	case nodeMode.FullCommand():
		node.InitNode(*masterAddr)
	case standaloneMode.FullCommand():
		standalone.InitStandAlone()
	}
}