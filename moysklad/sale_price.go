package moysklad

// SalePrice Цена продажи
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-ceny-prodazhi
type SalePrice struct {
	Value     *float64   `json:"value,omitempty"`     // Значение цены
	Currency  *Currency  `json:"currency,omitempty"`  // Ссылка на валюту в формате Метаданных
	PriceType *PriceType `json:"priceType,omitempty"` // Тип цены
}

func (s SalePrice) String() string {
	return Stringify(s)
}

type SalePrices = Iterator[SalePrice]
