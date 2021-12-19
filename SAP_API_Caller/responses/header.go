package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
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
			ToHeaderConfirmation          struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Confirmation"`
			ToHeaderDefect struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Defect"`
			ToHeaderPersonResponsible struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PersonResponsible"`
			ToHeaderReferenceObject struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_ReferenceObject"`
			ToItem struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Item"`
		} `json:"results"`
	} `json:"d"`
}
