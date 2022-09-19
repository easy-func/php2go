package in_array

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		needle any
		array  []any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case-int",
			args: args{
				needle: 1,
				array:  []any{1, 2},
			},
			want: true,
		},
		{
			name: "case-int-false",
			args: args{
				needle: 1,
				array:  []any{"1", 2},
			},
			want: false,
		},
		{
			name: "case-int-false",
			args: args{
				needle: 1,
				array:  []any{},
			},
			want: false,
		},
		{
			name: "case-int-false",
			args: args{
				needle: 1.0,
				array:  []any{"1.00"},
			},
			want: false,
		},
		{
			name: "case-double",
			args: args{
				needle: 1.0,
				array:  []any{1.00},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.needle, tt.args.array); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
