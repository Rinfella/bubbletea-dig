# IP Lookup

## This is a basic TUI tool that uses the `dig` command under the hood

## to lookup IPv4 addresses for domain names.

### Prerequisites

You must have the `dig` command installed in your system.
It is usually bundled with other DNS utilities.
It may come under different package names based on your distro or package managers.

#### For Ubuntu

```bash
sudo apt install dnsutils -y
```

#### For CentOs

```bash
sudo dnf install bind-utils -y
```

#### For Arch

```bash
sudo pacman -S bind
```
