package string

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"a", args{s: "egg", t: "add"}, true},
		{"a", args{s: "foo", t: "bar"}, false},
		{"a", args{s: "paper", t: "title"}, true},
		{"a", args{s: "abcd", t: "efgh"}, true},
		{"a", args{s: "badc", t: "baba"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
