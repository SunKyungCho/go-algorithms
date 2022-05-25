package tree

import "testing"

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	firstNode := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{3, nil, nil},
			&TreeNode{4, nil, nil},
		},
		&TreeNode{2,
			&TreeNode{4, nil, nil},
			&TreeNode{3, nil, nil},
		},
	}
	secondNode := &TreeNode{1,
		&TreeNode{2,
			nil,
			&TreeNode{3, nil, nil},
		},
		&TreeNode{2,
			nil,
			&TreeNode{3, nil, nil},
		},
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"true", args{firstNode}, true},
		{"false", args{secondNode}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
