package gridutils

import (
	"net"
	"os"

	"github.com/huin/goupnp/dcps/internetgateway2"
)

var dcps []*internetgateway2.WANIPConnection1

func init() {
	dcps, _, _ = internetgateway2.NewWANIPConnection1Clients()
}

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
func GetExternalIP() (string, error) {
	ip, err := dcps[0].GetExternalIPAddress()
	if err != nil {
		println(err.Error())
		return "", err
	}
	return ip, nil
}

//Forward forwards a proto on port with desc
func Forward(port uint16, proto string, desc string) error {
	err := dcps[0].AddPortMapping("", port, proto, port, GetInternalIP(), true, "GRID:"+desc, 0)
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

//Unforward removes a proto on port
func Unforward(port uint16, proto string) error {
	err := dcps[0].DeletePortMapping("", port, proto)
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}
