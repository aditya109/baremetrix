package greenpincreatepinallocation

// Response for GreenPin Create Pin Allocation POST API
type Response struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	A                PartyNumberResponseInfo `json:"a"`
	B                PartyNumberResponseInfo `json:"b"`
	DeallocationTime string                  `json:"deallocation_time"`
	GreenPinId       string                  `json:"greenpin_id"`
	TransactionId    string                  `json:"transaction_id"`
	Usage            string                  `json:"usage"`
}

type PartyNumberResponseInfo struct {
	Pin     string   `json:"pin"`
	VNS     []string `json:"vns"`
	Numbers []string `json:"numbers"`
}
