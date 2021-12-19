package responses

type Item struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
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
			ToItemPricingElement struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PricingElement"`
		} `json:"results"`
	} `json:"d"`
}
