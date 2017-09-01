package esp

import (
	"encoding/binary"
	"errors"
	"fmt"
)

const (
	PacketTypeRadioErp1        = 0x01
	PacketTypeResponse         = 0x02
	PacketTypeRadioSubTel      = 0x03
	PacketTypeEvent            = 0x04
	PacketTypeCommonCommand    = 0x05
	PacketTypeSmartAckCommand  = 0x06
	PacketTypeRemoteManCommand = 0x07
	PacketTypeRadioMessage     = 0x09
	PacketTypeRadioErp2        = 0x0a
)

var (
	InvalidLengthError = errors.New("Invalid length")
)

type Packet interface {
	GetType() byte
	DecodeData([]byte) error
	DecodeOptionalData([]byte) error
}

type PacketBase struct {
	Type byte
}

func (p *PacketBase) GetType() byte {
	return p.Type
}

func (p *PacketBase) DecodeData(data []byte) error {
	return nil
}

func (p *PacketBase) DecodeOptionalData(data []byte) error {
	return nil
}

func DecodePacket(data []byte) (Packet, error) {
	var err error

	if len(data) < 6 {
		return nil, fmt.Errorf("Too short")
	}

	if data[0] != 0x55 {
		return nil, fmt.Errorf("Invalid sync byte")
	}

	// check (Serial synchronization)
	dataLength := int(binary.BigEndian.Uint16(data[1:3]))
	optionalLength := int(data[3])
	packetType := data[4]
	// crc8h := data[5]
	// TODO: check CHC8H

	totalLength := 6 + dataLength + optionalLength + 1

	if len(data) != totalLength {
		// Too short
		return nil, fmt.Errorf("Too short")
	}

	var packet Packet

	packetBase := &PacketBase{Type: packetType}

	switch packetType {
	case PacketTypeRadioErp1:
		x := &RadioErp1Packet{}
		x.PacketBase = packetBase
		packet = x
	case PacketTypeRadioErp2:
		x := &RadioErp2Packet{}
		x.PacketBase = packetBase
		packet = x
	default:
		packet = packetBase
	}

	err = packet.DecodeData(data[6 : 6+dataLength])
	if err != nil {
		return nil, err
	}
	err = packet.DecodeOptionalData(data[6+dataLength : 6+dataLength+optionalLength])
	if err != nil {
		return nil, err
	}

	return packet, nil
}
