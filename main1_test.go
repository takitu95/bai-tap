package main

import "testing"

func TestKiemtrahieuduonghaykhong(t *testing.T) {
	type args struct {
		Numbera int
		Numberb int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happycase",
			args: args{
				Numbera: 1,
				Numberb: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Kiemtrahieuduonghaykhong(tt.args.Numbera, tt.args.Numberb); got != tt.want {
				t.Errorf("Kiemtrahieuduonghaykhong() = %v, want %v", got, tt.want)
			}
		})
	}
}
