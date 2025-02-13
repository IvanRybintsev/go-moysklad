package moysklad

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

// InventoryService
// Сервис для работы с инвентаризациями.
type InventoryService struct {
	Endpoint
	endpointGetList[Inventory]
	endpointCreate[Inventory]
	endpointCreateUpdateDeleteMany[Inventory]
	endpointDelete
	endpointGetById[Inventory]
	endpointUpdate[Inventory]
	endpointTemplate[Inventory]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[InventoryPosition]
	endpointAttributes
	endpointSyncId[Inventory]
	endpointRemove
}

func NewInventoryService(client *Client) *InventoryService {
	e := NewEndpoint(client, "entity/inventory")
	return &InventoryService{
		Endpoint:                       e,
		endpointGetList:                endpointGetList[Inventory]{e},
		endpointCreate:                 endpointCreate[Inventory]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Inventory]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[Inventory]{e},
		endpointUpdate:                 endpointUpdate[Inventory]{e},
		endpointTemplate:               endpointTemplate[Inventory]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[InventoryPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointSyncId:                 endpointSyncId[Inventory]{e},
		endpointRemove:                 endpointRemove{e},
	}
}

// Recalculate Запрос на пересчёт расчётных остатков у позиций инвентаризации.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-pereschet-raschetnogo-ostatka-w-inwentarizacii
func (s *InventoryService) Recalculate(ctx context.Context, id *uuid.UUID) (bool, *Response, error) {
	uri := fmt.Sprintf("rpc/inventory/%s/recalcCalculatedQuantity", id)
	e := NewEndpoint(s.client, uri)
	resp, err := NewRequestBuilder[any](e, ctx).do(http.MethodPut)
	if err != nil {
		return false, resp, err
	}
	ok := resp.StatusCode == http.StatusCreated
	return ok, resp, nil
}
