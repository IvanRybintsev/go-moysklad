package moysklad

import (
	"github.com/google/uuid"
)

// InvoiceOut Счет покупателю.
// Ключевое слово: invoiceout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-pokupatelu
type InvoiceOut struct {
	AccountId            *uuid.UUID                  `json:"accountId,omitempty"`            // ID учетной записи
	Agent                *Counterparty               `json:"agent,omitempty"`                // Метаданные контрагента
	AgentAccount         *AgentAccount               `json:"agentAccount,omitempty"`         // Метаданные счета контрагента
	Applicable           *bool                       `json:"applicable,omitempty"`           // Отметка о проведении
	Attributes           *Attributes                 `json:"attributes,omitempty"`           // Коллекция метаданных доп. полей
	Code                 *string                     `json:"code,omitempty"`                 // Код
	Contract             *Contract                   `json:"contract,omitempty"`             // Метаданные договора
	Created              *Timestamp                  `json:"created,omitempty"`              // Дата создания
	Deleted              *Timestamp                  `json:"deleted,omitempty"`              // Момент последнего удаления
	Description          *string                     `json:"description,omitempty"`          // Комментарий
	ExternalCode         *string                     `json:"externalCode,omitempty"`         // Внешний код
	Files                *Files                      `json:"files,omitempty"`                // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                *Group                      `json:"group,omitempty"`                // Отдел сотрудника
	Id                   *uuid.UUID                  `json:"id,omitempty"`                   // ID сущности
	Meta                 *Meta                       `json:"meta,omitempty"`                 // Метаданные
	Moment               *Timestamp                  `json:"moment,omitempty"`               // Дата документа
	Name                 *string                     `json:"name,omitempty"`                 // Наименование
	Organization         *Organization               `json:"organization,omitempty"`         // Метаданные юрлица
	OrganizationAccount  *AgentAccount               `json:"organizationAccount,omitempty"`  // Метаданные счета юрлица
	Owner                *Employee                   `json:"owner,omitempty"`                // Владелец (Сотрудник)
	PayedSum             *float64                    `json:"payedSum,omitempty"`             // Сумма входящих платежей по Счету
	PaymentPlannedMoment *Timestamp                  `json:"paymentPlannedMoment,omitempty"` // Планируемая дата оплаты
	Positions            *Positions[InvoicePosition] `json:"positions,omitempty"`            // Метаданные позиций
	Printed              *bool                       `json:"printed,omitempty"`              // Напечатан ли документ
	Project              *Project                    `json:"project,omitempty"`              // Проект
	Published            *bool                       `json:"published,omitempty"`            // Опубликован ли документ
	Rate                 *Rate                       `json:"rate,omitempty"`                 // Валюта
	Shared               *bool                       `json:"shared,omitempty"`               // Общий доступ
	ShippedSum           *float64                    `json:"shippedSum,omitempty"`           // Сумма отгруженного
	State                *State                      `json:"state,omitempty"`                // Метаданные статуса
	Store                *Store                      `json:"store,omitempty"`                // Метаданные склада
	Sum                  *float64                    `json:"sum,omitempty"`                  // Сумма
	SyncId               *uuid.UUID                  `json:"syncId,omitempty"`               // ID синхронизации. После заполнения недоступен для изменения
	Updated              *Timestamp                  `json:"updated,omitempty"`              // Момент последнего обновления
	VatEnabled           *bool                       `json:"vatEnabled,omitempty"`           // Учитывается ли НДС
	VatIncluded          *bool                       `json:"vatIncluded,omitempty"`          // Включен ли НДС в цену
	VatSum               *float64                    `json:"vatSum,omitempty"`               // Сумма включая НДС
	CustomerOrder        *CustomerOrder              `json:"customerOrder,omitempty"`        // Ссылка на Заказ Покупателя, с которым связан этот Счет покупателю в формате Метаданных
	Payments             *Payments                   `json:"payments,omitempty"`             // Массив ссылок на связанные операции в формате Метаданных
	Demands              Iterator[Demand]            `json:"demands,omitempty"`              // Массив ссылок на связанные отгрузки в формате Метаданных
}

func (i InvoiceOut) String() string {
	return Stringify(i)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (i InvoiceOut) GetMeta() *Meta {
	return i.Meta
}

func (i InvoiceOut) MetaType() MetaType {
	return MetaTypeInvoiceOut
}

// ConvertToOperation удовлетворяет интерфейсу OperationInType
func (i InvoiceOut) ConvertToOperation(linkedSum *float64) (*OperationIn, error) {
	return &OperationIn{}, nil //OperationFromEntity(c, linkedSum)
}

type InvoicesOut = Iterator[InvoiceOut]

// InvoiceOutPosition Позиция Счета покупателю.
// Ключевое слово: invoiceposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-pokupatelu-scheta-pokupatelqm-pozicii-scheta-pokupatelu
type InvoiceOutPosition struct {
	InvoicePosition
}

func (i InvoiceOutPosition) String() string {
	return Stringify(i)
}

func (i InvoiceOutPosition) MetaType() MetaType {
	return MetaTypeInvoicePosition
}
