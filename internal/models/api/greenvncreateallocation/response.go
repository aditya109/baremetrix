package greenvncreateallocation

// Response represents the response body from Create Allocation API
type Response struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

// Data represents the JSON payload of the response body which houses the various kinds of information required post-allocation
type Data struct {
	ApartyPins    []int    `json:"aparty_pins"`
	BpartyPins    []int    `json:"bparty_pins"`
	ApartyNumbers []string `json:"aparty_numbers"`
	BpartyNumbers []string `json:"bparty_numbers"`
	ConnectionID  string   `json:"connection_id"`
	GreenVN       string   `json:"greenvn"`
	State         string   `json:"state"`
	GreenVNID     string   `json:"greenvn_id"`
	Usage         string   `json:"usage"`
}
