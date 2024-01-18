# qbtuna

My own quick-and-dirty script to update qbittorrent's listen port with the current port-forwarded port from gluetun.

## Features

* Uses gluetun's control server, not the deprecated `forwarded_port` file
* No dependencies!

## Configuration

Use the following environment variables to configure it:

* `GLUETUN_HOST` - The URL of the gluetun control server, for example `http://localhost:8000`
* `QBITTORRENT_HOST` - The URL of qbittorrent, for example `http://localhost:8080`
* `QBITTORRENT_USERNAME` - The username to log into qbittorrent with
* `QBITTORRENT_PASSWORD` - The password to log into qbittorrent with
* `INTERVAL_S` - The update interval, in seconds (defaults to 60 if unset)

## Running it

You can download one of the [releases](https://github.com/mattbun/qbtuna/releases) or it's available as a docker image at [`ghcr.io/mattbun/qbtuna`](https://github.com/mattbun/qbtuna/pkgs/container/qbtuna).

Example `docker-compose.yml`:

```yaml
services:
  qbtuna:
    container_name: qbtuna
    image: ghcr.io/mattbun/qbtuna
    environment:
      - INTERVAL_S=60
      - GLUETUN_HOST=http://gluetun:8000
      - QBITTORRENT_HOST=http://qbittorrent:8080
      - QBITTORRENT_USERNAME
      - QBITTORRENT_PASSWORD
```
