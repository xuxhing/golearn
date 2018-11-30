package byteoder

type Slice struct {
	Addr uintptr
	Len  uint32
	Cap  uint32
}

type Head struct {
	Magic   uint16
	Version uint8
	Reserve uint32
	Len     uint16
}
