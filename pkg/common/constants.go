package common

type ByteSize int64
const (
	_           = iota
	KBinBytes ByteSize = 1 << (10 * iota)
	MBinBytes
	GBinBytes
	TBinBytes
	PBinBytes
	EBinBytes
)

const (
	DEFAULT_VOLUME_SIZE = 10

	MAX_BUS_NUMBER = 3
	MAX_UNIT_NUMBER = 15
)