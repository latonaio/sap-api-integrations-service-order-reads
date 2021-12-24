package sap_api_output_formatter

type ServiceOrder struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	ServiceOrder  string `json:"service_order"`
	Deleted       bool   `json:"deleted"`
}    
    
type Header struct {
	ServiceOrder                  string      `json:"ServiceOrder"`
	ServiceOrderType              string      `json:"ServiceOrderType"`
	ServiceOrderUUID              string      `json:"ServiceOrderUUID"`
	ServiceOrderDescription       string      `json:"ServiceOrderDescription"`
	ServiceObjectType             string      `json:"ServiceObjectType"`
	Language                      string      `json:"Language"`
	ServiceDocumentPriority       string      `json:"ServiceDocumentPriority"`
	RequestedServiceStartDateTime string      `json:"RequestedServiceStartDateTime"`
	RequestedServiceEndDateTime   string      `json:"RequestedServiceEndDateTime"`
	ServiceDocChangedDateTime     string      `json:"ServiceDocChangedDateTime"`
	PurchaseOrderByCustomer       string      `json:"PurchaseOrderByCustomer"`
	CustomerPurchaseOrderDate     string      `json:"CustomerPurchaseOrderDate"`
	ServiceOrderIsReleased        string      `json:"ServiceOrderIsReleased"`
	ServiceOrderIsCompleted       string      `json:"ServiceOrderIsCompleted"`
	ServiceOrderIsRejected        string      `json:"ServiceOrderIsRejected"`
	SalesOrganization             string      `json:"SalesOrganization"`
	DistributionChannel           string      `json:"DistributionChannel"`
	Division                      string      `json:"Division"`
	SalesOffice                   string      `json:"SalesOffice"`
	SalesGroup                    string      `json:"SalesGroup"`
	SoldToParty                   string      `json:"SoldToParty"`
	ShipToParty                   string      `json:"ShipToParty"`
	BillToParty                   string      `json:"BillToParty"`
	PayerParty                    string      `json:"PayerParty"`
	ContactPerson                 string      `json:"ContactPerson"`
	ServiceDocGrossAmount         string      `json:"ServiceDocGrossAmount"`
	ServiceDocNetAmount           string      `json:"ServiceDocNetAmount"`
	ServiceDocTaxAmount           string      `json:"ServiceDocTaxAmount"`
	TransactionCurrency           string      `json:"TransactionCurrency"`
	ReferenceServiceRequest       string      `json:"ReferenceServiceRequest"`
	ReferenceServiceContract      string      `json:"ReferenceServiceContract"`
	ReferenceServiceOrder         string      `json:"ReferenceServiceOrder"`
	SrvcOrdCreditStatus           string      `json:"SrvcOrdCreditStatus"`
	ServiceOrderRejectionReason   string      `json:"ServiceOrderRejectionReason"`
    ToHeaderConfirmation          string      `json:"to_Confirmation"`
	ToHeaderDefect                string      `json:"to_Defect"`
	ToHeaderPersonResponsible     string      `json:"to_PersonResponsible"`
	ToHeaderReferenceObject       string      `json:"to_ReferenceObject"`
    ToItem                        string      `json:"to_Item"`
}

