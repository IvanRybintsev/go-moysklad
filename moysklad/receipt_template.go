package moysklad

import (
	"github.com/google/uuid"
)

// ReceiptTemplate Шаблон печати кассовых чеков
// TODO: сервис не распознает тип сущности 'receipttemplate'
type ReceiptTemplate struct {
	AccountId *uuid.UUID `json:"accountId,omitempty"` // ID учетной записи
	Id        *uuid.UUID `json:"id,omitempty"`        // ID сущности
	Name      *string    `json:"name,omitempty"`      // Наименование
	Meta      *Meta      `json:"meta,omitempty"`      // Метаданные
	Header    *string    `json:"header,omitempty"`
	Footer    *string    `json:"footer,omitempty"`
}

func (r ReceiptTemplate) String() string {
	return Stringify(r)
}
