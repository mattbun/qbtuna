# qbtuna

My own quick-and-dirty script to update qbittorrent's listen port with the current port-forwarded port from gluetun.

## Configuration

Use the following environment variables to configure it:

* `GLUETUN_HOST` - The URL of the gluetun control server, for example `http://localhost:8000`
* `QBITTORRENT_HOST` - The URL of qbittorrent, for example `http://localhost:8080`
* `QBITTORRENT_USERNAME` - The username to log into qbittorrent with
* `QBITTORRENT_PASSWORD` - The password to log into qbittorrent with
* `INTERVAL_S` - The update interval, in seconds (defaults to 60 if unset)
