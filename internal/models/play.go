package models

import "bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models/api/greenvncreateallocation"

// Play is a scenario which is to be run by the load-testing framework,
// consists of acts which can either by parallelly or sequentially.
type Play struct {
	Id         string `json:"id"`              // Specifies the id of the play.
	Name       string `json:"name"`            // Specifies the name of the play.
	Type       string `json:"type"`            // Specifies the type of the play - internal or tenant.
	Tenant     string `json:"tenant"`          // Specifies the tenant name - internal or specifiec tenant name.
	Iterations int    `json:"iterations"`      // Specifies the number of interations the given play file has to run.
	Acts       []Act  `json:"acts,omitempty"`  // Specifies the acts in the playfile.
	Flows      []Flow `json:"flows,omitempty"` // Specifies the flows in the playfile.
}

// Act is an individual step in the scenario, again which can either by asynchronously, without no inherent ordering.
type Act struct {
	Name              string            `json:"name"`                      // Specifies the name of act.
	Allocateurl       []Allocateurl     `json:"allocate_url"`              // Specifies the list of allocate urls supported by the given act step.
	Api               string            `json:"api"`                       // Specifies the api type.
	Method            string            `json:"method"`                    // Specifies the method name.
	Headers           []Header          `json:"headers"`                   // Specifies the headers to be put.
	Endpoint          Endpoint          `json:"endpoint"`                  // Specifies the endpoint information to be put within the act.
	PreOpSequence     []string          `json:"pre_op_sequence,omitempty"` // Specifies the prerequisite api which has to called upon before actual load testing.
	DefaultParameters PayloadParameters `json:"default_parameters"`        // Specifies the default parameters for generating payload.
	Vegeta            Vegeta            `json:"vegeta"`                    // Specifies the parameters for vegeta.
}

// Allocateurl is the entity which contains usecase(api type) and corresponding URL.
type Allocateurl struct {
	UseCase string `json:"use_case"`
	URL     string `json:"url"`
}

// Header is a map of Headers which is a part of Http call, to be made in an act.
type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Vegeta contains all the tweak-able parameters provided by the load-testing framework, which can be dynamically changed by the user.
type Vegeta struct {
	RateOfRequests        int  `json:"rate_of_requests,omitempty"`        // Specifies the rate of requests per second.
	DurationInSeconds     int  `json:"duration_in_seconds,omitempty"`     // Specifies the duration of an act in seconds.
	DuplicacyPercentage   int  `json:"duplicacy_percentage,omitempty"`    // TBD: not implemented
	TimeoutInMilliSeconds int  `json:"timeout_in_milliseconds,omitempty"` // Specifies the vegeta timeout in milli seconds.
	KeepAlive             bool `json:"keep_alive,omitempty"`              // Specifies the keep alive connections for the current client. (keep it true for production-type scenarios)
}

// Endpoint contains the credentials which is required for request authorization.
type Endpoint struct {
	AuthorizationType  string `json:"authorization_type"`
	AuthorizationToken string `json:"authorization_token"`
	Sid                string `json:"sid"`
}

// PayloadParameters contains parameters which can be tweaked for varying payloads for each act.
type PayloadParameters struct {
	Seed               SeedLimit                                  `json:"seed,omitempty"`
	CountryCode        string                                     `json:"country_code"`
	Usage              string                                     `json:"usage"`
	DeallocationPolicy greenvncreateallocation.DeallocationPolicy `json:"deallocation_policy"`
	Strictness         string                                     `json:"strictness,omitempty"`
	Preferences        greenvncreateallocation.Preferences        `json:"preferences,omitempty"`
}

// SeedLimit contains payload field-specific seeds.
type SeedLimit struct {
	Apin          RangeLimit `json:"apin,omitempty"`
	Bpin          RangeLimit `json:"bpin,omitempty"`
	ConnectionId  RangeLimit `json:"connection_id,omitempty"`
	TransactionId RangeLimit `json:"transaction_id,omitempty"`
	PinLength     RangeLimit `json:"pin_length,omitempty"`
	VNS           RangeLimit `json:"vns,omitempty"`
	Numbers       RangeLimit `json:"numbers,omitempty"`
}

// RangeLimit contains parameters for setting seed for randomization.
type RangeLimit struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
}

// Flow is an individual step in the scenario, again which can either by asynchronously, with inherent ordering.
type Flow struct {
	Name     string   `json:"name"`     // Specifies the name of flow.
	FlowId   string   `json:"flow_id"`  // Specifies the id of the flow.
	Scenes   []Scene  `json:"scenes"`   // Specifies the scenes in the playfile.
	Endpoint Endpoint `json:"endpoint"` // Specifies the endpoint information to be put within the act for overall flow.
	Vegeta   Vegeta   `json:"vegeta"`   // Specifies the parameters for vegeta.
}

// Scene is an flow step in particular flow.
type Scene struct {
	Allocateurl       string            `json:"allocate_url"`              // Specifies the list of allocate urls supported by the given flow step.
	Api               string            `json:"api"`                       // Specifies the api type.
	Method            string            `json:"method"`                    // Specifies the method name.
	Headers           []Header          `json:"headers"`                   // Specifies the headers to be put.
	Endpoint          Endpoint          `json:"endpoint"`                  // Specifies the endpoint information to be put within the act, for scene-specific.
	PreOpSequence     []string          `json:"pre_op_sequence,omitempty"` // Specifies the prerequisite api which has to called upon before actual load testing.
	DefaultParameters PayloadParameters `json:"default_parameters"`        // Specifies the default parameters for generating payload.
	ExpectedRPS       int               `json:"expected_rps,omitempty"`    // Specifies the anticipated rps for the target set generation.
}
