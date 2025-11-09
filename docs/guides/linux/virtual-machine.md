---
title: All about virtual machines
---
# Installing the required component
```sh
pacman -S qemu-base qemu-base virt-install libvirt dnsmasq spice qemu-ui-opengl qemu-ui-spice-core qemu-chardev-spice qemu-chardev-spice qemu-audio-spice # if spice is required
pacman -S qemu-base qemu-base virt-install libvirt dnsmasq # if vnc is desired
```

# Installing virtual machine with virt-install
```sh
sudo virt-install \
    --name={VM-NAME} \
    --memory=4096 \
    --vcpus=4 \
    --disk size=30,bus=virtio,format=qcow2,io=io_uring,cache=none,discard=unmap \
    --cdrom {full-iso-path} \
    --os-variant=debian13 \
    --network network=default,model=virtio \
    --noautoconsole \
    --autostart \
    --graphics=spice,listen=0.0.0.0 
```

# Some useful commands
```sh
sudo virsh domifaddr vm-name --source agent # get the ip addresses assigned to vm
sudo virsh destroy vm-name # removes vm
sudo virsh undefine --remove-all-storage vm-name # removes vm image and it's data
sudo virsh console vm-name # aquire console access to vm
sudo virsh domdisplay --domain vm-name # get the ports on which vnc or spice started
```

