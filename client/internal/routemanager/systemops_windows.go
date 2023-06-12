//go:build windows
// +build windows

package routemanager

import (
	"net"
	"net/netip"

	"github.com/yusufpapurcu/wmi"
)

type Win32_IP4RouteTable struct {
	Destination string
	Mask        string
	NextHop     string
}

func existsInRouteTable(prefix netip.Prefix) (bool, error) {
	var routes []Win32_IP4RouteTable
	query := "SELECT Destination, Mask, NextHop FROM Win32_IP4RouteTable"

	err := wmi.Query(query, &routes)
	if err != nil {
		return true, err
	}

	for _, route := range routes {
		ip := net.ParseIP(route.Mask)
		mask := net.IPMask(ip)
		cidr, _ := mask.Size()
		if route.Destination == prefix.Addr().String() && cidr == prefix.Bits() {
			return true, nil
		}
	}
	return false, nil
}
