package main

import (
    "time"
    "net"
)

type Interface interface {
    // These methods manage a mapping between a port on the local
    // machine to a port that can be connected to from the internet.
    //
    // protocol is "UDP" or "TCP". Some implementations allow setting
    // a display name for the mapping. The mapping may be removed by
    // the gateway when its lifetime ends.
    AddMapping(protocol string, extport, intport int, name string, lifetime time.Duration) error
    DeleteMapping(protocol string, extport, intport int) error

    // This method should return the external (Internet-facing)
    // address of the gateway device.
    ExternalIP() (net.IP, error)

    // Should return name of the method. This is used for logging.
    String() string
}