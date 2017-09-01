package esp

import (
	"github.com/chibiegg/engocean/erp"
)

type RadioErp2Packet struct {
	*PacketBase
	Telegram  erp.Telegram
	SubTelNum uint8
	RSSI      int
}

func (p *RadioErp2Packet) DecodeData(data []byte) error {
	telegram, err := erp.DecodeTelegram(data)
	if err != nil {
		return err
	}
	p.Telegram = telegram
	return nil
}

func (p *RadioErp2Packet) DecodeOptionalData(data []byte) error {
	if len(data) != 2 {
		return InvalidLengthError
	}

	p.SubTelNum = data[0]
	p.RSSI = -1 * int(data[1])

	return nil
}
