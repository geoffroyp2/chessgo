package move

// All values for the moveType

const (
	MOVE          uint32 = 0b0000
	PMOVE2        uint32 = 0b0001
	KCASTLE       uint32 = 0b0010
	QCASTLE       uint32 = 0b0011
	CAPTURE       uint32 = 0b0100
	EPCAPTURE     uint32 = 0b0101
	QPROM         uint32 = 0b1000
	RPROM         uint32 = 0b1001
	BPROM         uint32 = 0b1010
	NPROM         uint32 = 0b1011
	QPROMCAPTURE  uint32 = 0b1100
	RPROMCAPTURE  uint32 = 0b1101
	BPROMCAPTURE  uint32 = 0b1110
	NPROMCAPTURE  uint32 = 0b1111
	 
	CAPTUREMASK	  uint32 = 0b100
	PROMMASK      uint32 = 0b1000
	PROMPIECEMASK uint32 = 0b11
)