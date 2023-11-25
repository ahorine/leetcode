package l75

import "testing"

func Test_gcdOfStrings(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Case 1",
			args{"ABCABC", "ABC"},
			"ABC",
		},
		{
			"Case 2",
			args{"ABABAB", "ABAB"},
			"AB",
		},
		{
			"Case 3",
			args{"LEET", "CODE"},
			"",
		},
		{
			"Case 4",
			args{"LEET", "LEET"},
			"LEET",
		},
		{
			"Case 5",
			args{"LEET", "CODER"},
			"",
		},
		{
			"Case 6",
			args{"AAA", "AAAAA"},
			"A",
		},
		{
			"Case 7",
			args{"ADDERSADDERS", "ADDERSADDERSADDERS"},
			"ADDERS",
		},
		{
			"Case 8",
			args{"", "ABC"},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcdOfStrings(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("gcdOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
