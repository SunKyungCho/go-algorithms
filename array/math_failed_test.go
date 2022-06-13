package array

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		answers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{1, 2, 3, 4, 5}}, []int{1}},
		{"", args{[]int{1, 3, 2, 4, 2}}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.answers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
