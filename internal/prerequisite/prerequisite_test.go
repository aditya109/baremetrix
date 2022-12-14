package prerequisite

import (
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models/api/greenvncreateallocation"
	"testing"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
)

func TestPreallocateGreenVNIds(t *testing.T) {
	type args struct {
		config             *models.Config
		act                models.Act
		apiType            string
		specialRequirement int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-without-special-requirement",
			args: args{
				config: &models.Config{
					Instance: models.Instance{
						Specs: models.Specs{
							PrerequisiteSpecs: models.PrerequisiteSpecs{
								HttpTimeoutInSeconds:   20,
								SupportedBatchSize:     300,
								ConditionalBatchBuffer: 2,
							},
						},
					},
				},
				act: models.Act{
					Name: "act-00",
					Allocateurl: []models.Allocateurl{
						{
							UseCase: "LA_GREENVN_DELETE_ALLOCATION",
							URL:     "http://leadassist-staging.exotel.in/v1/tenants/SID/greenvn/GREENVN_ID",
						},
						{
							UseCase: "LA_GREENVN_GET_ALLOCATION_DETAILS",
							URL:     "http://leadassist-staging.exotel.in/v1/tenants/SID/greenvn/GREENVN_ID",
						},
						{
							UseCase: "LA_GREENVN_CREATE_ALLOCATION",
							URL:     "http://leadassist-staging.exotel.in/v1/tenants/SID/greenvn",
						},
					},
					Api:    "LA_GREENVN_DELETE_ALLOCATION",
					Method: "DELETE",
					Headers: []models.Header{
						{
							Key:   "Content-type",
							Value: "application/json",
						},
					},
					Endpoint: models.Endpoint{
						AuthorizationType:  "Basic",
						AuthorizationToken: "ZXhvdGVsdDpleG90ZWx0",
						Sid:                "exotelt",
					},
					PreOpSequence: []string{"LA_GREENVN_CREATE_ALLOCATION"},
					DefaultParameters: models.PayloadParameters{
						Seed: models.SeedLimit{
							Apin: models.RangeLimit{
								Min: 1000000000,
								Max: 7000000000,
							},
							Bpin: models.RangeLimit{
								Min: 1000000000,
								Max: 9000000000,
							},
							ConnectionId: models.RangeLimit{
								Min: 0,
								Max: 1000000000000000000,
							},
						},
						CountryCode: "+91",
						Usage:       "twoway",
						DeallocationPolicy: greenvncreateallocation.DeallocationPolicy{
							Duration: "2m",
						},
					},
					Vegeta: models.Vegeta{
						RateOfRequests:        100,
						DurationInSeconds:     60,
						DuplicacyPercentage:   10,
						TimeoutInMilliSeconds: 10000,
						KeepAlive:             false,
					},
				},
				apiType:            "LA_GREENVN_DELETE_ALLOCATION",
				specialRequirement: 0,
			},
			want: 100,
		},
		{
			name: "test-with-special-requirement",
			args: args{
				config: &models.Config{
					Instance: models.Instance{
						Specs: models.Specs{
							PrerequisiteSpecs: models.PrerequisiteSpecs{
								HttpTimeoutInSeconds:   20,
								SupportedBatchSize:     150,
								ConditionalBatchBuffer: 2,
							},
						},
					},
				},
				act: models.Act{
					Name: "act-00",
					Allocateurl: []models.Allocateurl{
						{
							UseCase: "LA_GREENVN_DELETE_ALLOCATION",
							URL:     "http://leadassist-staging.exotel.in/v1/tenants/SID/greenvn/GREENVN_ID",
						},
						{
							UseCase: "LA_GREENVN_GET_ALLOCATION_DETAILS",
							URL:     "http://leadassist-staging.exotel.in/v1/tenants/SID/greenvn/GREENVN_ID",
						},
						{
							UseCase: "LA_GREENVN_CREATE_ALLOCATION",
							URL:     "http://leadassist-staging.exotel.in/v1/tenants/SID/greenvn",
						},
					},
					Api:    "LA_GREENVN_DELETE_ALLOCATION",
					Method: "DELETE",
					Headers: []models.Header{
						{
							Key:   "Content-type",
							Value: "application/json",
						},
					},
					Endpoint: models.Endpoint{
						AuthorizationType:  "Basic",
						AuthorizationToken: "ZXhvdGVsdDpleG90ZWx0",
						Sid:                "exotelt",
					},
					PreOpSequence: []string{"LA_GREENVN_CREATE_ALLOCATION"},
					DefaultParameters: models.PayloadParameters{
						Seed: models.SeedLimit{
							Apin: models.RangeLimit{
								Min: 1000000000,
								Max: 7000000000,
							},
							Bpin: models.RangeLimit{
								Min: 1000000000,
								Max: 9000000000,
							},
							ConnectionId: models.RangeLimit{
								Min: 0,
								Max: 1000000000000000000,
							},
						},
						CountryCode: "+91",
						Usage:       "twoway",
						DeallocationPolicy: greenvncreateallocation.DeallocationPolicy{
							Duration: "2m",
						},
					},
					Vegeta: models.Vegeta{
						RateOfRequests:        100,
						DurationInSeconds:     60,
						DuplicacyPercentage:   10,
						TimeoutInMilliSeconds: 10000,
						KeepAlive:             false,
					},
				},
				apiType:            "LA_GREENVN_DELETE_ALLOCATION",
				specialRequirement: 10,
			},
			want: 10,
		},
		{
			name: "test-with-special-requirement-but-exceeding-supported-batch-size",
			args: args{
				config: &models.Config{
					Instance: models.Instance{
						Specs: models.Specs{
							PrerequisiteSpecs: models.PrerequisiteSpecs{
								HttpTimeoutInSeconds:   20,
								SupportedBatchSize:     150,
								ConditionalBatchBuffer: 2,
							},
						},
					},
				},
				act: models.Act{
					Name: "act-00",
					Allocateurl: []models.Allocateurl{
						{
							UseCase: "LA_GREENVN_DELETE_ALLOCATION",
							URL:     "http://leadassist-staging.exotel.in/v1/tenants/SID/greenvn/GREENVN_ID",
						},
						{
							UseCase: "LA_GREENVN_GET_ALLOCATION_DETAILS",
							URL:     "http://leadassist-staging.exotel.in/v1/tenants/SID/greenvn/GREENVN_ID",
						},
						{
							UseCase: "LA_GREENVN_CREATE_ALLOCATION",
							URL:     "http://leadassist-staging.exotel.in/v1/tenants/SID/greenvn",
						},
					},
					Api:    "LA_GREENVN_DELETE_ALLOCATION",
					Method: "DELETE",
					Headers: []models.Header{
						{
							Key:   "Content-type",
							Value: "application/json",
						},
					},
					Endpoint: models.Endpoint{
						AuthorizationType:  "Basic",
						AuthorizationToken: "ZXhvdGVsdDpleG90ZWx0",
						Sid:                "exotelt",
					},
					PreOpSequence: []string{"LA_GREENVN_CREATE_ALLOCATION"},
					DefaultParameters: models.PayloadParameters{
						Seed: models.SeedLimit{
							Apin: models.RangeLimit{
								Min: 1000000000,
								Max: 7000000000,
							},
							Bpin: models.RangeLimit{
								Min: 1000000000,
								Max: 9000000000,
							},
							ConnectionId: models.RangeLimit{
								Min: 0,
								Max: 1000000000000000000,
							},
						},
						CountryCode: "+91",
						Usage:       "twoway",
						DeallocationPolicy: greenvncreateallocation.DeallocationPolicy{
							Duration: "2m",
						},
					},
					Vegeta: models.Vegeta{
						RateOfRequests:        100,
						DurationInSeconds:     60,
						DuplicacyPercentage:   10,
						TimeoutInMilliSeconds: 10000,
						KeepAlive:             false,
					},
				},
				apiType:            "LA_GREENVN_DELETE_ALLOCATION",
				specialRequirement: 500,
			},
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PreallocateGreenVNIds(tt.args.config, tt.args.act, tt.args.specialRequirement, tt.args.apiType)
			if tt.args.specialRequirement != 0 && len(got) < tt.want {
				t.Errorf("PreallocateGreenVNIds() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestGetGreenVNCreateAllocationTuplesConcurrently1(t *testing.T) {
	type args struct {
		config             *models.Config
		act                models.Act
		specialRequirement int
	}
	tests := []struct {
		name string
		args args
		want []greenvncreateallocation.Tuple
	}{
		{
			name: "",
			args: args{
				config: &models.Config{
					Instance: models.Instance{
						Specs: models.Specs{
							PrerequisiteSpecs: models.PrerequisiteSpecs{
								HttpTimeoutInSeconds:   660,
								SupportedBatchSize:     20,
								ConditionalBatchBuffer: 2,
							},
						},
					},
				},
				act: models.Act{
					Name: "",
					Allocateurl: []models.Allocateurl{
						{
							UseCase: "LA_GREENVN_CREATE_ALLOCATION",
							URL:     "http://leadassist-staging.exotel.in/v1/tenants/SID/greenvn",
						},
					},
					Api:    "LA_GREENVN_CREATE_ALLOCATION",
					Method: "POST",
					Headers: []models.Header{
						{
							Key:   "Content-type",
							Value: "application/json",
						},
					},
					Endpoint: models.Endpoint{
						AuthorizationType:  "Basic",
						AuthorizationToken: "ZXhvdGVsdDpleG90ZWx0",
						Sid:                "exotelt",
					},
					PreOpSequence: nil,
					DefaultParameters: models.PayloadParameters{
						Seed: models.SeedLimit{
							Apin: models.RangeLimit{
								Min: 1000000000,
								Max: 7000000000,
							},
							Bpin: models.RangeLimit{
								Min: 1000000000,
								Max: 9000000000,
							},
							ConnectionId: models.RangeLimit{
								Min: 0,
								Max: 1000000000000000000,
							},
						},
						CountryCode: "+91",
						Usage:       "twoway",
						DeallocationPolicy: greenvncreateallocation.DeallocationPolicy{
							Duration: "1h",
						},
					},
					Vegeta: models.Vegeta{
						RateOfRequests:        10,
						DurationInSeconds:     5,
						DuplicacyPercentage:   10,
						TimeoutInMilliSeconds: 10000,
						KeepAlive:             true,
					},
				},
				specialRequirement: 100,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGreenVNCreateAllocationTuplesConcurrently(tt.args.config, tt.args.act, tt.args.specialRequirement); !got[0].Response.Success {
				t.Errorf("GetGreenVNCreateAllocationTuplesConcurrently() = %v", got)
			}
		})
	}
}
