package in_array

import (
	"reflect"
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		mapArray map[string]any
		caseTo   CaseTo
	}
	tests := []struct {
		name string
		args args
		want map[string]any
	}{
		{
			name: "case - 小写转大写",
			args: args{
				mapArray: map[string]any{
					"a": 1,
				},
				caseTo: CaseUpper,
			},
			want: map[string]any{
				"A": 1,
			},
		},
		{
			name: "case - 大写转小写",
			args: args{
				mapArray: map[string]any{
					"A": 1,
				},
				caseTo: CaseLower,
			},
			want: map[string]any{
				"a": 1,
			},
		},
		{
			name: "case - 大小写转小写",
			args: args{
				mapArray: map[string]any{
					"AaA": 1,
				},
				caseTo: CaseLower,
			},
			want: map[string]any{
				"aaa": 1,
			},
		},
		{
			name: "case - 传值出错，自动转大写",
			args: args{
				mapArray: map[string]any{
					"AaA": 1,
				},
				caseTo: CaseTo(100),
			},
			want: map[string]any{
				"AAA": 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.mapArray, tt.args.caseTo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
