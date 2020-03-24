package addr

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"sync"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	loggerAVA "github.com/ver13/ava/pkg/common/logger"
	errorAddrUtilsAVA "github.com/ver13/ava/pkg/common/utils/addr/error"
)

type addr struct {
}

var (
	privateBlocks []*net.IPNet
)

var (
	l *addr

	once sync.Once
)

func init() {
	once.Do(func() {
		for _, b := range []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16", "100.64.0.0/10", "fd00::/8"} {
			if _, block, err := net.ParseCIDR(b); err == nil {
				privateBlocks = append(privateBlocks, block)
			}
		}

		l = &addr{}
	})
}

func GetInstance() *addr {
	return l
}

func (a *addr) IsPrivateIP(ipAddr string) bool {
	ip := net.ParseIP(ipAddr)
	for _, priv := range privateBlocks {
		if priv.Contains(ip) {
			return true
		}
	}
	return false
}
func IsPrivateIP(ipAddr string) bool {
	return GetInstance().IsPrivateIP(ipAddr)
}

// Extract returns a real ip
func (a *addr) Extract(addr string) (string, *errorAVA.Error) {
	// if addr specified then its returned
	if len(addr) > 0 && (addr != "0.0.0.0" && addr != "[::]" && addr != "::") {
		return addr, nil
	}

	ifaces, err := net.Interfaces()
	if err != nil {
		return "", errorAddrUtilsAVA.FailedGetInterfaces(err, fmt.Sprintf("Failed to get interfaces! Err: %v", err))
	}

	// nolint:prealloc
	var addrs []net.Addr
	var loAddrs []net.Addr
	for _, iface := range ifaces {
		ifaceAddrs, err := iface.Addrs()
		if err != nil {
			// ignore error, interface can dissapear from system
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			loAddrs = append(loAddrs, ifaceAddrs...)
			continue
		}
		addrs = append(addrs, ifaceAddrs...)
	}
	addrs = append(addrs, loAddrs...)

	var ipAddr []byte
	var publicIP []byte

	for _, rawAddr := range addrs {
		var ip net.IP
		switch addr := rawAddr.(type) {
		case *net.IPAddr:
			ip = addr.IP
		case *net.IPNet:
			ip = addr.IP
		default:
			continue
		}

		if !a.IsPrivateIP(ip.String()) {
			publicIP = ip
			continue
		}

		ipAddr = ip
		break
	}

	// return private ip
	if ipAddr != nil {
		return net.IP(ipAddr).String(), nil
	}

	// return public or virtual ip
	if publicIP != nil {
		return net.IP(publicIP).String(), nil
	}

	return "", errorAddrUtilsAVA.IPNotFount(nil, fmt.Sprintf("No IP address found, and explicit IP not provided. %s", addr))
}
func Extract(addr string) (string, *errorAVA.Error) {
	return GetInstance().Extract(addr)
}

// IPs returns all known ips
func (a *addr) IPs() []string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}

	var ipAddrs []string

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil {
				continue
			}

			// dont skip ipv6 addrs
			/*
				ip = ip.To4()
				if ip == nil {
					continue
				}
			*/

			ipAddrs = append(ipAddrs, ip.String())
		}
	}

	return ipAddrs
}
func IPs() []string {
	return GetInstance().IPs()
}

// ResolveIPFromHostsFile reads the final IP address of the /etc/hosts file. Works for docker, typically at least...
func (a *addr) ResolveIPFromHostsFile() (string, *errorAVA.Error) {
	data, err := ioutil.ReadFile("/etc/hosts")
	if err != nil {
		loggerAVA.Errorf("Problem reading /etc/hosts: %v", err.Error())
		return "", errorAddrUtilsAVA.ReadingEtcHostsWrong(err, fmt.Sprintf("Problem reading /etc/hosts."))
	}

	lines := strings.Split(string(data), "\n")

	// Get last line
	line := lines[len(lines)-1]

	if len(line) < 2 {
		line = lines[len(lines)-2]
	}

	parts := strings.Split(line, "\t")
	return parts[0], nil
}
func ResolveIPFromHostsFile() (string, *errorAVA.Error) {
	return GetInstance().ResolveIPFromHostsFile()
}

// GetIP returns the first non-loopback IP address
func (a *addr) GetIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "error"
	}

	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"

	panic("Unable to determine local IP address (non loopback). Exiting.")
}
func GetIP() string {
	return GetInstance().GetIP()
}

// GetIPWithPrefix returns the first non-loopback IP starting with the supplied prefix.
func (a *addr) GetIPWithPrefix(prefix string) string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "error"
	}

	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && strings.HasPrefix(ipnet.IP.String(), prefix) {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"

	panic("Unable to determine local IP address (non loopback). Exiting.")
}
func GetIPWithPrefix(prefix string) string {
	return GetInstance().GetIPWithPrefix(prefix)
}
