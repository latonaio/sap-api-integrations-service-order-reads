package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-service-order-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
	ServiceOrder:                  data.ServiceOrder,
	ServiceOrderType:              data.ServiceOrderType,
	ServiceOrderUUID:              data.ServiceOrderUUID,
	ServiceOrderDescription:       data.ServiceOrderDescription,
	ServiceObjectType:             data.ServiceObjectType,
	Language:                      data.Language,
	ServiceDocumentPriority:       data.ServiceDocumentPriority,
	RequestedServiceStartDateTime: data.RequestedServiceStartDateTime,
	RequestedServiceEndDateTime:   data.RequestedServiceEndDateTime,
	ServiceDocChangedDateTime:     data.ServiceDocChangedDateTime,
	PurchaseOrderByCustomer:       data.PurchaseOrderByCustomer,
	CustomerPurchaseOrderDate:     data.CustomerPurchaseOrderDate,
	ServiceOrderIsReleased:        data.ServiceOrderIsReleased,
	ServiceOrderIsCompleted:       data.ServiceOrderIsCompleted,
	ServiceOrderIsRejected:        data.ServiceOrderIsRejected,
	SalesOrganization:             data.SalesOrganization,
	DistributionChannel:           data.DistributionChannel,
	Division:                      data.Division,
	SalesOffice:                   data.SalesOffice,
	SalesGroup:                    data.SalesGroup,
	SoldToParty:                   data.SoldToParty,
	ShipToParty:                   data.ShipToParty,
	BillToParty:                   data.BillToParty,
	PayerParty:                    data.PayerParty,
	ContactPerson:                 data.ContactPerson,
	ServiceDocGrossAmount:         data.ServiceDocGrossAmount,
	ServiceDocNetAmount:           data.ServiceDocNetAmount,
	ServiceDocTaxAmount:           data.ServiceDocTaxAmount,
	TransactionCurrency:           data.TransactionCurrency,
	ReferenceServiceRequest:       data.ReferenceServiceRequest,
	ReferenceServiceContract:      data.ReferenceServiceContract,
	ReferenceServiceOrder:         data.ReferenceServiceOrder,
	SrvcOrdCreditStatus:           data.SrvcOrdCreditStatus,
	ServiceOrderRejectionReason:   data.ServiceOrderRejectionReason,
    ToHeaderConfirmation:          data.ToHeaderConfirmation.Deferred.URI,
	ToHeaderDefect:                data.ToHeaderDefect.Deferred.URI,
	ToHeaderPersonResponsible:     data.ToHeaderPersonResponsible.Deferred.URI,
	ToHeaderReferenceObject:       data.ToHeaderReferenceObject.Deferred.URI,
    ToItem:                        data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{
	ServiceOrder:                   data.ServiceOrder,
	ServiceOrderItem:               data.ServiceOrderItem,
	ServiceOrderItemUUID:           data.ServiceOrderItemUUID,
	ServiceOrderItemDescription:    data.ServiceOrderItemDescription,
	ServiceObjectType:              data.ServiceObjectType,
	ServiceDocumentItemObjectType:  data.ServiceDocumentItemObjectType,
	Language:                       data.Language,
	Product:                        data.Product,
	Quantity:                       data.Quantity,
	QuantityUnit:                   data.QuantityUnit,
	SrvcOrdItemReservedQuantity:    data.SrvcOrdItemReservedQuantity,
	ServiceDuration:                data.ServiceDuration,
	ServiceDurationUnit:            data.ServiceDurationUnit,
	ServiceOrderItemCategory:       data.ServiceOrderItemCategory,
	ServiceOrdItemRejectionReason:  data.ServiceOrdItemRejectionReason,
	BillableControl:                data.BillableControl,
	SoldToParty:                    data.SoldToParty,
	ShipToParty:                    data.ShipToParty,
	BillToParty:                    data.BillToParty,
	PayerParty:                     data.PayerParty,
	ContactPersonBusinessPartnerID: data.ContactPersonBusinessPartnerID,
	PersonResponsible:              data.PersonResponsible,
	ExecutingServiceEmployee:       data.ExecutingServiceEmployee,
	ServicePerformer:               data.ServicePerformer,
	ServiceOrderItemIsReleased:     data.ServiceOrderItemIsReleased,
	ServiceOrderItemIsCompleted:    data.ServiceOrderItemIsCompleted,
	ServiceOrderItemIsRejected:     data.ServiceOrderItemIsRejected,
	ReferenceServiceContract:       data.ReferenceServiceContract,
	ReferenceServiceContractItem:   data.ReferenceServiceContractItem,
	ReferenceServiceRequest:        data.ReferenceServiceRequest,
	ReferenceServiceRequestItem:    data.ReferenceServiceRequestItem,
	ParentServiceOrderItem:         data.ParentServiceOrderItem,
	PlannedServiceStartDateTime:    data.PlannedServiceStartDateTime,
	PlannedServiceEndDateTime:      data.PlannedServiceEndDateTime,
	RequestedServiceStartDateTime:  data.RequestedServiceStartDateTime,
	RequestedServiceEndDateTime:    data.RequestedServiceEndDateTime,
	ActualServiceDuration:          data.ActualServiceDuration,
	ActualServiceDurationUnit:      data.ActualServiceDurationUnit,
	ReferenceServiceOrderItem:      data.ReferenceServiceOrderItem,
	ProfitCenter:                   data.ProfitCenter,
	SrvcOrdItemCreditStatus:        data.SrvcOrdItemCreditStatus,
	ToItemPricingElement:           data.ToItemPricingElement.Deferred.URI,
		})
	}

	return item, nil
}

