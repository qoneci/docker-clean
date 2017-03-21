package main

import (
	"fmt"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// Version of build
var Version = "not set"

func main() {
	version := kingpin.Flag("version", "Show Version").Bool()
	all := kingpin.Flag("all", "Prune all container/images/volumes not used").Bool()
	containers := kingpin.Flag("containers", "Prune all container not used").Short('c').Bool()
	images := kingpin.Flag("images", "Prune all images not used").Short('i').Bool()
	volumes := kingpin.Flag("volumes", "Prune all volumes not used").Short('v').Bool()
	kingpin.Parse()
	if *version {
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}
	if !*containers && !*images && !*volumes && !*all {
		fmt.Printf("No thing to remove. check the flags with --help\n")
		os.Exit(1)
	}

	var cleaner Cleaner

	if *all {
		cleaner.unUsedVolumes = true
		cleaner.stoppedContainers = true
		cleaner.unUsedImages = true
	} else {
		cleaner.unUsedVolumes = *volumes
		cleaner.stoppedContainers = *containers
		cleaner.unUsedImages = *images
	}
	err := cleaner.CleanUp()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
