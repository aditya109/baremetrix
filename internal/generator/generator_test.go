package generator

import (
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils"
	"testing"
)

func TestPopRandomItemFromList(t *testing.T) {

	type args struct {
		list *[]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := utils.PopRandomItemFromList(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopRandomItemFromList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PopRandomItemFromList() = %v, want %v", got, tt.want)
			}
		})
	}
}
