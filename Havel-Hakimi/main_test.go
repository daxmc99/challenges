package main

import "testing"

func Test_havelHakimi(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want bool
	}{
		{
			"test1",
			[]int{5, 3, 0, 2, 0, 7, 2, 5},
			false,
		},
		{
			"test2",
			[]int{4, 2, 0, 1, 5, 0},
			false,
		},
		{
			"test4",
			[]int{3, 1, 2, 3, 1, 0},
			true,
		},
		{
			"test5",
			[]int{16, 9, 9, 15, 9, 7, 9, 11, 17, 11, 4, 9, 12, 14, 14, 12, 17, 0, 3, 16},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := havelHakimi(tt.arg); got != tt.want {
				t.Errorf("havelHakimi() = %v, want %v", got, tt.want)
			}
		})
	}
}
