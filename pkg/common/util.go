package common

import (
	"fmt"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"os"
)


func GetHostname() (string, error){
	name, err := os.Hostname()
	if err != nil {
		return "", fmt.Errorf("could not retrieve Hostname: %s", err)
	}

	return name, nil
}

// IsValidVolumeCapabilities is the helper function to validate capabilities of volume.
func IsValidVolumeCapabilities(volCaps []*csi.VolumeCapability) bool {
	hasSupport := func(cap *csi.VolumeCapability) bool {
		for _, c := range VolumeCaps {
			if c.GetMode() == cap.AccessMode.GetMode() {
				return true
			}
		}
		return false
	}
	foundAll := true
	for _, c := range volCaps {
		if !hasSupport(c) {
			foundAll = false
		}
	}
	return foundAll
}