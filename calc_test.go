package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if add(1, 2) != 3 {
		t.Fatalf("3 is expected, but %v\n", add(1, 2))
	}
}

func TestSub(t *testing.T) {
	if sub(1, 2) != -1 {
		t.Fatalf("-1 is expected, but %v\n", sub(1, 2))
	}
}

func TestMul(t *testing.T) {
	if mul(1, 2) != 2 {
		t.Fatalf("2 is expected, but %v\n", mul(1, 2))
	}
}

func TestDiv(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1/2", args{1, 2}, 0},
		{"2/1", args{2, 1}, 2},
		{"2/2", args{2, 2}, 1},
		{"2/3", args{2, 3}, 0},
		{"3/2", args{3, 2}, 1},
		{"3/3", args{3, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := div(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPow(t *testing.T) {
	type args struct {
		x int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2^0", args{2, 0}, 1},
		{"2^3", args{2, 3}, 8},
		{"2^4", args{2, 4}, 16},
		{"2^5", args{2, 5}, 32},
		{"2^6", args{2, 6}, 64},
		{"2^7", args{2, 7}, 128},
		{"2^24", args{2, 24}, 16777216},
		{"2^50", args{2, 50}, 1125899906842624},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComb(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4C2",
			args: args{a: 4, b: 2},
			want: 6,
		},
		{
			name: "4C3",
			args: args{a: 4, b: 3},
			want: 4,
		},
		{
			name: "4C0",
			args: args{a: 4, b: 0},
			want: 1,
		},
		{
			name: "4C4",
			args: args{a: 4, b: 4},
			want: 1,
		},
		{
			name: "10C5",
			args: args{a: 10, b: 5},
			want: 252,
		},
		{
			name: "10C0",
			args: args{a: 10, b: 0},
			want: 1,
		},
		{
			name: "10C10",
			args: args{a: 10, b: 10},
			want: 1,
		},
		{
			name: "30C15",
			args: args{a: 30, b: 15},
			want: 155117520,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := comb(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Comb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerm(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4P2",
			args: args{a: 4, b: 2},
			want: 12,
		},
		{
			name: "4P3",
			args: args{a: 4, b: 3},
			want: 24,
		},
		{
			name: "4P0",
			args: args{a: 4, b: 0},
			want: 1,
		},
		{
			name: "4P4",
			args: args{a: 4, b: 4},
			want: 24,
		},
		{
			name: "10P5",
			args: args{a: 10, b: 5},
			want: 30240,
		},
		{
			name: "10P0",
			args: args{a: 10, b: 0},
			want: 1,
		},
		{
			name: "10P10",
			args: args{a: 10, b: 10},
			want: 3628800,
		},
		{
			name: "20P10",
			args: args{a: 20, b: 10},
			want: 670442572800,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := perm(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Perm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFact(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0!", args{0}, 1},
		{"1!", args{1}, 1},
		{"2!", args{2}, 2},
		{"3!", args{3}, 6},
		{"4!", args{4}, 24},
		{"5!", args{5}, 120},
		{"6!", args{6}, 720},
		{"7!", args{7}, 5040},
		{"8!", args{8}, 40320},
		{"9!", args{9}, 362880},
		{"10!", args{10}, 3628800},
		{"11!", args{11}, 39916800},
		{"12!", args{12}, 479001600},
		{"13!", args{13}, 6227020800},
		{"14!", args{14}, 87178291200},
		{"15!", args{15}, 1307674368000},
		{"16!", args{16}, 20922789888000},
		{"17!", args{17}, 355687428096000},
		{"18!", args{18}, 6402373705728000},
		{"19!", args{19}, 121645100408832000},
		{"20!", args{20}, 2432902008176640000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fact(tt.args.n); got != tt.want {
				t.Errorf("Fact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalc(t *testing.T) {
	type args struct {
		rpn string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1+2", args{"1 2 +"}, 3},
		{"1-2", args{"1 2 -"}, -1},
		{"1*2", args{"1 2 *"}, 2},
		{"1/2", args{"1 2 /"}, 0},
		{"2^3", args{"2 3 ^"}, 8},
		{"4C2", args{"4 2 C"}, 6},
		{"4P2", args{"4 2 P"}, 12},
		{"1+2*3-4/2+2^3-1", args{"1 2 3 * + 4 2 / - 2 3 ^ + 1 -"}, 12},
		{"-5+1+2*3-4/2+2^3-1", args{"-5 1 + 2 3 * + 4 2 / - 2 3 ^ + 1 -"}, 7},
		{"-5+1+2*3-4/2+2^3-1+2C1", args{"-5 1 + 2 3 * + 4 2 / - 2 3 ^ + 1 - 2 1 C +"}, 9},
		{"-5+1+2*3-4/2+2^3-1+2C1-2P1", args{"-5 1 + 2 3 * + 4 2 / - 2 3 ^ + 1 - 2 1 C + 2 1 P -"}, 7},
		{"-5+1+2*3-4/2+2^3-1+2C1-2P1+2^3", args{"-5 1 + 2 3 * + 4 2 / - 2 3 ^ + 1 - 2 1 C + 2 1 P - 2 3 ^ +"}, 15},
		{"-5+1+2*3-4/2+2^3-1+2C1-2P1+2^3-3!", args{"-5 1 + 2 3 * + 4 2 / - 2 3 ^ + 1 - 2 1 C + 2 1 P - 2 3 ^ + 3 ! -"}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calc(tt.args.rpn); got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
