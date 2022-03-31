package hasuragraphql



type FindAcceptJson struct {
	Data struct {
		Accept []struct {
			Accepter string `json:"accepter"`
			Post     string `json:"post"`
		} `json:"accept"`
	} `json:"data"`
}