package helper

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetAbsolutePathForCorrectRelativePath(t *testing.T) {
	result, err := GetAbsolutePath(`/config/config.json`)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	var expectedResult string
	cwd, err := os.Getwd()
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	expectedResult = filepath.Join(strings.Split(cwd, "pkg/helper")[0], "config/config.json")
	if result != expectedResult {
		t.Errorf("got %q, wanted %q", result, expectedResult)
	}
}

func TestGetAbsolutePath(t *testing.T) {
	type args struct {
		relPath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "checking if correct absolute path is obtained, if file path is provided",
			args: args{
				relPath: "/play/tenants/zomato127/play_greenvn_post.json",
			},
			want:    "/home/adityakumar/go_workspace/src/bitbucket.org/exotel/ent-leadassist-test/baremetrix/play/tenants/zomato127/play_greenvn_post.json",
			wantErr: false,
		},
		{
			name: "checking if correct absolute path is obtained, if directory path is provided",
			args: args{
				relPath: "/play/tenants/zomato127",
			},
			want:    "/home/adityakumar/go_workspace/src/bitbucket.org/exotel/ent-leadassist-test/baremetrix/play/tenants/zomato127",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAbsolutePath(tt.args.relPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAbsolutePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAbsolutePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
