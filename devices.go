package main

import (
    "fmt"
    "devices/lsblk"
)

func main() {
    fmt.Println("Hi there!")
    lsblk.GetDeviceInfo()
}

