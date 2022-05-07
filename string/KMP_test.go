package string

import (
	"fmt"
	"testing"
)

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				haystack: "",
				needle:   "",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNext(t *testing.T) {
	needle := "aabaaf"
	n := len(needle)
	next := make([]int, n)
	getNext(next, needle)
	fmt.Println(next)
}
