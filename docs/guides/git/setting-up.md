---
title: getting started
---

# Setting up git for first time

## Setting up git config
---
- set user name `git config --global user.name "your_user_name"`
- set email `git config --global user.email "your_email"`

## Using HTTP
---
- Use subsequent command to set git to always use https url if if you clone using ssh. `git config --global url."https://github.com".insteadOf "git@github.com:"`
- install library `libsecret` the library name may change depdending upon linux distribution
- this will save credentials to gnome-keyring or libsecret `git config --global credential.helper /usr/lib/git-core/git-credential-libsecret`
- if you can't find this path `/usr/lib/git-core/git-credential-libsecret` then you can use `locate -b git-credential-libsecret`
- credentials store with libsecret are encrypted so they are secure

!!! tip "optional"
    you can try https://github.com/hickford/git-credential-oauth this library to authenticate using oauth installation instruction are given in repo.

## Using SSH
---
- Use subsequent command to set git to always use ssh url if if you clone using https. `git config --global url."ssh://git@github.com".insteadOf "https://github.com"`
- Install `openssh` package in your system
- Enable service to start on startup for ssh daemon and keygen using `sudo systemctl enable --now sshd sshdgenkeys`
- Run this command to generate keys `ssh-keygen -t ed25519 -C "your_email"`
- Provide filename to save generally i recommend to use githubusername-git
- There will be 2 files generated in current directory one public key with `.pub` extension and other without extension that's your private key
- Now go to github settings &rarr; SSH and GPG Keys &rarr; click on New SSH Key &rarr; put in title, and content of generated key with extension .pub
- done
