package messages

import "github.com/WenLopes/bank-transactions-api/domain"

const (
	GENERIC_ERROR                     = "ocorreu um erro ao processar a sua requisição"
	DOMAIN_ErrInvalidAmount           = domain.ErrInvalidAmount
	DOMAIN_BlockedAccount             = domain.BlockedAccount
	DOMAIN_InsufficientAccountBalance = domain.InsufficientAccountBalance
	INVALID_EVENT_TYPE                = "o tipo do evento não é válido"
	INVALID_AMOUNT_VALUE              = "o valor informado não é válido"
	DESTINATION_FIELD_REQUIRED        = "o campo destination é obrigatório para realizar essa operação"
	ORIGIN_FIELD_REQUIRED             = "o campo origin é obrigatório para realizar essa operação"
	REQUEST_INVALID                   = "erro na formatação do request"
)

var allMessages = map[int]string{
	1000: GENERIC_ERROR,
	1001: DOMAIN_ErrInvalidAmount,
	1002: DOMAIN_BlockedAccount,
	1003: DOMAIN_InsufficientAccountBalance,
	1004: INVALID_EVENT_TYPE,
	1005: INVALID_AMOUNT_VALUE,
	1006: DESTINATION_FIELD_REQUIRED,
	1007: ORIGIN_FIELD_REQUIRED,
	1008: REQUEST_INVALID,
}

type messageExport struct {
	Code     int
	ToString string
}

func Get(message string) messageExport {
	for index, value := range allMessages {
		if message == allMessages[index] {
			return messageExport{
				Code:     index,
				ToString: value,
			}
		}
	}
	return messageExport{}
}
