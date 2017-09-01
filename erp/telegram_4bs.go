package erp

type FourBSTelegram struct {
	*TelegramBase
	Data []byte
}

func (t *FourBSTelegram) DecodeData(data []byte) error {
	if len(data) != 4 {
		return InvalidLengthError
	}
	t.Data = data[0:4]
	return nil
}
