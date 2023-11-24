package daily

import "testing"

func Test_maxCoins(t *testing.T) {
	type args struct {
		piles []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{[]int{2, 4, 1, 2, 7, 8}},
			want: 9,
		},
		{
			name: "Case 2",
			args: args{[]int{2, 4, 5}},
			want: 4,
		},
		{
			name: "Case 3",
			args: args{[]int{9, 8, 7, 6, 5, 1, 2, 3, 4}},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoins(tt.args.piles); got != tt.want {
				t.Errorf("maxCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
