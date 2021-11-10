package main

import (
	"reflect"
	"testing"
)

func Test_newDownloadArgs(t *testing.T) {
	type args struct {
		rawRepoURL string
		outDirPath string
	}
	tests := []struct {
		name    string
		args    args
		want    *downloadArgs
		wantErr bool
	}{

		{
			name: "with query repo",
			args: args{
				rawRepoURL: "http://github.com/goleveldb/goleveldb.git?foo=bar",
				outDirPath: "",
			},
			want: &downloadArgs{
				repoURL:    "http://github.com/goleveldb/goleveldb.git",
				outDirPath: "/home/wangsaiyu/go/src/github.com/goleveldb/goleveldb",
			},
			wantErr: false,
		},
		{
			name: "with schema repo",
			args: args{
				rawRepoURL: "http://github.com/goleveldb/goleveldb.git",
				outDirPath: "",
			},
			want: &downloadArgs{
				repoURL:    "http://github.com/goleveldb/goleveldb.git",
				outDirPath: "/home/wangsaiyu/go/src/github.com/goleveldb/goleveldb",
			},
			wantErr: false,
		},
		{
			name: "no schema repo",
			args: args{
				rawRepoURL: "github.com/goleveldb/goleveldb",
				outDirPath: "",
			},
			want: &downloadArgs{
				repoURL:    "http://github.com/goleveldb/goleveldb.git",
				outDirPath: "/home/wangsaiyu/go/src/github.com/goleveldb/goleveldb",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newDownloadArgs(tt.args.rawRepoURL, tt.args.outDirPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("newDownloadArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDownloadArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
