package dto

type XMLData struct {
	Channel struct {
		Post struct {
			INACTENDD   string `json:"INACTENDD"`
			Error       string `json:"Error"`
			RETAILER    string `json:"RETAILER"`
			CLASS       string `json:"CLASS"`
			ACTENDD     string `json:"ACTENDD"`
			ADMINST     string `json:"ADMINST"`
			CREDITVIOCE string `json:"CREDITVIOCE"`
			CODE        string `json:"CODE"`
			PHONE       string `json:"PHONE"`
			RBAL        string `json:"RBAL"`
		} `xml:"item"`
	} `xml:"channel"`
}
