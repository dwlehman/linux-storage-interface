package main

import (
    "devices/lsblk"
)

func main() {
    devices := lsblk.GetDeviceInfo()
    lsblk.PrintBlkidDevices(devices, "")
}
