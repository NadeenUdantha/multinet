# MultiNet

**MultiNet** is a Go application that allows you to use multiple internet connections simultaneously. This can be used to combine the bandwidths of different connections.

## Features

- Increase download speeds by connecting through multiple internet connections.
- Path selection algorithms: `weighted-lru`, `roundrobin`, `random`, `hash`
- Ability to route traffic through network adapters or through another proxy server.

## How It Works

MultiNet operates by routing traffic through multiple network paths selected using an algorithm. When multiple network connections are available, it can split data transfer across them, effectively increasing total bandwidth by using the aggregate speeds of each connection.

This method works best for applications that establish multiple parallel connections. However, the impact is minimal for single-connection downloads, which remain limited to a single network path. MultiNet can be used alongside optimized downloaders, such as IDM, to maximize download speeds. This works for uploads too. You must have a computer with enough resources to handle these speeds (disk write speed, etc.)

### Practical Application

I made this for game downloading: when downloading a 20GB Valorant update using the Riot Client, it creates 4 connections to their servers. When using this application with tun2socks, these 4 connections go through different network adapters and different internet connections. My test setup had 3 internet connections with 3 ISPs: ~30MB/s fiber and 2x ~4MB/s USB tethering, giving ~38MB/s download speed, resulting in a ~27% speed increase.

## How to Use

1. Go to the [releases](https://github.com/NadeenUdantha/multinet/releases) page and download the latest correct binary for your platform/OS.

2. Create a configuration file (`config.yaml`) to set up your internet connection details. See below for the configuration file structure.

3. Run the application with your configuration file.

   ```bash
   ./multinet -cfg config.yaml
   ```

4. It will start a SOCKS5 proxy server at the address specified in your configuration. Point your applications to use this proxy server. You can also use `tun2socks` for VPN-like functionality.

### Configuration

You can configure MultiNet using a YAML file. Here’s an example configuration:

```yaml
# The algorithm used for selecting the network path.
# One of: weighted-lru, roundrobin, random, hash
algorithm: weighted-lru

# The address where the SOCKS5 server will listen for incoming connections.
# This should be in the format of IP:PORT.
listen: 192.168.1.69:1080

# If set to true, the port part of addr will be hashed when generating the address hash key.
hashport: true

# A list of network paths to use for routing traffic.
# Each path can be of type 'direct' or 'proxy'.
paths:
  - type: direct # Direct connection through the first internet connection.
    addr: 192.168.1.69 # IP address of Network Adapter #1.

  - type: direct # Direct connection through the second internet connection.
    addr: 192.168.57.206 # IP address of Network Adapter #2.

  - type: direct # Direct connection through the third internet connection.
    addr: 192.168.137.211 # IP address of Network Adapter #3.

  - type: proxy # Connection through the first proxy server.
    addr: socks5://127.0.0.1:1081 # SOCKS5 proxy address.

  - type: proxy # Connection through the second proxy server.
    addr: socks5://127.0.0.1:1082 # SOCKS5 proxy address.
```

### TODO
- Add tests
- Add default auto config
- Create a simple UI
- Add integrated VPN
- More path selection algorithms
- ???

[MultiNet](https://github.com/NadeenUdantha/multinet) by [NadeenUdantha](https://nadeen.lk) is licensed under [CC BY-NC-ND 4.0](https://creativecommons.org/licenses/by-nc-nd/4.0/) 💀

For questions or feedback, contact me: [NadeenUdantha](https://github.com/NadeenUdantha) me@nadeen.lk
