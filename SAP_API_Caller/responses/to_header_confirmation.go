package responses

type ToHeaderConfirmation struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ServiceOrder        string `json:"ServiceOrder"`
			ServiceConfirmation string `json:"ServiceConfirmation"`
		} `json:"results"`
	} `json:"d"`
}
