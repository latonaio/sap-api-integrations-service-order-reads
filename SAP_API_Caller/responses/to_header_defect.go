package responses

type ToHeaderDefect struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ServiceOrder                  string `json:"ServiceOrder"`
			SrvcDocTypeDefectCodeProfType string `json:"SrvcDocTypeDefectCodeProfType"`
			ServiceDefectSequence         int    `json:"ServiceDefectSequence"`
			SrvcDocTypeDefectCodeProfile  string `json:"SrvcDocTypeDefectCodeProfile"`
			ServiceDefectCodeCatalog      string `json:"ServiceDefectCodeCatalog"`
			ServiceDefectCodeGroup        string `json:"ServiceDefectCodeGroup"`
			ServiceDefectCode             string `json:"ServiceDefectCode"`
			ServiceDefectSchema           string `json:"ServiceDefectSchema"`
			ServiceDefectCategory         string `json:"ServiceDefectCategory"`
		} `json:"results"`
	} `json:"d"`
}
