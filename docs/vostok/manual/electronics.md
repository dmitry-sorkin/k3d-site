---

authors:
  - sorkin
icon: material/source-commit-local
title: VOSTOK - Электроника
description: Сборка и настройка электроники
---

# Настройка коммутационных плат

sudo apt install ifupdown
sudo nano /etc/network/interfaces
```
auto can0
allow-hotplug can0
iface can0 inet manual
    pre-up /sbin/ip link set $IFACE type can bitrate 1000000
    pre-up /sbin/ip link set $IFACE txqueuelen 128
    up /sbin/ip link set $IFACE up
    down /sbin/ip link set $IFACE down
```
sudo reboot