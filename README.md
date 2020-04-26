torrent-remove-trackers
=======================

Very simple command line tool for removing all trackers from torrent file(s).

Requirements: Go compiler 1.13 or later

Installation:
```
git clone https://github.com/snabb/torrent-remove-trackers.git
cd torrent-remove-trackers
go build .
```

Usage:
```
torrent-remove-trackers foo1.torrent ..
```

Uses Matt Joiner's excellent torrent library: https://github.com/anacrolix/torrent
