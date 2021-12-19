package responses

type ToItemPricingElement struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ServiceOrder            string `json:"ServiceOrder"`
			ServiceOrderItem        string `json:"ServiceOrderItem"`
			PricingProcedureStep    string `json:"PricingProcedureStep"`
			PricingProcedureCounter string `json:"PricingProcedureCounter"`
			ConditionType           string `json:"ConditionType"`
			ConditionRateValue      string `json:"ConditionRateValue"`
			ConditionCurrency       string `json:"ConditionCurrency"`
			ConditionQuantity       string `json:"ConditionQuantity"`
			ConditionQuantityUnit   string `json:"ConditionQuantityUnit"`
		} `json:"results"`
	} `json:"d"`
}
