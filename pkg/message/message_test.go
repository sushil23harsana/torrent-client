package message

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestSerialize(t *testing.T) {
	msg := &Message{
		ID:      MsgHave,
		Payload: []byte{1, 2, 3, 4},
	}
	buf := msg.Serialize()

	if len(buf) != 9 {
		t.Errorf("Expected 9 bytes, got %d", len(buf))
	}
	if buf[4] != byte(MsgHave) {
		t.Errorf("Expected ID %d, got %d", MsgHave, buf[4])
	}
}

func TestRead(t *testing.T) {
	data := []byte{0, 0, 0, 5, 4, 1, 2, 3, 4}
	msg, err := Read(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if msg.ID != MsgHave {
		t.Errorf("Expected ID %d, got %d", MsgHave, msg.ID)
	}
	if len(msg.Payload) != 4 {
		t.Errorf("Expected payload length 4, got %d", len(msg.Payload))
	}
}

func TestFormatRequest(t *testing.T) {
	msg := FormatRequest(5, 100, 16384)
	if msg.ID != MsgRequest {
		t.Error("Wrong ID")
	}
	if len(msg.Payload) != 12 {
		t.Errorf("Expected 12 bytes, got %d", len(msg.Payload))
	}
}

func TestParsePiece(t *testing.T) {
	buf := make([]byte, 100)
	payload := make([]byte, 13)
	binary.BigEndian.PutUint32(payload[0:4], 0)
	binary.BigEndian.PutUint32(payload[4:8], 10)
	copy(payload[8:], []byte{1, 2, 3, 4, 5})

	msg := &Message{
		ID:      MsgPiece,
		Payload: payload,
	}

	n, err := ParsePiece(0, buf, msg)
	if err != nil {
		t.Fatal(err)
	}
	if n != 5 {
		t.Errorf("Expected 5 bytes copied, got %d", n)
	}
}

func TestParseHave(t *testing.T) {
	payload := make([]byte, 4)
	binary.BigEndian.PutUint32(payload, 5)

	msg := &Message{
		ID:      MsgHave,
		Payload: payload,
	}

	index, err := ParseHave(msg)
	if err != nil {
		t.Fatal(err)
	}
	if index != 5 {
		t.Errorf("Expected index 5, got %d", index)
	}
}
