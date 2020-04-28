package api

import (
	"github.com/propero-oss/csi-vcloud/pkg/common"
	"github.com/vmware/go-vcloud-director/v2/govcd"
)

type Manager struct {
	Client *govcd.VCDClient
	Config *common.Config
	Vdc    *govcd.Vdc
}

/*type Vdc struct {
	*govcd.Vdc
}*/

type VM struct {
	*govcd.VM
}

type BusType string
const (
	IDE BusType = "5"
	SCSI = "6"
	SATA = "20"
)