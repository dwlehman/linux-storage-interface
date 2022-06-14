package device

type Device struct {
    id          int
    name        string
    path        string
    type        string
    size        uint64
    format      string  // needs a type
    parents     []int
    children    []int
}

type DeviceOps interface {
    Activate()      bool
    Deactivate()    bool

    Create()        bool
    Destroy()       bool
}