type Item struct {
	ServiceOrder                   string      `json:"ServiceOrder"`
	ServiceOrderItem               string      `json:"ServiceOrderItem"`
	ServiceOrderItemUUID           string      `json:"ServiceOrderItemUUID"`
	ServiceOrderItemDescription    string      `json:"ServiceOrderItemDescription"`
	ServiceObjectType              string      `json:"ServiceObjectType"`
	ServiceDocumentItemObjectType  string      `json:"ServiceDocumentItemObjectType"`
	Language                       string      `json:"Language"`
	Product                        string      `json:"Product"`
	Quantity                       string      `json:"Quantity"`
	QuantityUnit                   string      `json:"QuantityUnit"`
	SrvcOrdItemReservedQuantity    string      `json:"SrvcOrdItemReservedQuantity"`
	ServiceDuration                string      `json:"ServiceDuration"`
	ServiceDurationUnit            string      `json:"ServiceDurationUnit"`
	ServiceOrderItemCategory       string      `json:"ServiceOrderItemCategory"`
	ServiceOrdItemRejectionReason  string      `json:"ServiceOrdItemRejectionReason"`
	BillableControl                string      `json:"BillableControl"`
	SoldToParty                    string      `json:"SoldToParty"`
	ShipToParty                    string      `json:"ShipToParty"`
	BillToParty                    string      `json:"BillToParty"`
	PayerParty                     string      `json:"PayerParty"`
	ContactPersonBusinessPartnerID string      `json:"ContactPersonBusinessPartnerId"`
	PersonResponsible              string      `json:"PersonResponsible"`
	ExecutingServiceEmployee       string      `json:"ExecutingServiceEmployee"`
	ServicePerformer               string      `json:"ServicePerformer"`
	ServiceOrderItemIsReleased     string      `json:"ServiceOrderItemIsReleased"`
	ServiceOrderItemIsCompleted    string      `json:"ServiceOrderItemIsCompleted"`
	ServiceOrderItemIsRejected     string      `json:"ServiceOrderItemIsRejected"`
	ReferenceServiceContract       string      `json:"ReferenceServiceContract"`
	ReferenceServiceContractItem   string      `json:"ReferenceServiceContractItem"`
	ReferenceServiceRequest        string      `json:"ReferenceServiceRequest"`
	ReferenceServiceRequestItem    string      `json:"ReferenceServiceRequestItem"`
	ParentServiceOrderItem         string      `json:"ParentServiceOrderItem"`
	PlannedServiceStartDateTime    string      `json:"PlannedServiceStartDateTime"`
	PlannedServiceEndDateTime      string      `json:"PlannedServiceEndDateTime"`
	RequestedServiceStartDateTime  string      `json:"RequestedServiceStartDateTime"`
	RequestedServiceEndDateTime    string      `json:"RequestedServiceEndDateTime"`
	ActualServiceDuration          string      `json:"ActualServiceDuration"`
	ActualServiceDurationUnit      string      `json:"ActualServiceDurationUnit"`
	ReferenceServiceOrderItem      string      `json:"ReferenceServiceOrderItem"`
	ProfitCenter                   string      `json:"ProfitCenter"`
	SrvcOrdItemCreditStatus        string      `json:"SrvcOrdItemCreditStatus"`
	ToItemPricingElement           string      `json:"to_PricingElement"`
}

type ToHeaderConfirmation struct {
	ServiceOrder        string `json:"ServiceOrder"`
	ServiceConfirmation string `json:"ServiceConfirmation"`
}

type ToHeaderDefect struct {
	ServiceOrder                  string `json:"ServiceOrder"`
	SrvcDocTypeDefectCodeProfType string `json:"SrvcDocTypeDefectCodeProfType"`
	ServiceDefectSequence         int    `json:"ServiceDefectSequence"`
	SrvcDocTypeDefectCodeProfile  string `json:"SrvcDocTypeDefectCodeProfile"`
	ServiceDefectCodeCatalog      string `json:"ServiceDefectCodeCatalog"`
	ServiceDefectCodeGroup        string `json:"ServiceDefectCodeGroup"`
	ServiceDefectCode             string `json:"ServiceDefectCode"`
	ServiceDefectSchema           string `json:"ServiceDefectSchema"`
	ServiceDefectCategory         string `json:"ServiceDefectCategory"`
}

type ToHeaderPersonResponsible struct {
	ServiceOrder                 string `json:"ServiceOrder"`
	PersonResponsible            string `json:"PersonResponsible"`
	CustMgmtPartnerIsMainPartner bool   `json:"CustMgmtPartnerIsMainPartner"`
}

type ToHeaderReferenceObject struct {
	ServiceOrder                 string `json:"ServiceOrder"`
	ServiceReferenceEquipment    string `json:"ServiceReferenceEquipment"`
	ServiceRefFunctionalLocation string `json:"ServiceRefFunctionalLocation"`
	SrvcRefObjIsMainObject       bool   `json:"SrvcRefObjIsMainObject"`
}

