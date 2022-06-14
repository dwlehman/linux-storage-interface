package main

import (
    "linux-storage-interface/lsblk"
    "linux-storage-interface/device"
)

func DevicesFromBlkid(bdevs []lsblk.BlkidDevice) []device.Device {
    devices := make([]device.Device, 0)
    for _, bdev := range bdevs {
        device := *device.NewDevice(bdev)
        devices = append(devices, device)
    }
    return devices
}

func main() {
    devices := DevicesFromBlkid(lsblk.GetDeviceInfo())
    for _, dev := range devices {
        device.PrintDevice(dev)
    }
}
