---
title: Setting up NFS and SMB
---

# NFS
- Installation of packages
```sh
dnf install nfs-utils # fedora/redhat
pacman -S nfs-utils # archlinux
apt install nfs-kernel-server # ubuntu/debian
```
- Enable the service using `systemctl enable --now nfs-kernel-server` 
- Create export directory 
```sh
mkdir /mnt/service-store
chown nobody:nogroup /mnt/service-store
chmod 777 /mnt/service-store
```
- Configure exports using 
```sh
sudo tee /etc/exports > /dev/null <<EOF
/mnt/service-store 192.168.1.0/24(rw,sync,no_subtree_check,no_root_squash)
EOF
sudo exportfs -rav
```
- Optimizing client and mounting it
```sh
sudo mount -t nfs -o nfsvers=4.2,rsize=1048576,wsize=1048576,async,noatime,nodiratime,tcp <server-ip>:/mnt/service-store ${HOME}/nfs-service-store
```

