package string

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test1", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"Test2", args{[]string{"dog", "racecar", "car"}}, ""},
		{"a", args{[]string{"a"}}, "a"},
		{"ab, a", args{[]string{"ab", "a"}}, "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
