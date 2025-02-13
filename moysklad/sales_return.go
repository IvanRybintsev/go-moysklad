package moysklad

import (
	"github.com/google/uuid"
)

// SalesReturn Возврат покупателя.
// Ключевое слово: salesreturn
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-pokupatelq
type SalesReturn struct {
	AccountId           *uuid.UUID                      `json:"accountId,omitempty"`           // ID учетной записи
	Agent               *Counterparty                   `json:"agent,omitempty"`               // Метаданные контрагента
	AgentAccount        *AgentAccount                   `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	Applicable          *bool                           `json:"applicable,omitempty"`          // Отметка о проведении
	Attributes          *Attributes                     `json:"attributes,omitempty"`          // Коллекция метаданных доп. полей. Поля объекта
	Code                *string                         `json:"code,omitempty"`                // Код Возврата Покупателя
	Contract            *Contract                       `json:"contract,omitempty"`            // Метаданные договора
	Created             *Timestamp                      `json:"created,omitempty"`             // Дата создания
	Deleted             *Timestamp                      `json:"deleted,omitempty"`             // Момент последнего удаления Возврата Покупателя
	Description         *string                         `json:"description,omitempty"`         // Комментарий Возврата Покупателя
	ExternalCode        *string                         `json:"externalCode,omitempty"`        // Внешний код Возврата Покупателя
	Files               *Files                          `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group                          `json:"group,omitempty"`               // Отдел сотрудника
	Id                  *uuid.UUID                      `json:"id,omitempty"`                  // ID сущности
	Meta                *Meta                           `json:"meta,omitempty"`                // Метаданные
	Moment              *Timestamp                      `json:"moment,omitempty"`              // Дата документа
	Name                *string                         `json:"name,omitempty"`                // Наименование
	Organization        *Organization                   `json:"organization,omitempty"`        // Метаданные юрлица
	OrganizationAccount *AgentAccount                   `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	Owner               *Employee                       `json:"owner,omitempty"`               // Владелец (Сотрудник)
	Positions           *Positions[SalesReturnPosition] `json:"positions,omitempty"`           // Метаданные позиций Возврата Покупателя
	Printed             *bool                           `json:"printed,omitempty"`             // Напечатан ли документ
	Project             *Project                        `json:"project,omitempty"`             // Метаданные проекта
	Published           *bool                           `json:"published,omitempty"`           // Опубликован ли документ
	Rate                *Rate                           `json:"rate,omitempty"`                // Валюта
	SalesChannel        *SalesChannel                   `json:"salesChannel,omitempty"`        // Метаданные канала продаж
	Shared              *bool                           `json:"shared,omitempty"`              // Общий доступ
	State               *State                          `json:"state,omitempty"`               // Метаданные статуса Возврата Покупателя
	Store               *Store                          `json:"store,omitempty"`               // Метаданные склада
	Sum                 *float64                        `json:"sum,omitempty"`                 // Сумма
	SyncId              *uuid.UUID                      `json:"syncId,omitempty"`              // ID синхронизации. После заполнения недоступен для изменения
	Updated             *Timestamp                      `json:"updated,omitempty"`             // Момент последнего обновления
	VatEnabled          *bool                           `json:"vatEnabled,omitempty"`          // Учитывается ли НДС
	VatIncluded         *bool                           `json:"vatIncluded,omitempty"`         // Включен ли НДС в цену
	VatSum              *float64                        `json:"vatSum,omitempty"`              // Сумма включая НДС
	Demand              *Demand                         `json:"demand,omitempty"`              // Ссылка на отгрузку, по которой произошел возврат в формате Метаданных
	Losses              *Losses                         `json:"losses,omitempty"`              // Массив ссылок на связанные списания в формате Метаданных
	Payments            *Payments                       `json:"payments,omitempty"`            // Массив ссылок на связанные операции в формате Метаданных
	PayedSum            *float64                        `json:"payedSum,omitempty"`            // Сумма исходящих платежей по возврату покупателя
	FactureOut          *FactureOut                     `json:"factureOut,omitempty"`          // Ссылка на Счет-фактуру выданный, с которым связан этот возврат, в формате Метаданных
}

func (s SalesReturn) String() string {
	return Stringify(s)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (s SalesReturn) GetMeta() *Meta {
	return s.Meta
}

func (s SalesReturn) MetaType() MetaType {
	return MetaTypeSalesReturn
}

type SalesReturns = Iterator[SalesReturn]

// SalesReturnPosition Позиция Возврата покупателя.
// Ключевое слово: salesreturnposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-pokupatelq-vozwraty-pokupatelej-pozicii-vozwrata-pokupatelq
type SalesReturnPosition struct {
	AccountId  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Cost       *float64            `json:"cost,omitempty"`       // Себестоимость (выводится, если документ был создан без основания)
	Country    *Country            `json:"country,omitempty"`    // Метаданные Страны
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	GTD        *GTD                `json:"gtd,omitempty"`        // ГТД
	Id         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Slot       *Slot               `json:"slot,omitempty"`       // Метаданные
	Things     *Things             `json:"things,omitempty"`     // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Stock      *Stock              `json:"stock,omitempty"`      // Остатки и себестоимость `?fields=stock&expand=positions`
}

func (s SalesReturnPosition) String() string {
	return Stringify(s)
}

func (s SalesReturnPosition) MetaType() MetaType {
	return MetaTypeSalesReturnPosition
}
