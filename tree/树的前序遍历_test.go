package tree

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				/*
					    1
					  2   3
					1 2 3
					2 1 3
				*/
				root: buildTree([]int{1, 2, 3}, []int{2, 1, 3}),
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorderTraversal4(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				/*
					    1
					  2   3
					1 2 3
					2 1 3
				*/
				root: buildTree([]int{1, 2, 3}, []int{2, 1, 3}),
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal4(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal2() = %v, want %v", got, tt.want)
			}
		})
	}
}
