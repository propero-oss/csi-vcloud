package common

import "github.com/container-storage-interface/spec/lib/go/csi"

var (
	// VolumeCaps represents how the volume could be accessed.
	// It is SINGLE_NODE_WRITER since vSphere CNS Block volume could only be
	// attached to a single node at any given time.
	VolumeCaps = []csi.VolumeCapability_AccessMode{
		{
			Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER,
		},
	}
)
