package lsblk

import (
    "encoding/json"
    "fmt"
    "log"
    "os/exec"
)

// Use a json.RawMessage for Children, then unmarshal that separately;
// Maybe also define a DeviceInfo struct, then use it to avoid duplication;
type Blkid struct {
    Blockdevices []struct {
        Name        string `json:"name"`
        Majmin      string `json:"maj:min"`
        Size        string `json:"size"`
        Type        string `json:"type"`
        Mountpoints []string `json:"mountpoints"`
        Removable   bool `json:"rm"`
        ReadOnly    bool `json:"ro"`
        Children []struct {
            Name        string `json:"name"`
            Majmin      string `json:"maj:min"`
            Size        string `json:"size"`
            Type        string `json:"type"`
            Mountpoints []string `json:"mountpoints"`
            Removable   bool `json:"rm"`
            ReadOnly    bool `json:"ro"`
            Children []struct {
                Name        string `json:"name"`
                Majmin      string `json:"maj:min"`
                Size        string `json:"size"`
                Type        string `json:"type"`
                Mountpoints []string `json:"mountpoints"`
                Removable   bool `json:"rm"`
                ReadOnly    bool `json:"ro"`
            }
        } `json:"children"`
    } `json:"blockdevices"`
}

func GetDeviceInfo() {
    out, err := exec.Command("lsblk", "--json").Output()
    if err != nil {
        log.Fatal(err)
    }

    var devs Blkid
    err = json.Unmarshal(out, &devs)
    if err != nil {
        log.Fatal(err)
    }

    for _, v := range devs.Blockdevices {
        fmt.Printf("%v", v.Name)
        if v.Mountpoints != nil {
            fmt.Printf("  %v", v.Mountpoints[0])
        }
        fmt.Printf("\n")
        for _, c := range v.Children {
            fmt.Printf("  %v", c.Name)
            if c.Mountpoints != nil {
                fmt.Printf("  %v", c.Mountpoints[0])
            }
            fmt.Printf("\n")
            for _, d := range c.Children {
                fmt.Printf("    %v", d.Name)
                if d.Mountpoints != nil {
                    fmt.Printf("  %v", d.Mountpoints[0])
                }
                fmt.Printf("\n")
            }
        }
    }
}
