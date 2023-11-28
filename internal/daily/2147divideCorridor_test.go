package daily

import "testing"

func Test_numberOfWays(t *testing.T) {
	type args struct {
		corridor string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				corridor: "SSPPSPS",
			},
			want: 3,
		},
		{
			name: "Case 2",
			args: args{
				corridor: "PPSPSP",
			},
			want: 1,
		},
		{
			name: "Case 3",
			args: args{
				corridor: "S",
			},
			want: 0,
		},
		{
			name: "Case 4",
			args: args{
				corridor: "PPPPPPPPPPPPPPPPPPPP",
			},
			want: 0,
		},
		{
			name: "Case 5",
			args: args{
				corridor: "SSSSSSS",
			},
			want: 0,
		},
		{
			name: "Case 6",
			args: args{
				corridor: "PPSPSPPSPSPPSS",
			},
			want: 9,
		},
		{
			name: "Case 7",
			args: args{
				corridor: "PPPPPPSSPPPPPPP",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWays(tt.args.corridor); got != tt.want {
				t.Errorf("numberOfWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
