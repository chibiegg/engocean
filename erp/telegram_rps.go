package erp

type RPSTelegram struct {
	*TelegramBase
	Data byte
}

func (t *RPSTelegram) DecodeData(data []byte) error {
	if len(data) != 1 {
		return InvalidLengthError
	}
	t.Data = data[0]
	return nil
}
