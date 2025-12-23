package peers

import (
	"net"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    int
		wantErr bool
	}{
		{
			name: "valid two peers",
			input: []byte{
				192, 0, 2, 123, 0x1A, 0xE1,
				192, 0, 2, 124, 0x1A, 0xE2,
			},
			want:    2,
			wantErr: false,
		},
		{
			name:    "invalid length",
			input:   []byte{192, 0, 2, 123, 0x1A},
			want:    0,
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   []byte{},
			want:    0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			peers, err := Unmarshal(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(peers) != tt.want {
				t.Errorf("Unmarshal() got %d peers, want %d", len(peers), tt.want)
			}

			if !tt.wantErr && len(peers) == 2 {
				if peers[0].IP.String() != "192.0.2.123" {
					t.Errorf("First peer IP = %s, want 192.0.2.123", peers[0].IP.String())
				}
				if peers[0].Port != 6881 {
					t.Errorf("First peer Port = %d, want 6881", peers[0].Port)
				}
			}
		})
	}
}

func TestPeerString(t *testing.T) {
	p := Peer{
		IP:   net.IPv4(192, 168, 1, 1),
		Port: 6881,
	}

	got := p.String()
	want := "192.168.1.1:6881"

	if got != want {
		t.Errorf("Peer.String() = %s, want %s", got, want)
	}
}
