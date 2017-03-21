package main

import (
	"fmt"
	"os"
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// Version of build
var Version = "not set"

// Args config struct
type Args struct {
	all, version, containers, images, volumes, demon *bool
	demonInterval                                    *int
}

func runCleaner(a Args) {
	var cleaner Cleaner

	if *a.all {
		cleaner.unUsedVolumes = true
		cleaner.stoppedContainers = true
		cleaner.unUsedImages = true
	} else {
		cleaner.unUsedVolumes = *a.volumes
		cleaner.stoppedContainers = *a.containers
		cleaner.unUsedImages = *a.images
	}
	err := cleaner.CleanUp()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func setInterval(a Args) int {
	var interval int
	if *a.demonInterval == 0 {
		interval = 300 //SEC
	} else {
		interval = *a.demonInterval
	}
	return interval
}

func main() {
	var args Args
	args.version = kingpin.Flag("version", "Show Version").Bool()
	args.all = kingpin.Flag("all", "Prune all container/images/volumes not used").Bool()
	args.containers = kingpin.Flag("containers", "Prune all container not used").Short('c').Bool()
	args.images = kingpin.Flag("images", "Prune all images not used").Short('i').Bool()
	args.volumes = kingpin.Flag("volumes", "Prune all volumes not used").Short('v').Bool()
	args.demon = kingpin.Flag("demon", "run as demon in interval").Bool()
	args.demonInterval = kingpin.Flag("demon-interval", "demon in interval in SEC, default 300").Int()
	kingpin.Parse()
	if *args.version {
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}
	if !*args.containers && !*args.images && !*args.volumes && !*args.all {
		fmt.Printf("No thing to remove. check the flags with --help\n")
		os.Exit(1)
	}
	if *args.demon {
		interval := setInterval(args)
		for {
			var sleepTime = time.Duration(interval) * time.Second
			fmt.Printf("Will sleep for: %v nanoseconds", sleepTime)
			time.Sleep(sleepTime)
			runCleaner(args)
		}
	} else {
		runCleaner(args)
	}
}
