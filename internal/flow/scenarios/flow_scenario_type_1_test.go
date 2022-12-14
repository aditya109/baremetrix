package scenarios

import (
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils/targetstack"
	"reflect"
	"testing"
)

func TestGetStackListFromFlowScenarioType1(t *testing.T) {
	type args struct {
		config *models.Config
		flow   models.Flow
	}
	tests := []struct {
		name string
		args args
		want []targetstack.FlowStack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStackListFromFlowScenarioType1(tt.args.config, tt.args.flow); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStackListFromFlowScenarioType1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFlowStack(t *testing.T) {
	type args struct {
		parameters flowStackParameters
	}
	tests := []struct {
		name    string
		args    args
		want    targetstack.FlowStack
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFlowStack(tt.args.parameters)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFlowStack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFlowStack() got = %v, want %v", got, tt.want)
			}
		})
	}
}
