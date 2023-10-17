package moysklad

import (
	"encoding/json"
)

// PaymentInDocument Входящий платеж, Приходный ордер
type PaymentDocument struct {
	Meta *Meta `json:"meta"`
	raw  json.RawMessage
}

// UnmarshalJSON анмаршалит Входящий платеж, Приходный ордер, при expand=payments
func (p *PaymentDocument) UnmarshalJSON(data []byte) (err error) {
	type alias PaymentDocument
	var t alias

	if err = json.Unmarshal(data, &t); err != nil {
		return err
	}

	t.raw = data
	*p = PaymentDocument(t)

	return nil
}

// PaymentsIn Входящие платежи, Приходные ордеры
type Payments []PaymentDocument

//
//func (p PaymentInDocument) Data() json.RawMessage {
//	return p.w.raw
//}
//
//func (p PaymentInDocument) Meta() Meta {
//	return utils.Deref[Meta](p.w.Meta)
//}
//
//func (p PaymentInDocument) CashIn() (CashIn, bool) {
//	return ElementAsType[CashIn](p)
//}
//
//func (p PaymentInDocument) PaymentIn() (PaymentIn, bool) {
//	return ElementAsType[PaymentIn](p)
//}
//
////func (p PaymentsIn) GetCashesIn() Slice[CashIn] {
////	return getElementsByType[CashIn](p)
////}
////
////func (p PaymentsIn) GetPaymentsIn() Slice[PaymentIn] {
////	return getElementsByType[PaymentIn](p)
////}
//
//func (p PaymentOutDocument) Data() json.RawMessage {
//	return p.w.raw
//}
//
//func (p PaymentOutDocument) Meta() Meta {
//	return utils.Deref[Meta](p.w.Meta)
//}
//
//func (p PaymentOutDocument) CashOut() (CashOut, bool) {
//	return ElementAsType[CashOut](p)
//}
//
//func (p PaymentOutDocument) PaymentOut() (PaymentOut, bool) {
//	return ElementAsType[PaymentOut](p)
//}
//
////func (p PaymentsOut) GetCashesOut() Slice[CashOut] {
////	return getElementsByType[CashOut](p)
////}
////
////func (p PaymentsOut) GetPaymentsOut() Slice[PaymentOut] {
////	return getElementsByType[PaymentOut](p)
////}
