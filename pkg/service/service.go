package service

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"sync"
)

const (
	// Name is the name of the CSI plug-in.
	Name = "csi-vcloud.propero-oss.de"

	// VendorVersion is the version returned by GetPluginInfo.
	VendorVersion = "0.0.1"
)

var Manifest = map[string]string{
	"url": "https://github.com/rexray/gocsi/tree/master/mock",
}

// Service is the CSI Mock service provider.
type Service interface {
	csi.ControllerServer
	csi.IdentityServer
	csi.NodeServer
}

type service struct {
	sync.Mutex
	nodeID   string
	vols     []csi.Volume
	snaps    []csi.Snapshot
	volsRWL  sync.RWMutex
	snapsRWL sync.RWMutex
	volsNID  uint64
	snapsNID uint64
}

func New() Service {
	s := &service{nodeID: Name}
	return s
}