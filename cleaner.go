package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	units "github.com/docker/go-units"
)

const (
	noColor        = "\033[0m"
	lightBlueColor = "\033[1;34m"
	lightRedColor  = "\033[1;31m"
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
		return fmt.Errorf("Failed to connect to docker client! %v", err)
	}
	defer cli.Close()

	var SpaceReclaimed uint64
	SpaceReclaimed = 0

	if c.stoppedContainers {
		pruneFilter := filters.NewArgs()
		pruneFilter.Add("status", "exited")
		report, err := cli.ContainersPrune(context.Background(), pruneFilter)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		if len(report.ContainersDeleted) >= 1 {
			fmt.Printf("%sRemoving containers %s\n", lightBlueColor, noColor)
		}
		for _, containerID := range report.ContainersDeleted {
			fmt.Printf("%v\n", containerID[:10])
		}
		SpaceReclaimed += report.SpaceReclaimed
	}

	if c.unUsedImages {
		pruneFilter := filters.NewArgs()
		pruneFilter.Add("dangling", "false")
		report, err := cli.ImagesPrune(context.Background(), pruneFilter)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		if len(report.ImagesDeleted) >= 1 {
			fmt.Printf("%sRemoving un used images %s\n", lightBlueColor, noColor)
		}
		for _, image := range report.ImagesDeleted {
			fmt.Printf("%+v\n", image)
		}
		SpaceReclaimed += report.SpaceReclaimed
	}

	if c.unUsedVolumes {
		pruneFilter := filters.NewArgs()
		report, err := cli.VolumesPrune(context.Background(), pruneFilter)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		if len(report.VolumesDeleted) >= 1 {
			fmt.Printf("%sRemoving un used volumes %s\n", lightBlueColor, noColor)
		}

		for _, volume := range report.VolumesDeleted {
			fmt.Printf("%+v\n", volume[:10])
		}
		SpaceReclaimed += report.SpaceReclaimed
	}
	if SpaceReclaimed >= 1 {
		fmt.Printf("Total reclaimed space: %v\n", units.HumanSize(float64(SpaceReclaimed)))
	}
	return nil
}
