package main

import "testing"

func TestJewelsAndStones(t *testing.T) {
	tests := []struct {
		name   string
		jewels string
		stones string
		want   int
	}{
		{
			name:   "example1",
			jewels: "aA",
			stones: "aAAbbbb",
			want:   3,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := numJewelsInStones(tt.jewels, tt.stones)
			if got != tt.want {
				t.Fatalf("numJewelsInStones(%v, %v) = %d; want %d", tt.jewels, tt.stones, got, tt.want)
			}
		})
	}

}