type ToItem struct {
	ServiceOrder                   string      `json:"ServiceOrder"`
	ServiceOrderItem               string      `json:"ServiceOrderItem"`
	ServiceOrderItemUUID           string      `json:"ServiceOrderItemUUID"`
	ServiceOrderItemDescription    string      `json:"ServiceOrderItemDescription"`
	ServiceObjectType              string      `json:"ServiceObjectType"`
	ServiceDocumentItemObjectType  string      `json:"ServiceDocumentItemObjectType"`
	Language                       string      `json:"Language"`
	Product                        string      `json:"Product"`
	Quantity                       string      `json:"Quantity"`
	QuantityUnit                   string      `json:"QuantityUnit"`
	SrvcOrdItemReservedQuantity    string      `json:"SrvcOrdItemReservedQuantity"`
	ServiceDuration                string      `json:"ServiceDuration"`
	ServiceDurationUnit            string      `json:"ServiceDurationUnit"`
	ServiceOrderItemCategory       string      `json:"ServiceOrderItemCategory"`
	ServiceOrdItemRejectionReason  string      `json:"ServiceOrdItemRejectionReason"`
	BillableControl                string      `json:"BillableControl"`
	SoldToParty                    string      `json:"SoldToParty"`
	ShipToParty                    string      `json:"ShipToParty"`
	BillToParty                    string      `json:"BillToParty"`
	PayerParty                     string      `json:"PayerParty"`
	ContactPersonBusinessPartnerID string      `json:"ContactPersonBusinessPartnerId"`
	PersonResponsible              string      `json:"PersonResponsible"`
	ExecutingServiceEmployee       string      `json:"ExecutingServiceEmployee"`
	ServicePerformer               string      `json:"ServicePerformer"`
	ServiceOrderItemIsReleased     string      `json:"ServiceOrderItemIsReleased"`
	ServiceOrderItemIsCompleted    string      `json:"ServiceOrderItemIsCompleted"`
	ServiceOrderItemIsRejected     string      `json:"ServiceOrderItemIsRejected"`
	ReferenceServiceContract       string      `json:"ReferenceServiceContract"`
	ReferenceServiceContractItem   string      `json:"ReferenceServiceContractItem"`
	ReferenceServiceRequest        string      `json:"ReferenceServiceRequest"`
	ReferenceServiceRequestItem    string      `json:"ReferenceServiceRequestItem"`
	ParentServiceOrderItem         string      `json:"ParentServiceOrderItem"`
	PlannedServiceStartDateTime    string      `json:"PlannedServiceStartDateTime"`
	PlannedServiceEndDateTime      string      `json:"PlannedServiceEndDateTime"`
	RequestedServiceStartDateTime  string      `json:"RequestedServiceStartDateTime"`
	RequestedServiceEndDateTime    string      `json:"RequestedServiceEndDateTime"`
	ActualServiceDuration          string      `json:"ActualServiceDuration"`
	ActualServiceDurationUnit      string      `json:"ActualServiceDurationUnit"`
	ReferenceServiceOrderItem      string      `json:"ReferenceServiceOrderItem"`
	ProfitCenter                   string      `json:"ProfitCenter"`
	SrvcOrdItemCreditStatus        string      `json:"SrvcOrdItemCreditStatus"`
	ToItemPricingElement           string      `json:"to_PricingElement"`
}

type ToItemPricingElement struct {
	ServiceOrder            string `json:"ServiceOrder"`
	ServiceOrderItem        string `json:"ServiceOrderItem"`
	PricingProcedureStep    string `json:"PricingProcedureStep"`
	PricingProcedureCounter string `json:"PricingProcedureCounter"`
	ConditionType           string `json:"ConditionType"`
	ConditionRateValue      string `json:"ConditionRateValue"`
	ConditionCurrency       string `json:"ConditionCurrency"`
	ConditionQuantity       string `json:"ConditionQuantity"`
	ConditionQuantityUnit   string `json:"ConditionQuantityUnit"`
}
