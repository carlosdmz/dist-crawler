package utils

import (
	"fmt"
    "flag"
	"os"
)

func InitArgs() map[string]string {
    args := make(map[string]string)

    var mode = flag.String("mode", "standalone", "Web crawler mode")
    var seed = flag.String("s" , "", "Initial URL to crawl (seed)")
    var node = flag.String("n", "", "IP/FQDN of node")

    flag.Parse()

    args["mode"] = *mode
    args["seed"] = *seed
    args["node"] = *node

    return args
}

// Print main command usage
func PrintUsage() {
    fmt.Fprintf(os.Stdout, "dist-crawler - Distributed Web Crawler\n\n")
    fmt.Fprintf(os.Stdout, "Usage: ./dist-crawler [-h --help] mode ")
    fmt.Fprintf(os.Stdout, "[-s seed -n nodeAddr]\n")
    fmt.Fprintf(os.Stdout, "Modes:\n\tmain\t\tMain node mode\n")
    fmt.Fprintf(os.Stdout, "\tnode\t\tNode publish mode\n")
    fmt.Fprintf(os.Stdout, "\tstandalone\tStandalone mode\n")
    fmt.Fprintf(os.Stdout, "Flags:\n")
    fmt.Fprintf(os.Stdout, "\t-s\t\tInitial URL to be crawled (seed)\n")
    fmt.Fprintf(os.Stdout, "\t-n\t\tNode IP/FQDN with port number\n")
    fmt.Fprintf(os.Stdout, "\t-h --help\tShow this message\n")
}

// Print main node subcommand usage
func PrintMainNodeUsage() {
    fmt.Fprintf(os.Stdout, "Main node mode\n")
    fmt.Fprintf(os.Stdout, "Usage: ./dist-crawler main [-h --help] -s seed -n nodeAddr\n")
    fmt.Fprintf(os.Stdout, "Flags:\n")
    fmt.Fprintf(os.Stdout, "\t-s\t\tInitial URL to be crawled (seed)\n")
    fmt.Fprintf(os.Stdout, "\t-n\t\tNode IP/FQDN with port number\n")
    fmt.Fprintf(os.Stdout, "\t-h --help\tShow this message\n")
}

// Print node subcommand usage
func PrintNodeUsage() {
    fmt.Fprintf(os.Stdout, "Node mode\n")
    fmt.Fprintf(os.Stdout, "Usage: ./dist-crawler node [--help]\n\n")
    fmt.Fprintf(os.Stdout, "Flags:\n")
    fmt.Fprintf(os.Stdout, "\t-h --help\tShow this message\n")
}

// Print standalone subcommand usage
func PrintStandaloneUsage() {
    fmt.Fprintf(os.Stdout, "Standalone mode\n")
    fmt.Fprintf(os.Stdout, "Usage: ./dist-crawler standalone [-h --help] -s seed")
    fmt.Fprintf(os.Stdout, "Flags:\n")
    fmt.Fprintf(os.Stdout, "\t-s\t\tInitial URL to be crawled (seed)\n")
    fmt.Fprintf(os.Stdout, "\t-h --help\tShow this message\n")
}
