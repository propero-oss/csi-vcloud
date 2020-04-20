package api

import "github.com/vmware/go-vcloud-director/v2/govcd"

type Config struct {
	User     string
	Password string
	Org      string
	Href     string
	VDC      string
	VApp	 string
	Insecure bool
}

type Manager struct {
	Client *govcd.VCDClient
	Config *Config
	Vdc *Vdc
}

type Vdc struct {
	*govcd.Vdc
}

type VM struct {
	*govcd.VM
}

type BusType string
const (
	IDE BusType = "5"
	SCSI = "6"
	SATA = "20"
)

type ByteSize float64
const (
	_           = iota
	KiB ByteSize = 1 << (10 * iota)
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)