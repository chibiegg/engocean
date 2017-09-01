package erp

type OneBSTelegram struct {
	*TelegramBase
	Data byte
}

func (t *OneBSTelegram) DecodeData(data []byte) error {
	if len(data) != 1 {
		return InvalidLengthError
	}
	t.Data = data[0]
	return nil
}
