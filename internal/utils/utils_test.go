package utils

import (
	"reflect"
	"testing"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
)

func TestFindItemFromList(t *testing.T) {
	type args struct {
		list       interface{}
		searchTerm string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Test-with-GroupType-parameter",
			args: args{
				list: []models.GraphType{
					{
						Name:      "rpm_vs_latency",
						IsEnabled: true,
					},
					{
						Name:      "rpm_vs_5xx",
						IsEnabled: true,
					},
					{
						Name:      "rpm_vs_timeout",
						IsEnabled: true,
					},
				},
				searchTerm: "rpm_vs_latency",
			},
			want: models.GraphType{
				Name:      "rpm_vs_latency",
				IsEnabled: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindItemFromListWithKey(tt.args.list, tt.args.searchTerm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindItemFromListWithKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindItemFromListWithKeyWithAllocateURLType(t *testing.T) {
	type args struct {
		list interface{}
		key  string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindItemFromListWithKey(tt.args.list, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindItemFromListWithKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