func ConvertToToHeaderConfirmation(raw []byte, l *logger.Logger) ([]ToHeaderConfirmation, error) {
	pm := &responses.ToHeaderConfirmation{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToHeaderConfirmation. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toHeaderConfirmation := make([]ToHeaderConfirmation, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toHeaderConfirmation = append(toHeaderConfirmation, ToHeaderConfirmation{
	ServiceOrder:        data.ServiceOrder,
	ServiceConfirmation: data.ServiceConfirmation,
		})
	}

	return toHeaderConfirmation, nil
}

func ConvertToToHeaderDefect(raw []byte, l *logger.Logger) ([]ToHeaderDefect, error) {
	pm := &responses.ToHeaderDefect{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToHeaderDefect. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toHeaderDefect := make([]ToHeaderDefect, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toHeaderDefect = append(toHeaderDefect, ToHeaderDefect{
	ServiceOrder:                  data.ServiceOrder,
	SrvcDocTypeDefectCodeProfType: data.SrvcDocTypeDefectCodeProfType,
	ServiceDefectSequence:         data.ServiceDefectSequence,
	SrvcDocTypeDefectCodeProfile:  data.SrvcDocTypeDefectCodeProfile,
	ServiceDefectCodeCatalog:      data.ServiceDefectCodeCatalog,
	ServiceDefectCodeGroup:        data.ServiceDefectCodeGroup,
	ServiceDefectCode:             data.ServiceDefectCode,
	ServiceDefectSchema:           data.ServiceDefectSchema,
	ServiceDefectCategory:         data.ServiceDefectCategory,
		})
	}

	return toHeaderDefect, nil
}

func ConvertToToHeaderPersonResponsible(raw []byte, l *logger.Logger) ([]ToHeaderPersonResponsible, error) {
	pm := &responses.ToHeaderPersonResponsible{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToHeaderPersonResponsible. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toHeaderPersonResponsible := make([]ToHeaderPersonResponsible, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toHeaderPersonResponsible = append(toHeaderPersonResponsible, ToHeaderPersonResponsible{
	ServiceOrder:                 data.ServiceOrder,
	PersonResponsible:            data.PersonResponsible,
	CustMgmtPartnerIsMainPartner: data.CustMgmtPartnerIsMainPartner,
		})
	}

	return toHeaderPersonResponsible, nil
}

func ConvertToToHeaderReferenceObject(raw []byte, l *logger.Logger) ([]ToHeaderReferenceObject, error) {
	pm := &responses.ToHeaderReferenceObject{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToHeaderReferenceObject. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toHeaderReferenceObject := make([]ToHeaderReferenceObject, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toHeaderReferenceObject = append(toHeaderReferenceObject, ToHeaderReferenceObject{
	ServiceOrder:                 data.ServiceOrder,
	ServiceReferenceEquipment:    data.ServiceReferenceEquipment,
	ServiceRefFunctionalLocation: data.ServiceRefFunctionalLocation,
	SrvcRefObjIsMainObject:       data.SrvcRefObjIsMainObject,
		})
	}

	return toHeaderReferenceObject, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
	ServiceOrder:                   data.ServiceOrder,
	ServiceOrderItem:               data.ServiceOrderItem,
	ServiceOrderItemUUID:           data.ServiceOrderItemUUID,
	ServiceOrderItemDescription:    data.ServiceOrderItemDescription,
	ServiceObjectType:              data.ServiceObjectType,
	ServiceDocumentItemObjectType:  data.ServiceDocumentItemObjectType,
	Language:                       data.Language,
	Product:                        data.Product,
	Quantity:                       data.Quantity,
	QuantityUnit:                   data.QuantityUnit,
	SrvcOrdItemReservedQuantity:    data.SrvcOrdItemReservedQuantity,
	ServiceDuration:                data.ServiceDuration,
	ServiceDurationUnit:            data.ServiceDurationUnit,
	ServiceOrderItemCategory:       data.ServiceOrderItemCategory,
	ServiceOrdItemRejectionReason:  data.ServiceOrdItemRejectionReason,
	BillableControl:                data.BillableControl,
	SoldToParty:                    data.SoldToParty,
	ShipToParty:                    data.ShipToParty,
	BillToParty:                    data.BillToParty,
	PayerParty:                     data.PayerParty,
	ContactPersonBusinessPartnerID: data.ContactPersonBusinessPartnerID,
	PersonResponsible:              data.PersonResponsible,
	ExecutingServiceEmployee:       data.ExecutingServiceEmployee,
	ServicePerformer:               data.ServicePerformer,
	ServiceOrderItemIsReleased:     data.ServiceOrderItemIsReleased,
	ServiceOrderItemIsCompleted:    data.ServiceOrderItemIsCompleted,
	ServiceOrderItemIsRejected:     data.ServiceOrderItemIsRejected,
	ReferenceServiceContract:       data.ReferenceServiceContract,
	ReferenceServiceContractItem:   data.ReferenceServiceContractItem,
	ReferenceServiceRequest:        data.ReferenceServiceRequest,
	ReferenceServiceRequestItem:    data.ReferenceServiceRequestItem,
	ParentServiceOrderItem:         data.ParentServiceOrderItem,
	PlannedServiceStartDateTime:    data.PlannedServiceStartDateTime,
	PlannedServiceEndDateTime:      data.PlannedServiceEndDateTime,
	RequestedServiceStartDateTime:  data.RequestedServiceStartDateTime,
	RequestedServiceEndDateTime:    data.RequestedServiceEndDateTime,
	ActualServiceDuration:          data.ActualServiceDuration,
	ActualServiceDurationUnit:      data.ActualServiceDurationUnit,
	ReferenceServiceOrderItem:      data.ReferenceServiceOrderItem,
	ProfitCenter:                   data.ProfitCenter,
	SrvcOrdItemCreditStatus:        data.SrvcOrdItemCreditStatus,
	ToItemPricingElement:           data.ToItemPricingElement.Deferred.URI,
		})
	}

	return toItem, nil
}

func ConvertToToItemPricingElement(raw []byte, l *logger.Logger) ([]ToItemPricingElement, error) {
	pm := &responses.ToItemPricingElement{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemPricingElement. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItemPricingElement := make([]ToItemPricingElement, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemPricingElement = append(toItemPricingElement, ToItemPricingElement{
	ServiceOrder:            data.ServiceOrder,
	ServiceOrderItem:        data.ServiceOrderItem,
	PricingProcedureStep:    data.PricingProcedureStep,
	PricingProcedureCounter: data.PricingProcedureCounter,
	ConditionType:           data.ConditionType,
	ConditionRateValue:      data.ConditionRateValue,
	ConditionCurrency:       data.ConditionCurrency,
	ConditionQuantity:       data.ConditionQuantity,
	ConditionQuantityUnit:   data.ConditionQuantityUnit,
		})
	}

	return toItemPricingElement, nil
}
