package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	units "github.com/docker/go-units"
)

// Cleaner struct
type Cleaner struct {
	stoppedContainers bool
	unUsedImages      bool
	unUsedVolumes     bool
}

// CleanUp state of cleaner settings
// can remove:
// stopped containers
// un used images and old versions of it.
func (c Cleaner) CleanUp() error {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Printf("Failed to connect to docker client! \n %v\n", err)
	}

	switch {
	case c.stoppedContainers:
		// status=exited
		pruneFilter := filters.NewArgs()
		pruneFilter.Add("status", "exited")
		report, err := cli.ContainersPrune(context.Background(), pruneFilter)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		log.Printf("Removed containers")
		for _, containerID := range report.ContainersDeleted {
			log.Printf("%v\n", containerID)
		}
		fmt.Printf("Total reclaimed space: %v", units.HumanSize(float64(report.SpaceReclaimed)))

	case c.unUsedImages:
		// dangling=true
		pruneFilter := filters.NewArgs()
		pruneFilter.Add("dangling", "true")
		report, err := cli.ImagesPrune(context.Background(), pruneFilter)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		for _, image := range report.ImagesDeleted {
			log.Printf("%+v", image)
		}
		fmt.Printf("Total reclaimed space: %v", units.HumanSize(float64(report.SpaceReclaimed)))

	case c.unUsedVolumes:
		pruneFilter := filters.NewArgs()
		report, err := cli.VolumesPrune(context.Background(), pruneFilter)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		for _, volume := range report.VolumesDeleted {
			log.Printf("%+v", volume)
		}

		fmt.Printf("Total reclaimed space: %v", units.HumanSize(float64(report.SpaceReclaimed)))
	}
	return nil
}
