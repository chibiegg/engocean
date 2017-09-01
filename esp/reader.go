package esp

import (
	"encoding/binary"
	"io"
)

func ReadPacket(reader io.Reader) (Packet, error) {
	buffer := make([]byte, 0, 100)

	for {
		b := make([]byte, 1)
		readLen, err := reader.Read(b)
		if err != nil {
			return nil, err
		}
		if readLen == 0 {
			continue
		}

		buffer = append(buffer, b...)

		for len(buffer) > 0 {
			if buffer[0] != 0x55 {
				// skip buffer
				buffer = buffer[1:]
			}

			if len(buffer) < 6 {
				break
			}

			// check (Serial synchronization)
			dataLength := int(binary.BigEndian.Uint16(buffer[1:3]))
			optionalLength := int(buffer[3])
			// packetType := buffer[4]
			// crc8h := buffer[5]
			// TODO: check CHC8H

			totalLength := 6 + dataLength + optionalLength + 1

			if len(buffer) < totalLength {
				// Too short
				break
			}

			rawPacket := buffer[:totalLength]
			buffer = buffer[totalLength:]

			return DecodePacket(rawPacket)
		}
	}
}
