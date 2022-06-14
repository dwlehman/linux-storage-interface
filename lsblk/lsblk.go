package lsblk

import (
    "encoding/json"
    "fmt"
    "log"
    "os/exec"
)

type BlkidDevice struct {
    Name        string `json:"name"`
    Majmin      string `json:"maj:min"`
    Size        string `json:"size"`
    Type        string `json:"type"`
    Mountpoints []string `json:"mountpoints"`
    Removable   bool `json:"rm"`
    ReadOnly    bool `json:"ro"`
    Children    []BlkidDevice `json:"children"`
}

type Blkid struct {
    Blockdevices []BlkidDevice `json:"blockdevices"`
}

func PrintBlkidDevices(devices []BlkidDevice, prefix string) {
    for _, device := range devices {
        fmt.Printf("%v%v: %v %v", prefix, device.Name, device.Size, device.Type)
        if len(device.Mountpoints[0]) > 0 {
            fmt.Printf(" [%v]", device.Mountpoints[0])
        }
        fmt.Printf("\n")
        PrintBlkidDevices(device.Children, prefix + string("  "))
    }
}

func GetDeviceInfo() []BlkidDevice {
    out, err := exec.Command("lsblk", "--json").Output()
    if err != nil {
        log.Fatal(err)
    }

    var blkid_data Blkid
    err = json.Unmarshal(out, &blkid_data)
    if err != nil {
        log.Fatal(err)
    }

    //PrintBlkidDevices(blkid_data.Blockdevices, string(""))
    return blkid_data.Blockdevices
}
