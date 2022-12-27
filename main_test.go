package main

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestIsPrime",
			args: args{
				n: 5,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.args.n); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
