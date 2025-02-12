//go:build android && cgo

package state

import (
	"fmt"
	"math/rand"
	"time"
)

var DefaultIpv4Address string
var DefaultDnsAddress string
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

// generateRandomIP generates a random IP address where each octet is in the range of 10-250.
func generateRandomIP() {
	rand.Seed(time.Now().UnixNano())

	// Генерация случайных чисел для каждого октета от 10 до 250.
	octet2 := rand.Intn(241) + 10  // Генерация от 10 до 250
	octet3 := rand.Intn(241) + 10  // Генерация от 10 до 250
	octet4 := rand.Intn(241) + 10  // Генерация от 10 до 250

	// Формируем IP-адрес.
	ipAddress := fmt.Sprintf("10.%d.%d.%d", octet2, octet3, octet4)
	
	// Формируем DNS-адрес (последний октет + 1).
	dnsAddress := fmt.Sprintf("10.%d.%d.%d", octet2, octet3, octet4+1)

	// Возвращаем IP-адрес и DNS-адрес.
	DefaultIpv4Address = ipAddress + "/30"
	DefaultDnsAddress = dnsAddress
	
}

func init() {
	// Генерация случайного IP и DNS.
	generateRandomIP()
}
