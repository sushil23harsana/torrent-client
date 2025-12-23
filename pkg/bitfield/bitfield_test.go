package bitfield

import "testing"

func TestHasPiece(t *testing.T) {
	bf := Bitfield{0b01010100, 0b01010100}

	tests := []struct {
		input int
		want  bool
	}{
		{0, false}, {1, true}, {2, false}, {3, true},
		{4, false}, {5, true}, {6, false}, {7, false},
	}

	for _, tt := range tests {
		got := bf.HasPiece(tt.input)
		if got != tt.want {
			t.Errorf("HasPiece(%d) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestSetPiece(t *testing.T) {
	bf := Bitfield{0b00000000, 0b00000000}
	bf.SetPiece(1)
	bf.SetPiece(9)

	if !bf.HasPiece(1) {
		t.Error("Failed to set piece 1")
	}
	if !bf.HasPiece(9) {
		t.Error("Failed to set piece 9")
	}
	if bf.HasPiece(0) {
		t.Error("Piece 0 should not be set")
	}
}
