package main

import (
	"testing"
)

func TestStringToTokens(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1 + 21 * 33",
			args: args{s: "1 + 21 * 33"},
			want: []string{"1", "+", "21", "*", "33"},
		},
		{
			name: "1 + 21 * 33 - 4",
			args: args{s: "1 + 21 * 33 - 4"},
			want: []string{"1", "+", "21", "*", "33", "-", "4"},
		},
		{
			name: "1 + 21 * 33 - 4 / 5",
			args: args{s: "1 + 21 * 33 - 4 / 5"},
			want: []string{"1", "+", "21", "*", "33", "-", "4", "/", "5"},
		},
		{
			name: "4 ^ 3 + 21 * 33 - 4 / 5 + 6",
			args: args{s: "4 ^ 3 + 21 * 33 - 4 / 5 + 6"},
			want: []string{"4", "^", "3", "+", "21", "*", "33", "-", "4", "/", "5", "+", "6"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringToTokens(tt.args.s)
			if len(got) != len(tt.want) {
				t.Errorf("stringToTokens() = %v, want %v", got, tt.want)
				return
			}
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("stringToTokens() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}

func TestConvRPN(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1 + 21 * 33",
			args: args{s: "1 + 21 * 33"},
			want: "1 21 33 * +",
		},
		{
			name: "1 + 21 * 33 - 4",
			args: args{s: "1 + 21 * 33 - 4"},
			want: "1 21 33 * + 4 -",
		},
		{
			name: "1 + 21 * 33 - 4 / 5",
			args: args{s: "1 + 21 * 33 - 4 / 5"},
			want: "1 21 33 * + 4 5 / -",
		},
		{
			name: "4 ^ 3 + 21 * 33 - 4 / 5 + 6",
			args: args{s: "4 ^ 3 + 21 * 33 - 4 / 5 + 6"},
			want: "4 3 ^ 21 33 * + 4 5 / - 6 +",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvRPN(tt.args.s)
			if got != tt.want {
				t.Errorf("ConvRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}
