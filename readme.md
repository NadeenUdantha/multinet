# MultiNet

**MultiNet** is a Go application that allows you to use multiple internet connections simultaneously. This can be used to combine the bandwidths of different connections.

> **Warning**: Things might break. Use this application with care.

## Features

- Increase download speeds by connecting through multiple internet connections.
- Path selection algorithms: `weighted-lru`, `roundrobin`, `random`, `hash`
- Ability to route traffic through network adapters or via another proxy server.

## How to Use

1. Clone the repository:

   ```bash
   git clone https://github.com/NadeenUdantha/multinet.git
   cd multinet/multinet
   ```

2. Edit the configuration file (`config.yaml`) to set up your internet connections.

3. Build and run the application with your configuration file:

   ```bash
   go run main.go -cfg config.yaml
   ```

4. It will start a SOCKS5 proxy server at the address specified in your configuration. Point your applications to use this proxy. You can also use `tun2socks` for VPN-like functionality.

### Configuration

You configure MultiNet using a YAML file. Here’s an example configuration:

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

For questions or feedback, contact me: [NadeenUdantha](https://github.com/NadeenUdantha) me@nadeen.lk
