package main

import "testing"

func TestGoodPairs(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example1",
			nums: []int{1, 2, 3, 1, 1, 3},
			want: 4,
		},
		{
			name: "example1",
			nums: []int{1, 1, 1, 1},
			want: 6,
		},
		{
			name: "example1",
			nums: []int{1, 2, 3},
			want: 0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := numIdenticalPairs(tt.nums)
			if got != tt.want {
				t.Fatalf("numIdenticalPairs(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}

}
