package main

import (
  "net"
)

/**
Get the local ip
 */
func Get() (err error, result string) {
  var (
    addresses []net.Addr
  )
  if addresses, err = net.InterfaceAddrs(); err != nil {
    return
  }
  for _, address := range addresses {
    if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
      if ipnet.IP.To4() != nil {
        result = ipnet.IP.String()
        return
      }
    }
  }
  return
}
