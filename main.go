package main

import (
	"log"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// Version of build
var Version = "not set"

func main() {
	version := kingpin.Flag("version", "Show Version").Short('v').Bool()
	all := kingpin.Flag("all", "Prune all container/images/volumes not used").Short('a').Bool()
	containers := kingpin.Flag("containers", "Prune all container not used").Short('c').Bool()
	images := kingpin.Flag("images", "Prune all images not used").Short('i').Bool()
	volumes := kingpin.Flag("volumes", "Prune all volumes not used").Short('v').Bool()
	kingpin.Parse()
	if *version {
		log.Printf("Version: %s\n", Version)
		os.Exit(0)
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
		log.Printf("%v", err)
	}
}
