package greenpincreatepinallocation

import "bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models/api/greenvncreateallocation"

// Payload for GreenPin Create Pin Allocation POST API
type Payload struct {
	TransactionId      string                                     `json:"transaction_id"`
	A                  PartyNumberRequestInfo                     `json:"a"`
	B                  PartyNumberRequestInfo                     `json:"b"`
	Usage              string                                     `json:"usage"`
	DeallocationPolicy greenvncreateallocation.DeallocationPolicy `json:"Deallocation_policy"`
}

type PartyNumberRequestInfo struct {
	PinLength int      `json:"pin_length"`
	VNS       []string `json:"vns"`
	Numbers   []string `json:"numbers"`
}
