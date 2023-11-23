package daily

import (
	"reflect"
	"testing"
)

func Test_checkArithmeticSubarrays(t *testing.T) {
	type args struct {
		nums []int
		l    []int
		r    []int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "Case 1",
			args: args{
				nums: []int{4, 6, 5, 9, 3, 7},
				l:    []int{0, 0, 2},
				r:    []int{2, 3, 5},
			},
			want: []bool{true, false, true},
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10},
				l:    []int{0, 1, 6, 4, 8, 7},
				r:    []int{4, 4, 9, 7, 9, 10},
			},
			want: []bool{false, true, false, false, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkArithmeticSubarrays(tt.args.nums, tt.args.l, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkArithmeticSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isArithmetic(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{
				nums: []int{4, 6, 5},
			},
			want: true,
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{4, 6, 5, 9},
			},
			want: false,
		},
		{
			name: "Case 3",
			args: args{
				nums: []int{5, 9, 3, 7},
			},
			want: true,
		},
		{
			name: "Case 4",
			args: args{
				nums: []int{5, 9},
			},
			want: true,
		},
		{
			name: "Case 5",
			args: args{
				nums: []int{5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArithmetic(tt.args.nums); got != tt.want {
				t.Errorf("isArithmetic() = %v, want %v", got, tt.want)
			}
		})
	}
}
