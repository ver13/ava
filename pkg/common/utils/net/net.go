package net

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"

	errorGmf "github.com/ValentinEncinasRojas/ava/errors"
	errorNetUtilsGmf "github.com/ValentinEncinasRojas/ava/pkg/common/utils/net/error"
)

type netGmf struct {
}

var (
	f *netGmf

	once sync.Once
)

func init() {
	once.Do(func() {
		f = &netGmf{}
	})
}

func GetInstance() *netGmf {
	return f
}

// HostPort format addr and port suitable for dial
func (n *netGmf) HostPort(addr string, port interface{}) string {
	host := addr
	if strings.Count(addr, ":") > 0 {
		host = fmt.Sprintf("[%s]", addr)
	}
	// when port is blank or 0, host is a queue name
	if v, ok := port.(string); ok && v == "" {
		return host
	} else if v, ok := port.(int); ok && v == 0 && net.ParseIP(host) == nil {
		return host
	}

	return fmt.Sprintf("%s:%v", host, port)
}
func HostPort(addr string, port interface{}) string {
	return GetInstance().HostPort(addr, port)
}

// Listen takes addr:portmin-portmax and binds to the first available port
// Example: Listen("localhost:5000-6000", fn)
func (n *netGmf) Listen(addr string, fn func(string) (net.Listener, errorGmf.ErrorGmfI)) (net.Listener, errorGmf.ErrorGmfI) {

	if strings.Count(addr, ":") == 1 && strings.Count(addr, "-") == 0 {
		return fn(addr)
	}

	// host:port || host:min-max
	host, ports, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, errorNetUtilsGmf.HostPortWrong(err, fmt.Sprintf("%v.", addr))
	}

	// try to extract port range
	prange := strings.Split(ports, "-")

	// single port
	if len(prange) < 2 {
		return fn(addr)
	}

	// we have a port range

	// extract min port
	min, err := strconv.Atoi(prange[0])
	if err != nil {
		return nil, errorNetUtilsGmf.PortWrong(err, fmt.Sprintf("Unable to extract port range. %s", prange[0]))
	}

	// extract max port
	max, err := strconv.Atoi(prange[1])
	if err != nil {
		return nil, errorNetUtilsGmf.PortWrong(err, fmt.Sprintf("Unable to extract port range. %s", prange[1]))
	}

	// range the ports
	for port := min; port <= max; port++ {
		// try bind to host:port
		ln, err := fn(HostPort(host, port))
		if err == nil {
			return ln, nil
		}

		// hit max port
		if port == max {
			return nil, err
		}
	}

	// why are we here?
	return nil, errorNetUtilsGmf.AddrWrong(err, fmt.Sprintf("Unable to bind to %s", addr))
}
func Listen(addr string, fn func(string) (net.Listener, errorGmf.ErrorGmfI)) (net.Listener, errorGmf.ErrorGmfI) {
	return GetInstance().Listen(addr, fn)
}
