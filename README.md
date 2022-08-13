# UDPFLOOD

udp amplification attack, used in server stress testing and DDOS defense testing.

## Use

### Build

```bash
git clone https://github.com/AH-dark/udpflood
cd udpflood
go build -o ../udpflood-exe
cd ..
```

### Run

```bash
./udpflood -h <hostname> -p <port> -t <threads> -s <size>
```

Four parameters are supported:

* `-h <hostname>` Set the IP or domain name of the test server. Default is `localhost`.
* `-p <port>` Set the port for sending UDP packets. Default is `8080`.
* `-t <threads>` Set the number of concurrent attacks. Default is `100`.
* `-s <size>` Set the size of packets. Default is `65507`.
