package device

import (
    "fmt"
    "linux-storage-interface/lsblk"
)

type Device struct {
    Name        string
    Path        string
    Type        string
    Size        string

    FormatType  string
    Mountpoint  string

    Parents     []Device
    Children    []Device
}

func NewDevice(bdev lsblk.BlkidDevice) *Device {
    device := Device{Name: bdev.Name, Type: bdev.Type, Size: bdev.Size}
    for _, bchild := range bdev.Children {
        child := *NewDevice(bchild)
        device.Children = append(device.Children, child)
        child.Parents = append(child.Parents, device)
    }
    return &device
}

func PrintDevice(device Device) {
    fmt.Printf("%v\n", device.Name)
    for _, child := range device.Children {
        PrintDevice(child)
    }
}

type DeviceOps interface {
    Activate()      bool
    Deactivate()    bool

    Create()        bool
    Destroy()       bool

    Wipe()          bool
    Format()        bool

    Mount()         bool
    Unmount()       bool
}

