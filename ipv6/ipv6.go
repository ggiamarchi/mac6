package ipv6

import (
	"fmt"
	"strconv"

	"github.com/ggiamarchi/mac6/mac"
)

func intToBitArray(val uint8) []uint8 {
	bits := []uint8{}
	for i := 0; i < 8; i++ {
		bits = append([]uint8{val & 0x1}, bits...)
		val = val >> 1
	}
	return bits
}

func parseBinStringToHex(s string) string {
	ui, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		return "error"
	}

	return fmt.Sprintf("%x", ui)
}

// TODO Implementation can be done in a better way using bitwise operations
func invertBit7(hexOctet string) (string, error) {
	i, err := strconv.ParseUint(hexOctet, 16, 8)
	if err != nil {
		return "", err
	}
	s := ""
	for k, b := range intToBitArray(uint8(i)) {
		if k == 6 {
			s = fmt.Sprintf("%s%d", s, (b+1)%2)
		} else {
			s = fmt.Sprintf("%s%d", s, b)
		}
	}

	return parseBinStringToHex(s), nil
}

func Compute(macAddress string) (string, error) {

	nMac, err := mac.NormalizeMACAddress(macAddress)

	if err != nil {
		return "", err
	}

	s, err := invertBit7(nMac[:2])
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("fe80::%s%s:%sff:fe%s:%s%s", s, nMac[2:4], nMac[4:6], nMac[6:8], nMac[8:10], nMac[10:]), nil
}
