package main

import (
	"fmt"
    "os"
    "github.com/carlosdamazio/dist-crawler/pkg/standalone"
    "github.com/carlosdamazio/dist-crawler/pkg/node"
    "github.com/carlosdamazio/dist-crawler/pkg/master"
    "github.com/carlosdamazio/dist-crawler/pkg/utils"
)

func main() {
    if len(os.Args) <= 1 {
        fmt.Fprintf(os.Stderr, "Error: insufficient arguments for program\n\n")
        utils.PrintUsage()
        os.Exit(1)
	}
    args := utils.InitArgs()

    switch args["mode"] {
        case "main":
            if args["node"] == "" || args["seed"] == "" {
                fmt.Fprintf(os.Stderr, "Error: invalid argument set\n")
                utils.PrintMainNodeUsage()
                goto FINISH
            }
            
            master.InitMaster(args["node"], args["seed"]) 
        case "node":
            // no checks, just ignore additional flags
            node.InitNode()
        case "standalone":
            if args["seed"] == "" {
                fmt.Fprintf(os.Stderr, "Error: invalid argument set\n")
                utils.PrintMainNodeUsage()
                goto FINISH
            }

            standalone.InitStandAlone()
        default:
            fmt.Fprintf(os.Stderr, "Error: invalid crawler mode\n\n")
            utils.PrintUsage()
            goto FINISH
    }
FINISH:
    os.Exit(1)
}
