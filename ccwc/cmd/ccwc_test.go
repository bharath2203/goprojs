package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"testing"
)

func Test_handle(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "file",
			args: args{
				args: []string{"file.txt"},
				cmd:  nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(os.Getwd())
			if err := handle(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func BenchmarkHandle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := handle(nil, []string{"file.txt"})
		if err != nil {
			return 
		}
	}
}

func BenchmarkHandleParallel(b *testing.B) {
	options.runInAsync = true
	for i := 0; i < b.N; i++ {
		err := handle(nil, []string{"file.txt"})
		if err != nil {
			return 
		}
	}
}

