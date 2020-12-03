package gridutils

import (
	"net"
	"os"

	"github.com/huin/goupnp/dcps/internetgateway1"
)

func getInternalIP() string {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	return ""
}

func Forward(port uint16, proto string, desc string) {
	dcps, _, _ := internetgateway1.NewWANIPConnection1Clients()
	err := dcps[0].AddPortMapping("", port, proto, port, getInternalIP(), true, "GRID:"+desc, 0)
	if err != nil {
		println(err)
	}
}

func Unforward(port uint16, proto string) {
	dcps, _, _ := internetgateway1.NewWANIPConnection1Clients()
	err := dcps[0].DeletePortMapping("", port, proto)
	if err != nil {
		println(err)
	}
}
