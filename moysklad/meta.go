package moysklad

import (
	"fmt"
	"reflect"
)

type HasMeta interface {
	MetaTyper
	GetMeta() *Meta
}

type MetaTyper interface {
	MetaType() MetaType
}

// Meta Метаданные объекта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-metadannye-metadannye-ob-ekta
type Meta struct {
	Href         *string  `json:"href,omitempty"`         // Ссылка на объект
	MetadataHref *string  `json:"metadataHref,omitempty"` // Ссылка на метаданные сущности (Другой вид метаданных. Присутствует не во всех сущностях)
	Type         MetaType `json:"type,omitempty"`         // Тип объекта
	MediaType    *string  `json:"mediaType,omitempty"`    // Тип данных, который приходят в ответ от сервиса, либо отправляется в теле запроса
	UUIDHref     *string  `json:"uuidHref,omitempty"`     // Ссылка на объект на UI. Присутствует не во всех сущностях. Может быть использована для получения uuid
	DownloadHref *string  `json:"downloadHref,omitempty"` // Ссылка на скачивание Изображения и миниатюр изображений. Данный параметр указывается только в meta для Изображения у Товара или Комплекта, а также в поле miniature везде, где используются изображения. Если миниатюра не была создана, то значение данного поля равно null. Для создания миниатюры нужно перейти по ссылке, указанной в href в миниатюре.
}

func (m Meta) String() string {
	return Stringify(m)
}

type MetaWrapper struct {
	Meta Meta `json:"meta"`
}

type MetaName struct {
	Meta Meta   `json:"meta"`
	Name string `json:"name"`
}

func (m MetaWrapper) String() string {
	return Stringify(m)
}

// MetaCollection Метаданные коллекции.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-metadannye-metadannye-kollekcii
type MetaCollection struct {
	Href         string `json:"href,omitempty"`         // Ссылка на объект
	Type         string `json:"type,omitempty"`         // Тип объекта
	MediaType    string `json:"mediaType,omitempty"`    // Тип данных, который приходят в ответ от сервиса, либо отправляется в теле запроса
	Size         int    `json:"size,omitempty"`         // Количество элементов в коллекции
	Limit        int    `json:"limit,omitempty"`        // Максимальное число элементов в коллекции, возвращаемых за один запрос
	Offset       int    `json:"offset,omitempty"`       // Смещение выборки коллекции от первого элемента
	NextHref     string `json:"nextHref,omitempty"`     // Ссылка на следующую страницу коллекции
	PreviousHref string `json:"previousHref,omitempty"` // Ссылка на предыдущую страницу коллекции
}

func (m MetaCollection) String() string {
	return Stringify(m)
}

type Context struct {
	Employee MetaWrapper `json:"employee,omitempty"`
}

func (m Context) String() string {
	return Stringify(m)
}

type AssortmentResult struct {
	Context Context        `json:"context,omitempty"`
	Meta    MetaCollection `json:"meta,omitempty"`
	Rows    Assortment     `json:"rows,omitempty"`
}

func (m AssortmentResult) String() string {
	return Stringify(m)
}

// MetaArray Объект с полями meta и rows, где rows - массив объектов
type MetaArray[T any] struct {
	Meta MetaCollection `json:"meta,omitempty"`
	Rows Slice[T]       `json:"rows,omitempty"`
}

func (m MetaArray[T]) String() string {
	return Stringify(m)
}

type MetaType string

func (m MetaType) String() string {
	return string(m)
}

