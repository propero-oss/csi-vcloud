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
	ZBinBytes
	YBinBytes
)

const (
	DEFAULT_VOLUME_SIZE = 10
)