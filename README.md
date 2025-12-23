# BitTorrent Client

A complete BitTorrent client implementation written in Go. This client supports downloading files via the BitTorrent protocol with multi-peer connections, piece verification, and tracker communication.

## Features

- **Multi-peer downloading** - Connects to multiple peers simultaneously for faster downloads
- **Piece verification** - Uses SHA-1 hashes to verify downloaded pieces
- **Tracker communication** - Communicates with trackers to discover peers
- **Progress tracking** - Real-time download progress with percentage completion
- **Error resilience** - Handles peer timeouts and network issues gracefully
- **Clean architecture** - Modular design with separate packages for different components

## Requirements

- Go 1.20 or later
- Internet connection for tracker communication

## Installation

### Method 1: Clone and Build
```bash
git clone https://github.com/sushil23harsana/torrent-client.git
cd torrent-client
go build -o torrent-client ./cmd/torrent-client
```

### Method 2: Direct Install
```bash
go install github.com/sushil23harsana/torrent-client/cmd/torrent-client@latest
```

### Method 3: Run from Source
```bash
git clone https://github.com/sushil23harsana/torrent-client.git
cd torrent-client
go run ./cmd/torrent-client <torrent-file> <output-file>
```

## Usage

### Basic Usage
```bash
./torrent-client <path-to-torrent-file> <output-file-path>
```

### Examples

**Download a Linux ISO:**
```bash
./torrent-client ubuntu-22.04.torrent ubuntu-22.04.iso
```

**Download to specific directory:**
```bash
./torrent-client ~/Downloads/debian.torrent ~/Downloads/debian-12.iso
```

**Using absolute paths:**
```bash
./torrent-client "C:\torrents\movie.torrent" "C:\Downloads\movie.mp4"
```

### Command Line Arguments

- **torrent-file**: Path to the `.torrent` file containing metadata
- **output-file**: Path where you want to save the downloaded file

## How It Works

1. **Parse Torrent File** - Reads and decodes the `.torrent` file to extract metadata
2. **Contact Tracker** - Communicates with the tracker to get a list of peers
3. **Connect to Peers** - Establishes connections with multiple peers simultaneously  
4. **Download Pieces** - Downloads file pieces from different peers in parallel
5. **Verify Integrity** - Validates each piece using SHA-1 hashes
6. **Assemble File** - Combines verified pieces into the final file

## Project Structure

```
├── cmd/torrent-client/    # Main application entry point
├── pkg/
│   ├── bitfield/         # Bitfield operations for piece tracking
│   ├── client/           # Peer client connection handling
│   ├── handshake/        # BitTorrent handshake protocol
│   ├── message/          # BitTorrent protocol messages
│   ├── p2p/              # Peer-to-peer download coordination
│   ├── peers/            # Peer discovery and management
│   └── torrentfile/      # Torrent file parsing and tracker communication
├── go.mod                # Go module definition
└── README.md            # This file
```

## Sample Output

```bash
$ ./torrent-client debian.torrent debian.iso
2025/12/23 18:41:15 Starting download for debian-12.11.0-armel-netinst.iso
2025/12/23 18:41:16 Completed handshake with 178.62.85.20
2025/12/23 18:41:16 Completed handshake with 91.121.91.13
2025/12/23 18:41:16 Completed handshake with 176.9.57.163
2025/12/23 18:41:17 (0.05%) Downloaded piece #0 from 18 peers
2025/12/23 18:41:17 (0.10%) Downloaded piece #1 from 18 peers
...
2025/12/23 18:44:13 (100.00%) Downloaded piece #2019 from 16 peers
2025/12/23 18:44:14 Download complete!
```

## Dependencies

- [`github.com/jackpal/bencode-go`](https://github.com/jackpal/bencode-go) - Bencode encoding/decoding for torrent files

## Protocol Support

This client implements core BitTorrent protocol features:
- BitTorrent Protocol v1.0
- Piece-based downloading
- Peer Wire Protocol messages
- Tracker HTTP/HTTPS communication
- SHA-1 piece verification

## Limitations

- Single-file torrents only (no multi-file torrent support)
- HTTP/HTTPS trackers only (no UDP tracker support)
- No DHT (Distributed Hash Table) support
- No peer exchange (PEX) support
- No magnet link support


## Acknowledgments

- BitTorrent protocol specification
- Go standard library documentation
- Various BitTorrent client implementations for reference

## Troubleshooting

### Common Issues

**"no such file or directory"**
- Ensure the torrent file path is correct
- Use quotes around paths with spaces

**"connection refused" or timeout errors**
- Check your internet connection
- Some trackers may be temporarily unavailable
- Firewall may be blocking connections

**"permission denied" when saving file**
- Ensure you have write permissions to the output directory
- Try running with appropriate permissions

**Very slow download speeds**
- Limited number of available peers
- Network conditions
- Tracker may have few active peers

For more issues, please check the [Issues](https://github.com/sushil23harsana/torrent-client/issues) page.