const (
	MetaTypeAccount                          MetaType = "account"
	MetaTypeAccumulationDiscount             MetaType = "accumulationdiscount"
	MetaTypeApplication                      MetaType = "application"
	MetaTypeAssortment                       MetaType = "assortment"
	MetaTypeAssortmentSettings               MetaType = "assortmentsettings"
	MetaTypeAsync                            MetaType = "async"
	MetaTypeAttribute                        MetaType = "attributemetadata"
	MetaTypeAudit                            MetaType = "audit"
	MetaTypeAuditEvent                       MetaType = "auditevent"
	MetaTypeBonusProgram                     MetaType = "bonusprogram"
	MetaTypeBonusTransaction                 MetaType = "bonustransaction"
	MetaTypeBundle                           MetaType = "bundle"
	MetaTypeBundleComponent                  MetaType = "bundlecomponent"
	MetaTypeCashier                          MetaType = "cashier"
	MetaTypeCashIn                           MetaType = "cashin"
	MetaTypeCashOut                          MetaType = "cashout"
	MetaTypeCharacteristic                   MetaType = "attributemetadata"
	MetaTypeCommissionReportIn               MetaType = "commissionreportin"
	MetaTypeCommissionReportInPosition       MetaType = "commissionreportinposition"
	MetaTypeCommissionReportInReturnPosition MetaType = "commissionreportinreturnedposition"
	MetaTypeCommissionReportOut              MetaType = "commissionreportout"
	MetaTypeCommissionReportOutPosition      MetaType = "commissionreportoutposition"
	MetaTypeCompanySettings                  MetaType = "companysettings"
	MetaTypeConsignment                      MetaType = "consignment"
	MetaTypeContactPerson                    MetaType = "contactperson"
	MetaTypeContract                         MetaType = "contract"
	MetaTypeCounterparty                     MetaType = "counterparty"
	MetaTypeCounterPartyAdjustment           MetaType = "counterpartyadjustment"
	MetaTypeCountry                          MetaType = "country"
	MetaTypeCurrency                         MetaType = "currency"
	MetaTypeCustomerOrder                    MetaType = "customerorder"
	MetaTypeCustomerOrderPosition            MetaType = "customerorderposition"
	MetaTypeCustomEntity                     MetaType = "customentity"
	MetaTypeCustomTemplate                   MetaType = "customtemplate"
	MetaTypeDemand                           MetaType = "demand"
	MetaTypeDemandPosition                   MetaType = "demandposition"
	MetaTypeDiscount                         MetaType = "discount"
	MetaTypeEmbeddedTemplate                 MetaType = "embeddedtemplate"
	MetaTypeEmployee                         MetaType = "employee"
	MetaTypeEmployeeContext                  MetaType = "employee"
	MetaTypeEnter                            MetaType = "enter"
	MetaTypeEnterPosition                    MetaType = "enterposition"
	MetaTypeExpenseItem                      MetaType = "expenseitem"
	MetaTypeFactureIn                        MetaType = "facturein"
	MetaTypeFactureOut                       MetaType = "factureout"
	MetaTypeFiles                            MetaType = "files"
	MetaTypeGroup                            MetaType = "group"
	MetaTypeImage                            MetaType = "image"
	MetaTypeInternalOrder                    MetaType = "internalorder"
	MetaTypeInternalOrderPosition            MetaType = "internalorderposition"
	MetaTypeInventory                        MetaType = "inventory"
	MetaTypeInventoryPosition                MetaType = "inventoryposition"
	MetaTypeInvoiceIn                        MetaType = "invoicein"
	MetaTypeInvoiceOut                       MetaType = "invoiceout"
	MetaTypeInvoicePosition                  MetaType = "invoiceposition"
	MetaTypeLoss                             MetaType = "loss"
	MetaTypeLossPosition                     MetaType = "lossposition"
	MetaTypeNamedFilter                      MetaType = "namedfilter"
	MetaTypeMove                             MetaType = "move"
	MetaTypeMovePosition                     MetaType = "moveposition"
	MetaTypeNote                             MetaType = "note"
	MetaTypeNotification                     MetaType = "notification"
	MetaTypePublication                      MetaType = "operationpublication"
	MetaTypeOrganization                     MetaType = "organization"
	MetaTypePaymentIn                        MetaType = "paymentin"
	MetaTypePaymentOut                       MetaType = "paymentout"
	MetaTypePersonalDiscount                 MetaType = "personaldiscount"
	MetaTypePrepayment                       MetaType = "prepayment"
	MetaTypePrepaymentPosition               MetaType = "prepaymentposition"
	MetaTypePrepaymentReturn                 MetaType = "prepaymentreturn"
	MetaTypePrepaymentReturnPosition         MetaType = "prepaymentreturnposition"
	MetaTypePriceList                        MetaType = "pricelist"
	MetaTypePriceListPosition                MetaType = "pricelistrow"
	MetaTypePriceType                        MetaType = "pricetype"
	MetaTypeProcessing                       MetaType = "processing"
	MetaTypeProcessingOrder                  MetaType = "processingorder"
	MetaTypeProcessingOrderPosition          MetaType = "processingorderposition"
	MetaTypeProcessingPlan                   MetaType = "processingplan"
	MetaTypeProcessingPlanMaterial           MetaType = "processingplanmaterial"
	MetaTypeProcessingPlanProduct            MetaType = "processingplanresult"
	MetaTypeProcessingPositionMaterial       MetaType = "processingpositionmaterial"
	MetaTypeProcessingPositionProduct        MetaType = "processingpositionresult"
	MetaTypeProcessingProcess                MetaType = "processingprocess"
	MetaTypeProcessingProcessPosition        MetaType = "processingprocessposition"
	MetaTypeProcessingStage                  MetaType = "processingstage"
	MetaTypeProduct                          MetaType = "product"
	MetaTypeProductFolder                    MetaType = "productfolder"
	MetaTypeProject                          MetaType = "project"
	MetaTypePurchaseOrder                    MetaType = "purchaseorder"
	MetaTypePurchaseOrderPosition            MetaType = "purchaseorderposition"
	MetaTypePurchaseReturn                   MetaType = "purchasereturn"
	MetaTypePurchaseReturnPosition           MetaType = "purchasereturnposition"
	MetaTypeReceiptTemplate                  MetaType = "receipttemplate"
	MetaTypeRegion                           MetaType = "region"
	MetaTypeRetailDemand                     MetaType = "demand"
	MetaTypeRetailDemandPosition             MetaType = "demandposition"
	MetaTypeRetailDrawerCashIn               MetaType = "retaildrawercashin"
	MetaTypeRetailDrawerCashOut              MetaType = "retaildrawercashout"
	MetaTypeRetailSalesReturn                MetaType = "retailsalesreturn"
	MetaTypeRetailSalesReturnPosition        MetaType = "salesreturnposition"
	MetaTypeRetailShift                      MetaType = "retailshift"
	MetaTypeRetailStore                      MetaType = "retailstore"
	MetaTypeRoundOffDiscount                 MetaType = "discount"
	MetaTypeSalesReturn                      MetaType = "salesreturn"
	MetaTypeSalesReturnPosition              MetaType = "salesreturnposition"
	MetaTypeService                          MetaType = "service"
	MetaTypeSlot                             MetaType = "slot"
	MetaTypeSpecialPriceDiscount             MetaType = "specialpricediscount"
	MetaTypeState                            MetaType = "state"
	MetaTypeStore                            MetaType = "store"
	MetaTypeStoreZone                        MetaType = "storezone"
	MetaTypeSupply                           MetaType = "supply"
	MetaTypeSupplyPosition                   MetaType = "supplyposition"
	MetaTypeTask                             MetaType = "task"
	MetaTypeTaskNote                         MetaType = "tasknote"
	MetaTypeUom                              MetaType = "uom"
	MetaTypeVariant                          MetaType = "variant"
	MetaTypeWebhook                          MetaType = "webhook"
	MetaTypeCounterpartySettings             MetaType = "counterpartysettings"
	MetaTypeRole                             MetaType = "role"
	MetaTypeSystemRole                       MetaType = "systemrole"
	MetaTypeIndividualRole                   MetaType = "individualrole"
	MetaTypeCustomRole                       MetaType = "customrole"
	MetaTypeUserSettings                     MetaType = "usersettings"
	MetaTypeSubscription                     MetaType = "subscription"
	MetaTypeSalesChannel                     MetaType = "saleschannel"
	MetaTypeMetadata                         MetaType = "metadata"
	MetaTypeTaxRate                          MetaType = "taxrate"
	MetaTypeToken                            MetaType = "token"
	MetaTypeReportStock                      MetaType = "stock"
	MetaTypeReportStockByOperation           MetaType = "stockbyoperation"
	MetaTypeReportStockByStore               MetaType = "stockbystore"
	MetaTypeReportMoney                      MetaType = "money"
	MetaTypeReportMoneyPlotSeries            MetaType = "moneyplotseries"
	MetaTypeReportProfitByCounterparty       MetaType = "salesbyCounterparty"
	MetaTypeReportProfitByEmployee           MetaType = "salesbyemployee"
	MetaTypeReportProfitByProduct            MetaType = "salesbyproduct"
	MetaTypeReportProfitBySalesChannel       MetaType = "salesbysaleschannel"
	MetaTypeReportProfitByVariant            MetaType = "salesbyvariant"
	MetaTypeReportOrders                     MetaType = "ordersplotseries"
	MetaTypeReportSales                      MetaType = "salesplotseries"
	MetaTypeReportTurnover                   MetaType = "turnover"
	MetaTypeReportDashboard                  MetaType = "dashboard"
	MetaTypeReportCounterparty               MetaType = "counterparty"
	MetaTypeWebhookStock                     MetaType = "webhookstock"
	MetaTypeProcessingPlanFolder             MetaType = "processingplanfolder"
)

