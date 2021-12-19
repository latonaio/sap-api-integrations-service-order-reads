package responses

type ToHeaderPersonResponsible struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ServiceOrder                 string `json:"ServiceOrder"`
			PersonResponsible            string `json:"PersonResponsible"`
			CustMgmtPartnerIsMainPartner bool   `json:"CustMgmtPartnerIsMainPartner"`
		} `json:"results"`
	} `json:"d"`
}
