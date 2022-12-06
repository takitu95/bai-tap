package main

import (
	"testing"
)

func TestPrintTeamNumber(t *testing.T) {
	type args struct {
		Number int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintTeamNumber(tt.args.Number); got != tt.want {
				t.Errorf("PrintTeamNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
