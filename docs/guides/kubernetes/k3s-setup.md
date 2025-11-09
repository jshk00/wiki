# Setting up k3s
Below command will install k3s with network policy and flannel disabled.
Starting kube-proxy in ipvs mode for faster iteration.
Afterward apply kube-router as a cni and network proxy handler
```sh
curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="server --flannel-backend none \
        --disable-network-policy \
        --write-kubeconfig-mode 644 \
        --kube-proxy-arg proxy-mode=ipvs" sh -s -
curl -s https://raw.githubusercontent.com/cloudnativelabs/kube-router/master/daemonset/generic-kuberouter.yaml | kubectl apply -f -
```

Uninstalling k3s. This will completely cleanup the k3s from host 
```sh
sudo systemctl stop k3s || true
sudo /usr/local/bin/k3s-uninstall.sh || true
sudo ip link delete kube-bridge || true
sudo ip link delete kube-dummy-if || true
sudo rm -rf /etc/cni/net.d /var/lib/cni /etc/cni/ /etc/rancher /var/lib/kube-router
```

# Routing baremetal services using k3s and traefik

```yaml
apiVersion: discovery.k8s.io/v1
kind: EndpointSlice
metadata:
  name: torrent-host
  namespace: default
  labels:
    kubernetes.io/service-name: torrent-host
addressType: IPv4
endpoints:
  - addresses:
      - "192.168.1.250"
    conditions:
      ready: true
ports:
  - name: http
    port: 3244
    protocol: TCP
---
apiVersion: v1
kind: Service
metadata:
  name: torrent-host
spec:
  type: ""
  clusterIP: None
  ports:
    - port: 3244
      targetPort: 3244
      name: http
      protocol: TCP
---
apiVersion: traefik.io/v1alpha1
kind: IngressRoute
metadata:
  name: torrent-route
spec:
  routes:
    - match: Host(`torrent.internal.lan`)
      kind: Rule
      services:
        - name: torrent-host
          port: 8080
```
