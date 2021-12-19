package responses

type ToHeaderReferenceObject struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ServiceOrder                 string `json:"ServiceOrder"`
			ServiceReferenceEquipment    string `json:"ServiceReferenceEquipment"`
			ServiceRefFunctionalLocation string `json:"ServiceRefFunctionalLocation"`
			SrvcRefObjIsMainObject       bool   `json:"SrvcRefObjIsMainObject"`
		} `json:"results"`
	} `json:"d"`
}
