package esp

import (
	"encoding/binary"
)

type RadioErp1Packet struct {
	*PacketBase
	Data          []byte
	SubTelNum     uint8
	DestinationID uint32
	RSSI          int
	IsEncrypted   bool
}

func (p *RadioErp1Packet) DecodeData(data []byte) error {
	p.Data = data
	return nil
}

func (p *RadioErp1Packet) DecodeOptionalData(data []byte) error {
	if len(data) != 7 {
		return InvalidLengthError
	}

	p.SubTelNum = data[0]
	p.DestinationID = binary.BigEndian.Uint32(data[1:5])
	p.RSSI = -1 * int(data[5])
	p.IsEncrypted = false
	if data[6] == 0x01 {
		p.IsEncrypted = true
	}

	return nil
}
