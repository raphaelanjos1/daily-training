package main

import "testing"

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "example1",
			arr:  []int{1, 2, 2, 1, 1, 3},
			want: true,
		},
		{
			name: "example1",
			arr:  []int{1, 2},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := uniqueOccurrences(tt.arr)
			if got != tt.want {
				t.Fatalf("numIdenticalPairs(%v) = %t; want %t", tt.arr, got, tt.want)
			}
		})
	}

}
