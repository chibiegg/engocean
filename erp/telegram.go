package erp

import (
	"errors"
	"fmt"
)

const (
	TelegramTypeRps = 0x00
	TelegramType1bs = 0x01
	TelegramType4bs = 0x02
)

const (
	TelegramAddressControl24BitWithoutDestinationId = 0x00
	TelegramAddressControl32BitWithoutDestinationId = 0x01
	TelegramAddressControl32Bit                     = 0x02
	TelegramAddressControl48BitWithoutDestinationId = 0x03
)

var (
	NotImplementedError = errors.New("Not implemented")
	InvalidLengthError  = errors.New("Invalid length")
)

type Telegram interface {
	GetType() byte
	DecodeData([]byte) error
}

type TelegramBase struct {
	Type           uint8
	AddressControl uint8
	OriginatorId   []byte
	DestinationId  []byte
}

func (t *TelegramBase) GetType() uint8 {
	return t.Type
}

func (t *TelegramBase) DecodeData([]byte) error {
	return nil
}

func DecodeTelegram(data []byte) (Telegram, error) {
	var err error
	fmt.Printf("%#v\n", data)

	var telegram Telegram

	if len(data) <= 6 {
		return nil, NotImplementedError
	}

	telegramBase := &TelegramBase{}
	telegramBase.AddressControl = (data[0] >> 5) & 0x07
	telegramBase.Type = data[0] & 0x0f

	switch telegramBase.AddressControl {
	case TelegramAddressControl24BitWithoutDestinationId:
		telegramBase.OriginatorId = data[1:4]
		data = data[4:]

	case TelegramAddressControl32BitWithoutDestinationId:
		telegramBase.OriginatorId = data[1:5]
		data = data[5:]

	case TelegramAddressControl32Bit:
		telegramBase.OriginatorId = data[1:5]
		telegramBase.DestinationId = data[6:10]
		data = data[10:]

	case TelegramAddressControl48BitWithoutDestinationId:
		telegramBase.OriginatorId = data[1:7]
		data = data[7:]
	}

	switch telegramBase.Type {
	case TelegramTypeRps:
		telegram = &RPSTelegram{TelegramBase: telegramBase}
	case TelegramType1bs:
		telegram = &OneBSTelegram{TelegramBase: telegramBase}
	case TelegramType4bs:
		telegram = &FourBSTelegram{TelegramBase: telegramBase}
	default:
		telegram = telegramBase
	}

	// TODO: Check CRC
	// crc := data[len(data)-1]
	data = data[:len(data)-1]

	err = telegram.DecodeData(data)
	if err != nil {
		return nil, err
	}

	return telegram, nil
}
