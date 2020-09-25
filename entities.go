package sila

func (client Client) CheckHandle(userHandle string) *CheckHandle {
	return &CheckHandle{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "header_msg",
	}
}

func (client Client) Register(userHandle string) *Register {
	return &Register{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "entity_msg",
	}
}

func (client Client) RequestKyc(userHandle string) *RequestKyc {
	return &RequestKyc{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "header_msg",
	}
}

func (client Client) CheckKyc(userHandle string) *CheckKyc {
	return &CheckKyc{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "header_msg",
	}
}

type SuccessResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
}