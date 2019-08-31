# quic-socks
quic-socks implements socks5 server using custom protocol in the back end, due to the use of QUIC , 2x faster than shadowsocks, and safer.
## Features
* implements socks5 server in the front end for less RTT
* using custom protocol in the back end(client <-> server), only need 1 RTT
* the experience is still good in the case of weak networks
* the experience will not deteriorate in the case of mobile networks
* client <-> server using TLS 1.3(QUIC), less RTT

## Protocol
password + type + host + port\
see protocol.go
