package moysklad

// EnterService
// Сервис для работы с оприходованиями.
type EnterService struct {
	endpointGetList[Enter]
	endpointCreate[Enter]
	endpointCreateUpdateDeleteMany[Enter]
	endpointDelete
	endpointGetById[Enter]
	endpointUpdate[Enter]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[EnterPosition]
	endpointAttributes
	endpointPublication
	endpointSyncId[Enter]
	endpointNamedFilter
	endpointRemove
}

func NewEnterService(client *Client) *EnterService {
	e := NewEndpoint(client, "entity/enter")
	return &EnterService{
		endpointGetList:                endpointGetList[Enter]{e},
		endpointCreate:                 endpointCreate[Enter]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Enter]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[Enter]{e},
		endpointUpdate:                 endpointUpdate[Enter]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[EnterPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncId:                 endpointSyncId[Enter]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointRemove:                 endpointRemove{e},
	}
}
