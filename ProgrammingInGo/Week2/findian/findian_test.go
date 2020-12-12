package main

import "testing"

func Test_check(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "one", args: args{in: "ian"}, want: true},
		{name: "two", args: args{in: "Ian"}, want: true},
		{name: "three", args: args{in: "iuiygaygn"}, want: true},
		{name: "four", args: args{in: "I d skd a efju N"}, want: true},
		{name: "five", args: args{in: "ihhhhhn"}, want: false},
		{name: "six", args: args{in: "ina"}, want: false},
		{name: "seven", args: args{in: "xian"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.in); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
