---
title: DNS
--- 

Setting up dns using [raspberry pi zero 2w](https://www.raspberrypi.com/products/raspberry-pi-zero-2-w)

## Downloading and flashing image
- Download rpi-imager and select alpine image of 64 bit for the raspberry pi. flash the image
- Download [headless.apkovl.tar.gz](https://is.gd/apkovl_master) from repo `https://github.com/macmpi/alpine-linux-headless-bootstrap/`.
- run following for connectiong wifi at boot time
```sh
cat > wpa_supplicant.conf <<EOF
country=IN 
network={
    ssid="wifi-name"
    psk="wifi-password"
}
EOF
```
- mount the flashed sd card and put `wpa_supplicant.conf` and `headless.apkovl.tar.gz`  into root of folder

## Setting up raspberry pi
- ssh into raspberry pi using root@ip it does not contain any password
- hit `setup-alpine` there and select options required

## Setting up DNS and Unbound

### Install necessary tools and prequistes
go into root using `su -` and you will be able to install
```sh
# packages
apk add curl unbound bind-tools vim coredns
# unbound and coredns folders
mkdir -p /etc/coredns/zones /var/lib/unbound
# adblock list for coredns
curl -s https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts > /etc/coredns/adblock.hosts
# unbound root.keys and trust anchors
chown -R unbound:unbound /var/lib/unbound
unbound-anchor -a /var/lib/unbound/root.key
wget https://www.internic.net/domain/named.cache -O /etc/unbound/root.hints
```

## Unbound configuration
```sh
tee "/etc/unbound/unbound.conf" > /dev/null <<EOF
server:
    # Network
    interface: 0.0.0.0
    port: 1053
    do-ip4: yes
    do-ip6: no
    so-reuseport: yes
    do-udp: yes
    do-tcp: yes
    access-control: 0.0.0.0/0 allow

    # Performance & memory
    num-threads: 2
    msg-cache-size: 32m
    rrset-cache-size: 64m
    outgoing-range: 256
    num-queries-per-thread: 128
    so-rcvbuf: 2m
    so-sndbuf: 2m
    so-reuseport: yes

    # Cache tuning
    cache-min-ttl: 3600
    cache-max-ttl: 86400

    # DNSSEC
    auto-trust-anchor-file: "/var/lib/unbound/root.key"
    val-clean-additional: yes
    val-log-level: 1
    harden-dnssec-stripped: yes
    harden-glue: yes
    harden-short-bufsize: yes
    harden-large-queries: yes
    harden-below-nxdomain: yes
    harden-referral-path: yes

    # Files
    root-hints: "/etc/unbound/root.hints"
    logfile: "/var/log/unbound.log"
    log-queries: no
    log-replies: no
    verbosity: 1

    # Security
    hide-identity: yes
    hide-version: yes
    unwanted-reply-threshold: 10000
    prefetch: yes
    prefetch-key: yes
    qname-minimisation: yes
EOF
```

## Coredns configuration
```sh
# zone files
tee "/etc/coredns/zones/db.192.168.1" > /dev/null <<'EOF'
$ORIGIN 1.168.192.in-addr.arpa.
$TTL 3600
@   IN  SOA ns1.jshk.dev. admin.jshk.dev. (
        2025110201 ; serial
        3600       ; refresh
        1800       ; retry
        1209600    ; expire
        86400 )    ; minimum
    IN  NS  ns1.jshk.dev.

249  IN  PTR pi.jshk.dev.
250  IN  PTR router.jshk.dev.
250  IN  PTR orangepi.jshk.dev
EOF

tee "/etc/coredns/zones/db.jshk.dev" > /dev/null <<'EOF'
$ORIGIN jshk.dev.
@   3600 IN SOA ns1.jshk.dev. admin.jshk.dev. (
            2025110201 ; serial
            3600       ; refresh
            1800       ; retry
            1209600    ; expire
            86400 )    ; minimum
         IN  NS  ns1.jshk.dev.

pi       IN A 192.168.1.249 ; dnspi domain which is hosting coredns and unbound
postgres IN A 192.168.1.250 ; postgres database
valkey   IN A 192.168.1.250 ; valkey database
torrent  IN A 192.168.1.250 ; web torrent version of qbittorent
minio    IN A 192.168.1.250 ; minio service for block storage
router   IN A 192.168.1.250 ; router domain for main router
registry IN A 192.168.1.250 ; zot registry
bookmark IN A 192.168.1.250 ; self hosted bookmarks manager
ci	     IN A 192.168.1.250 ; ci/cd pipeline
wiki     IN CNAME jshk00.github.io. ;self hosted wiki
EOF

# Corefile
tee "/etc/coredns/Corefile" > /dev/null <<EOF
jshk.dev {
    file /etc/coredns/zones/db.jshk.dev
}

1.168.192.in-addr.arpa {
    file /etc/coredns/zones/db.192.168.1
}

.:53 {
    log
    errors

    hosts /etc/coredns/adblock.hosts {
       fallthrough
    }

    forward . 127.0.0.1:1053
}
EOF
```

# Openrc system 
```sh
rc-update add unbound default
rc-service unbound start
rc-update add coredns default
rc-service coredns start
```
commands for openrc
```sh
rc-service <service> start
rc-service <service> stop
rc-service <service> restart
# enable service at boot
rc-update add <service> default
# disable service at boot
rc-update del <service> default
rc-service <service> status
```

## Finally some networking config
```sh
tee "/etc/resolv.conf" > /dev/null <<EOF
search jshk.dev
namerserver 127.0.0.1
EOF

tee "/etc/network/interfaces" > /dev/null <<EOF
auto lo
iface lo inet loopback

auto eth0
iface eth0 inet static
    address 192.168.1.249
    netmask 255.255.255.0
    gateway 192.168.1.1
    dns-nameservers 127.0.0.1

auto wlan0
iface wlan0 inet dhcp
EOF
```

## Unreachable node
keep wifi also on during installation and configure using dhcp. sometimes due to power or usb hub timing issue LAN may not comeup so we can manually restart it.
```sh
cat >/etc/local.d/eth0.start <<'EOF'
#!/bin/sh

LOGFILE="/var/log/eth0up.log"
log() {
    printf '%s %s\n' "$(date '+%Y-%m-%d %H:%M:%S')" "$*" >> "$LOGFILE"
}
log "Script started"

TIMEOUT=300
INTERVAL=2
ELAPSED=0

ifup wlan0 >/dev/null 2>&1

while [ "$ELAPSED" -lt "$TIMEOUT" ]; do
    log "Checking eth0 (elapsed ${ELAPSED}s)"
    ip link set eth0 up
    
    if ip link show eth0 | grep -q "LOWER_UP"; then
        log "Ethernet detected"
        ifup eth0
        ifdown wlan0
        log "Switched to Ethernet"
        exit 0
    fi
    
    sleep "$INTERVAL"
    ELAPSED=$((ELAPSED + INTERVAL))
done

log "Timed out waiting for Ethernet. Staying on Wi-Fi."
ifup wlan0
exit 0
EOF

chmod +x /etc/local.d/eth0.start
rc-update add local default
```

