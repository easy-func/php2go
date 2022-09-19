package in_array

import (
	"reflect"
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		array  []any
		length int
	}
	tests := []struct {
		name string
		args args
		want [][]any
	}{
		{
			name: "case - 1",
			args: args{
				array:  []any{1, 2},
				length: 1,
			},
			want: [][]any{
				[]any{1},
				[]any{2},
			},
		},
		{
			name: "case - length = 0",
			args: args{
				array:  []any{1, 2},
				length: 0,
			},
			want: [][]any{
				[]any{1},
				[]any{2},
			},
		},
		{
			name: "case - 1",
			args: args{
				array:  []any{},
				length: 2,
			},
			want: [][]any{},
		},
		{
			name: "case - 2",
			args: args{
				array:  []any{1, 2, 3},
				length: 2,
			},
			want: [][]any{
				[]any{1, 2},
				[]any{3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.array, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
