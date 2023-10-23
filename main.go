// main.go

package main

import (
	"flag"
	"log"
	"os"

	"tectonic_assets/commands"
)

func main() {
	browseCmd := flag.NewFlagSet("browse", flag.ExitOnError)
	optimizeCmd := flag.NewFlagSet("optimize", flag.ExitOnError)

	// Parse flags
	switch os.Args[1] {
	case "browse":
		err := browseCmd.Parse(os.Args[2:])
		if err != nil {
			return
		}
		commands.BrowseAssets()
	case "optimize":
		err := optimizeCmd.Parse(os.Args[2:])
		if err != nil {
			return
		}
		commands.OptimizeAssets()
	default:
		log.Fatal("Unknown command")
	}
}
