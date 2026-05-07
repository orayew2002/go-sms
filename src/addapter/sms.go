package addapter

import (
	"fmt"

	"github.com/fiorix/go-smpp/smpp"
	"github.com/richxcame/gosms/src/utils"
)

type SmsService struct {
	tx *smpp.Transmitter
}

func DefaultSmsService() (*SmsService, error) {
	tx := &smpp.Transmitter{
		Addr:   utils.GetEnvD("SMS_IP", "localhost") + ":" + utils.GetEnvD("SMS_PORT", "5555"),
		User:   utils.GetEnvD("SMS_LOGIN", "12345"),
		Passwd: utils.GetEnvD("SMS_PASSWORD", "strong_passowrd"),
	}

	conn := tx.Bind()
	var status smpp.ConnStatus
	if status = <-conn; status.Error() != nil {
		return nil, fmt.Errorf("Unable to connect, aborting: %w", status.Error())
	}

	return &SmsService{tx}, nil
}

func (s *SmsService) Send(message *smpp.ShortMessage) (string, error) {
	sm, err := s.tx.Submit(message)
	if err != nil {
		return "", err
	}

	return sm.RespID(), nil
}
