package main

import "testing"

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
		{"1+2*3-4/2+2^3-1", args{"1 2 3 * + 4 2 / - 2 3 ^ + 1 -"}, 12},
		{"-5+1+2*3-4/2+2^3-1", args{"-5 1 + 2 3 * + 4 2 / - 2 3 ^ + 1 -"}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calc(tt.args.rpn); got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
