package handshake

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var infoHash, peerID [20]byte
	copy(infoHash[:], "12345678901234567890")
	copy(peerID[:], "ABCDEFGHIJ1234567890")

	h := New(infoHash, peerID)

	if h.Pstr != "BitTorrent protocol" {
		t.Errorf("Expected 'BitTorrent protocol', got %s", h.Pstr)
	}
}

func TestSerialize(t *testing.T) {
	var infoHash, peerID [20]byte
	h := New(infoHash, peerID)
	buf := h.Serialize()

	if len(buf) != 68 {
		t.Errorf("Expected 68 bytes, got %d", len(buf))
	}
	if buf[0] != 19 {
		t.Errorf("Expected pstrlen 19, got %d", buf[0])
	}
}

func TestRead(t *testing.T) {
	var infoHash, peerID [20]byte
	copy(infoHash[:], "12345678901234567890")
	copy(peerID[:], "ABCDEFGHIJ1234567890")

	h := New(infoHash, peerID)
	serialized := h.Serialize()

	parsed, err := Read(bytes.NewReader(serialized))
	if err != nil {
		t.Fatal(err)
	}

	if parsed.Pstr != "BitTorrent protocol" {
		t.Errorf("Expected 'BitTorrent protocol', got %s", parsed.Pstr)
	}
	if !bytes.Equal(parsed.InfoHash[:], infoHash[:]) {
		t.Error("InfoHash mismatch")
	}
}
