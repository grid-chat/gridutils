package gridutils

import (
	"net"
	"os"

	"github.com/huin/goupnp/dcps/internetgateway1"
)

//GetInternalIP gets the first ipv4 non loopback ip
func GetInternalIP() string {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil && !ipv4.IsLoopback() {
			return ipv4.String()
		}
	}
	return ""
}

//GetExternalIP gets the external IP
func GetExternalIP() string {
	dcps, _, _ := internetgateway1.NewWANIPConnection1Clients()
	ip, err := dcps[0].GetExternalIPAddress()
	if err != nil {
		println(err)
		return ""
	}
	return ip
}

//Forward forwards a proto on port with desc
func Forward(port uint16, proto string, desc string) {
	dcps, _, _ := internetgateway1.NewWANIPConnection1Clients()
	err := dcps[0].AddPortMapping("", port, proto, port, GetInternalIP(), true, "GRID:"+desc, 0)
	if err != nil {
		println(err)
	}
}

//Unforward removes a proto on port
func Unforward(port uint16, proto string) {
	dcps, _, _ := internetgateway1.NewWANIPConnection1Clients()
	err := dcps[0].DeletePortMapping("", port, proto)
	if err != nil {
		println(err)
	}
}
