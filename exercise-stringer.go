package main

import "fmt"

type IPAddr [4]byte

func (addr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", addr[0], addr[1], addr[2], addr[3])
}

func main() {
	addrs := map[string](IPAddr){
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		c := &a
		var b IPAddr = *c
		fmt.Printf("%v: %v\n", n, a)
		fmt.Printf("%v: %v\n", n, &a)
		fmt.Printf("%v: %v\n", n, b)
	}
}