package greenvncreateallocation

// Payload is the model for GreenVN - Create Allocation API request payload.
type Payload struct {
	ConnectionID       string             `json:"connection_id"`
	ApartyNumbers      []string           `json:"aparty_numbers"`
	BpartyNumbers      []string           `json:"bparty_numbers"`
	ApartyPins         []int              `json:"aparty_pins"`
	BpartyPins         []int              `json:"bparty_pins"`
	DeallocationPolicy DeallocationPolicy `json:"deallocation_policy"`
	Usage              string             `json:"usage"`
	Strictness         string             `json:"strictness,omitempty"`
	Preferences        Preferences        `json:"preferences,omitempty"`
}

// DeallocationPolicy is used to schedule an automatic deallocation. If not set, default deallocation period (as in configuration) is used.
type DeallocationPolicy struct {
	Duration string `json:"duration"`
}

// Preferences a set of preferred options that the returned greenvn should satisfy.
type Preferences struct {
	Greenvn string `json:"greenvn,omitempty"`
	Region  string `json:"region,omitempty"`
	Type    string `json:"type,omitempty"`
}
