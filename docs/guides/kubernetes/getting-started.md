---
title: Getting started
---

Kubernetes is the go-to platform for managing containerized applications at scale. However, for beginners, it can feel overwhelming.
In this series, we’ll walk through how to install Kubernetes locally, deploy some applications, and explain how Kubernetes works. We assume you have some experience with containers, but no prior Kubernetes knowledge is required.
We'll start by setting up our system and installing Kubernetes. While the upstream version of Kubernetes can be complex and difficult to install, there are lightweight Kubernetes distributions available that can be set up with just a single command.
Don't worry — in upcoming articles, we’ll also dive into what Kubernetes is and what it actually does.

## Installing Kind
Kind is a tool for running local Kubernetes clusters using Docker container “nodes”. Kind was primarily designed for testing Kubernetes itself, but may be used for local development or CI. Kind is what we refer as easy to install local kubernetes. You need [docker](https://docs.docker.com/engine/install/), [podman](https://podman.io/docs/installation) or [nerdctl](https://github.com/containerd/nerdctl) and [Go](https://go.dev/doc/install) 1.17+ installed on system.

- Go to github [releases](https://github.com/kubernetes-sigs/kind/releases) of kind
- Select appropriate version for your system generally on normal devices it's amd64, for Raspberry pi it arm64 and for MacOS it's darwin-arm64(only for M seriers chips)
- Download the binary, make it executable and move it in `/usr/local/bin`
- You can also run below commands for quick installation. copy it into your terminal and kind will be installed
```sh
curl -sSL https://raw.githubusercontent.com/ark-j/dotfiles/refs/heads/main/.config/scripts/generic/kind.sh | bash
```
- Try running `kind version` for checking successful installation

## Creating kind cluster
Before creation of cluster kubectl needed to be installed on system so that you can run commands necessary. To install kubectl follow official [guide](https://kubernetes.io/docs/tasks/tools/)

To create kind cluster run command `kind create cluster` which will create cluster with name kind and will take some time. After completion you may get output like following
<div class="termy">

```console
$ kind create cluster
Creating cluster "kind" ...
 ✓ Ensuring node image (kindest/node:v1.32.2) 🖼
 ✓ Preparing nodes 📦
 ✓ Writing configuration 📜
 ✓ Starting control-plane 🕹️
 ✓ Installing CNI 🔌
 ✓ Installing StorageClass 💾
Set kubectl context to "kind-kind"
You can now use your cluster with:

kubectl cluster-info --context kind-kind

Have a nice day! 👋
```

</div>

Verify the cluster info with `kubectl cluster-info --context kind-kind`. You will below output. Your urls may look different

<div class="termy">

```console
kubectl cluster-info --context kind-kind

Kubernetes control plane is running at https://127.0.0.1:35135
CoreDNS is running at https://127.0.0.1:35135/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
```

</div>

!!! note
    For our convinience we will set shortcut to `kubectl` as `kc`. Open `.zshrc` or `.bashrc` in your favourite editor and type `alias kc="kubectl"`
