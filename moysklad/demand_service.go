package moysklad

// DemandService
// Сервис для работы с отгрузками.
type DemandService struct {
	endpointGetList[Demand]
	endpointCreate[Demand]
	endpointCreateUpdateDeleteMany[Demand]
	endpointDelete
	endpointGetById[Demand]
	endpointUpdate[Demand]
	endpointTemplate[Demand]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[DemandPosition]
	endpointAttributes
	endpointPublication
	endpointSyncId[Demand]
	endpointNamedFilter
	endpointRemove
}

func NewDemandService(client *Client) *DemandService {
	e := NewEndpoint(client, "entity/demand")
	return &DemandService{
		endpointGetList:                endpointGetList[Demand]{e},
		endpointCreate:                 endpointCreate[Demand]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Demand]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[Demand]{e},
		endpointUpdate:                 endpointUpdate[Demand]{e},
		endpointTemplate:               endpointTemplate[Demand]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[DemandPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncId:                 endpointSyncId[Demand]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointRemove:                 endpointRemove{e},
	}
}
