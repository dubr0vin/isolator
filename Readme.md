# Isolator
Simple tool for isolating programms

## Installing

Download binary from [releases](https://github.com/dubr0vin/isolator/releases/) page.

## Basic usage

```shell
isolator /path/to/program arg1 arg2 ...
```

You, probably, need to select custom path to chroot or disable it:

```shell
isolator -chroot-dir=/path/to/chroot /path/to/program arg1 arg2 ...
isolator -disable-chroot /path/to/program arg1 arg2 ...
```
