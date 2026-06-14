---
title: Networking
---
Useful linux networking tool and commands

## Interface and Routing
`ip` (replacement for ifocnfig, route, arp)

### Show interfaces 

```sh
ip link
ip link show eth0
ip -j link show eth0 # prints in json format
```
### Bring interface up/down

```sh
sudo ip link set eth0 up
sudo ip link set eth0 down
```
### IP adresses

```sh
ip addr # show all
ip addr show eth0 # show eth0
ip -4 addr # show ipv4
ip -6 addr # show ipv6
sudo ip addr add 192.168.1.100/24 dev eth0 # assign
sudo ip addr del 192.168.1.100/24 dev eth0 # delete
```
### Routing

```sh
ip route
ip route get 8.8.8.8 # route through the internet
ip route get 192.168.1.1 # router ip address route 
```

Example &rarr;

Generally the src parameter in router ip and external ip should be same in network debugging.
<div class="termy">

```sh
$ ip route
default via 192.168.1.1 dev wlan0 proto static metric 600
172.17.0.0/16 dev docker0 proto kernel scope link src 172.17.0.1 linkdown
192.168.1.0/24 dev wlan0 proto kernel scope link src 192.168.1.241 metric 600

$ ip route get 192.168.1.1
192.168.1.1 dev wlan0 src 192.168.1.241 uid 1000
    cache

$ ip route get 8.8.8.8
8.8.8.8 via 192.168.1.1 dev wlan0 src 192.168.1.241 uid 1000
    cache
```

</div>

### Neighbours (ARP)
Kernel adds the ARP(Address Resolution Protocol) basically entry of ip to mac in cache if you communicated with the node recently.
```sh
ip neigh
ip neigh show 192.168.1.249 # check for specific neighbour
```

## DNS
DNS debugging tool

### Lookup
<div class="termy">

```sh
$ dig google.com

; <<>> DiG 9.20.23 <<>> google.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 62428
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 1232
;; QUESTION SECTION:
;google.com.			IN	A

;; ANSWER SECTION:
google.com.		1655	IN	A	142.251.223.14

;; Query time: 6 msec
;; SERVER: 192.168.1.20#53(192.168.1.249) (UDP)
;; WHEN: Sun Jun 14 18:54:40 IST 2026
;; MSG SIZE  rcvd: 65
```

</div>

### Only answer
```sh
dig +short google.com
```

### Specific records
```sh
dig A google.com # IPV4
dig AAAA google.com # IPV6
dig MX gmail.com # MAIL
dig TXT google.com # VALIDATION
dig NS github.com  # Authorative Nameserver
dig CNAME www.example.com # Alias names
```

### Use Specific DNS server
```sh
dig @8.8.8.8 google.com
dig @192.168.1.20 -p 1053 google.com
```

### Trace delegation
```sh
dig +trace google.com
```

### Reverse lookup
```sh
dig -x 8.8.8.8
```

### NSLookup
Older tool
```sh
nslookup google.com
nslookup github.com 8.8.8.8

nslookup
> google.com
> set type=mx
> gmail.com
```

### Host
Very quick
```sh
host google.com
host 8.8.8.8
host -t MX gmail.com
```

## Connectivity

### `ping` ICMP echo
```sh
ping google.com
ping -c 4 google.com    # number of pings
ping -i 0.2 google.com  # interval
ping6 google.com        # IPV6
```

### traceroute
```sh
traceroute google.com    # show packaet path
traceroute -I google.com # icmp
traceroute -T google.com # tcp
```

### tracepath
```sh
sudo tracepath google.com
```

### ss `Socket Statistic`

- **State:** Indicates the current status of the socket, such as LISTEN (waiting for connections) or ESTABLISHED (active communication between systems)
- **Recv-Q / Send-Q:** Shows the amount of data queued for receiving or sending, which helps identify delays or bottlenecks in communication
- **Local Address:Port:** Displays the IP address and port number on your system where the socket is created or listening for connections
- **Peer Address:Port -** Represents the remote system’s IP address and port number connected to your machine
- **Protocol Type:** Specifies the communication protocol used by the socket, such as TCP (connection-oriented) or UDP (connectionless

```sh
ss -l    # listening ports
ss -t    # display tcp
ss -u    # display udp
ss -lt   # listening tcp
ss -ltnp # listening tcp with raw numeric ip and process names which started it
ss -tan  # established connectionless
```

## Port testing
### nc (netcat)
```sh
nc -vz google.com 443   # test tcp port
nc localhost 8080       # connect manually
nc -l 8080              # Listen
nc -l 9000 > file.txt   # transfer to file
nc host 9000 < file.txt # transfer from file
```

### telnet
```sh
telnet google.com 80 # simple tcp testing
```

## Network Scanning
```sh
nmap 192.168.1.0/24     # host discovery
nmap google.com         # common ports
nmap -p 22,80,443 host  # specific ports
nmap -sV host           # version detection
sudo nmap -O host       # os detection
sudo nmap -A host       # aggressive scan
sudo nmap -sS host      # TCP SYN
sudo nmap -sU host      # UDP ports
```

## Packet capture using tcpdump
```sh
tcpdump -D                      # view interfaces
sudo tcpdump                    # capture
sudo tcpdump -i eth0            # interface
sudo tcpdump host 8.8.8.8       # host
sudo tcpdump port 443           # port
sudo tcpdump udp port 53        # dns
sudo tcpdump -w capture.pcap    # pcap capture
tcpdump -r capture.pcap         # pcap read
```

## Bandwidth check using `iperf3`
```sh
iperf3 -s           # server
iperf3 -c SERVER_IP # client
iperf3 -R           # reverse
iperf3 -u           # udp
```

## Network manager
```sh
systemctl status NetworkManager     # network manager running
nmcli general status                # general status
nmcli dev                           # show devices
nmcli device show iface_name        # details information
nmcli con                           # show connection
nmcli con show con_name             # detailed information about connection
nmcli con show --active             # show active connection
nmcli -f IP4.DNS device show wlan0  # filter
nmcli con up con_profile            # activate the connection
nmcli con down con_profile          # deactive the connection

# add static connection profile if not present is use `con mod`
nmcli con add \
    type ethernet \
    ifname eth0 \
    con-name "Wired Connection 1" \
    ipv4.method manual \
    ipv4.addresses 192.168.1.100/24 \
    ipv4.gateway 192.168.1.1 \
    ipv4.dns "1.1.1.1 8.8.8.8"

# vlan 
nmcli con add \
    type vlan \
    con-name vlan10 \
    ifname eth0.10 \
    dev eth0 \
    id 10 \
    ipv4.method manual \
    ipv4.addresses 192.168.10.20/24 \
    ipv4.gateway 192.168.10.1 \
    ipv4.dns "1.1.1.1 8.8.8.8"
```