func MetaTypeFromEntity(v any) (MetaType, error) {
	var metaType MetaType
	var err error

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		v = val.Elem().Interface()
	}

	// TODO
	//switch any(*new(T)).(type) {
	switch v.(type) {
	case ContextEmployee:
		metaType = MetaTypeEmployeeContext
	case AccumulationDiscount:
		metaType = MetaTypeAccumulationDiscount
	case Assortment:
		metaType = MetaTypeAssortment
	case Application:
		metaType = MetaTypeApplication
	case BonusProgram:
		metaType = MetaTypeBonusProgram
	case BonusTransaction:
		metaType = MetaTypeBonusTransaction
	case Bundle:
		metaType = MetaTypeBundle
	case BundleComponent:
		metaType = MetaTypeBundleComponent
	case Cashier:
		metaType = MetaTypeCashier
	case CashIn:
		metaType = MetaTypeCashIn
	case CashOut:
		metaType = MetaTypeCashOut
	case CommissionReportIn:
		metaType = MetaTypeCommissionReportIn
	case CommissionReportInPosition:
		metaType = MetaTypeCommissionReportInPosition
	case CommissionReportOut:
		metaType = MetaTypeCommissionReportOut
	case CommissionReportOutPosition:
		metaType = MetaTypeCommissionReportOutPosition
	case CompanySettings:
		metaType = MetaTypeCompanySettings
	case Consignment:
		metaType = MetaTypeConsignment
	case ContactPerson:
		metaType = MetaTypeContactPerson
	case Contract:
		metaType = MetaTypeContract
	case Counterparty:
		metaType = MetaTypeCounterparty
	case Country:
		metaType = MetaTypeCountry
	case Currency:
		metaType = MetaTypeCurrency
	case CustomerOrder:
		metaType = MetaTypeCustomerOrder
	case CustomerOrderPosition:
		metaType = MetaTypeCustomerOrderPosition
	case CustomEntity:
	case CustomEntityElement:
		metaType = MetaTypeCustomEntity
	case Demand:
		metaType = MetaTypeDemand
	case DemandPosition:
		metaType = MetaTypeDemandPosition
	case Discount:
		metaType = MetaTypeDiscount
	case Employee:
		metaType = MetaTypeEmployee
	case Enter:
		metaType = MetaTypeEnter
	case EnterPosition:
		metaType = MetaTypeEnterPosition
	case ExpenseItem:
		metaType = MetaTypeExpenseItem
	case FactureIn:
		metaType = MetaTypeFactureIn
	case FactureOut:
		metaType = MetaTypeFactureOut
	case Group:
		metaType = MetaTypeGroup
	case Image:
		metaType = MetaTypeImage
	case InternalOrder:
		metaType = MetaTypeInternalOrder
	case InternalOrderPosition:
		metaType = MetaTypeInternalOrderPosition
	case Inventory:
		metaType = MetaTypeInventory
	case InventoryPosition:
		metaType = MetaTypeInventoryPosition
	case InvoiceIn:
		metaType = MetaTypeInvoiceIn
	case InvoiceOut:
		metaType = MetaTypeInvoiceOut
	case InvoicePosition:
		metaType = MetaTypeInvoicePosition
	case Loss:
		metaType = MetaTypeLoss
	case LossPosition:
		metaType = MetaTypeLossPosition
	case Move:
		metaType = MetaTypeMove
	case MovePosition:
		metaType = MetaTypeMovePosition
	case Note:
		metaType = MetaTypeNote
	case Organization:
		metaType = MetaTypeOrganization
	case PaymentIn:
		metaType = MetaTypePaymentIn
	case PaymentOut:
		metaType = MetaTypePaymentOut
	case PersonalDiscount:
		metaType = MetaTypePersonalDiscount
	case Prepayment:
		metaType = MetaTypePrepayment
	case PrepaymentPosition:
		metaType = MetaTypePrepaymentPosition
	case PrepaymentReturn:
		metaType = MetaTypePrepaymentReturn
	case PrepaymentReturnPosition:
		metaType = MetaTypePrepaymentReturnPosition
	case PriceList:
		metaType = MetaTypePriceList
	case PriceListPosition:
		metaType = MetaTypePriceListPosition
	case Processing:
		metaType = MetaTypeProcessing
	case ProcessingOrder:
		metaType = MetaTypeProcessingOrder
	case ProcessingOrderPosition:
		metaType = MetaTypeProcessingOrderPosition
	case ProcessingPlan:
		metaType = MetaTypeProcessingPlan
	case ProcessingPlanMaterial:
		metaType = MetaTypeProcessingPlanMaterial
	case ProcessingPlanProduct:
		metaType = MetaTypeProcessingPlanProduct
	case ProcessingPositionMaterial:
		metaType = MetaTypeProcessingPositionMaterial
	case ProcessingPositionProduct:
		metaType = MetaTypeProcessingPositionProduct
	case ProcessingProcess:
		metaType = MetaTypeProcessingProcess
	case ProcessingStage:
		metaType = MetaTypeProcessingStage
	case Product:
		metaType = MetaTypeProduct
	case ProductFolder:
		metaType = MetaTypeProductFolder
	case Project:
		metaType = MetaTypeProject
	case PurchaseOrder:
		metaType = MetaTypePurchaseOrder
	case PurchaseOrderPosition:
		metaType = MetaTypePurchaseOrderPosition
	case PurchaseReturn:
		metaType = MetaTypePurchaseReturn
	case PurchaseReturnPosition:
		metaType = MetaTypePurchaseReturnPosition
	case ReceiptTemplate:
		metaType = MetaTypeReceiptTemplate
	case Region:
		metaType = MetaTypeRegion
	case RetailDemand:
		metaType = MetaTypeRetailDemand
	case RetailDrawerCashIn:
		metaType = MetaTypeRetailDrawerCashIn
	case RetailDrawerCashOut:
		metaType = MetaTypeRetailDrawerCashOut
	case RetailSalesReturn:
		metaType = MetaTypeRetailSalesReturn
	case RetailShift:
		metaType = MetaTypeRetailShift
	case RetailStore:
		metaType = MetaTypeRetailStore
	case SalesReturn:
		metaType = MetaTypeSalesReturn
	case SalesReturnPosition:
		metaType = MetaTypeSalesReturnPosition
	case Service:
		metaType = MetaTypeService
	case SpecialPriceDiscount:
		metaType = MetaTypeSpecialPriceDiscount
	case State:
		metaType = MetaTypeState
	case Store:
		metaType = MetaTypeStore
	case Supply:
		metaType = MetaTypeSupply
	case SupplyPosition:
		metaType = MetaTypeSupplyPosition
	case Task:
		metaType = MetaTypeTask
	case TaskNote:
		metaType = MetaTypeTaskNote
	case Uom:
		metaType = MetaTypeUom
	case Variant:
		metaType = MetaTypeVariant
	case Webhook:
		metaType = MetaTypeWebhook
	case CounterpartySettings:
		metaType = MetaTypeCounterpartySettings
	case UserSettings:
		metaType = MetaTypeUserSettings
	case Subscription:
		metaType = MetaTypeSubscription
	case Role:
		metaType = MetaTypeRole
	case PriceType:
		metaType = MetaTypePriceType
	case AssortmentSettings:
		metaType = MetaTypeAssortmentSettings
	case SalesChannel:
		metaType = MetaTypeSalesChannel
	case CounterPartyAdjustment:
		metaType = MetaTypeCounterPartyAdjustment
	case Metadata:
		metaType = MetaTypeMetadata
	case Token:
		metaType = MetaTypeToken
	case TaxRate:
		metaType = MetaTypeTaxRate
	case WebhookStock:
		metaType = MetaTypeWebhookStock
	case CustomTemplate:
		metaType = MetaTypeCustomTemplate
	case EmbeddedTemplate:
		metaType = MetaTypeEmbeddedTemplate
	case ProcessingPlanFolder:
		metaType = MetaTypeProcessingPlanFolder
	case Publication:
		metaType = MetaTypePublication
	case Attribute:
		metaType = MetaTypeAttribute
	case Async:
		metaType = MetaTypeAsync
	default:
		err = fmt.Errorf("unrecognized meta type %T", v)
	}
	return metaType, err
}
