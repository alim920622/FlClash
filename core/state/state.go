//go:build android && cgo

package state

var DefaultIpv4Address = "192.168.43.231/30"
var DefaultDnsAddress = "192.168.43.232"
var DefaultIpv6Address = "fdfe:dcba:9876::1/126"

type AndroidVpnOptions struct {
	Enable           bool           `json:"enable"`
	Port             int            `json:"port"`
	AccessControl    *AccessControl `json:"accessControl"`
	AllowBypass      bool           `json:"allowBypass"`
	SystemProxy      bool           `json:"systemProxy"`
	BypassDomain     []string       `json:"bypassDomain"`
	RouteAddress     []string       `json:"routeAddress"`
	Ipv4Address      string         `json:"ipv4Address"`
	Ipv6Address      string         `json:"ipv6Address"`
	DnsServerAddress string         `json:"dnsServerAddress"`
}

type AccessControl struct {
	Mode              string   `json:"mode"`
	AcceptList        []string `json:"acceptList"`
	RejectList        []string `json:"rejectList"`
	IsFilterSystemApp bool     `json:"isFilterSystemApp"`
}

type AndroidVpnRawOptions struct {
	Enable        bool           `json:"enable"`
	AccessControl *AccessControl `json:"accessControl"`
	AllowBypass   bool           `json:"allowBypass"`
	SystemProxy   bool           `json:"systemProxy"`
	RouteAddress  []string       `json:"routeAddress"`
	Ipv6          bool           `json:"ipv6"`
	BypassDomain  []string       `json:"bypassDomain"`
}

type State struct {
	AndroidVpnRawOptions
	CurrentProfileName string `json:"currentProfileName"`
}

var CurrentState = &State{}

func GetIpv6Address() string {
	if CurrentState.Ipv6 {
		return DefaultIpv6Address
	} else {
		return ""
	}
}

func GetDnsServerAddress() string {
	return DefaultDnsAddress
}
