package daily

import (
	"reflect"
	"testing"
)

func Test_getSumAbsoluteDifferences(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Case 1",
			args{[]int{2, 3, 5}},
			[]int{4, 3, 5},
		},
		{
			"Case 2",
			args{[]int{1, 4, 6, 8, 10}},
			[]int{24, 15, 13, 15, 21},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSumAbsoluteDifferences(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSumAbsoluteDifferences() = %v, want %v", got, tt.want)
			}
		})
	}
}
