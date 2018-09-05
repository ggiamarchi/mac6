package mac6

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func isHexChar(char byte) bool {
	switch char {
	case
		byte('a'), byte('b'), byte('c'), byte('d'),
		byte('e'), byte('f'), byte('0'), byte('1'),
		byte('2'), byte('3'), byte('4'), byte('5'),
		byte('6'), byte('7'), byte('8'), byte('9'):
		return true
	}
	return false
}

// normalizeMACAddress takes the input MAC address and remove every non hexa symbol
// and lowercase everything else
func normalizeMACAddress(mac string) string {
	var buffer bytes.Buffer

	macArray := strings.Split(strings.ToLower(mac), ":")

	for i := 0; i < len(macArray); i++ {
		m := macArray[i]
		if len(m) == 1 {
			buffer.WriteByte(byte('0'))
		}
		for j := 0; j < len(m); j++ {
			if isHexChar(m[j]) {
				buffer.WriteByte(m[j])
			}
		}
	}
	return buffer.String()
}

// This could also return []bool
func asBits(val uint64) []uint64 {
	bits := []uint64{}
	for i := 0; i < 8; i++ {
		bits = append([]uint64{val & 0x1}, bits...)
		// or
		// bits = append(bits, val & 0x1)
		// depending on the order you want
		val = val >> 1
	}
	return bits
}

func parseBinToHex(s string) string {
	ui, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		return "error"
	}

	return fmt.Sprintf("%x", ui)
}

func invertBit7(hexOctet string) (string, error) {
	i, err := strconv.ParseUint(hexOctet, 16, 8)
	if err != nil {
		return "", err
	}
	s := ""
	for k, b := range asBits(i) {
		if k == 6 {
			s = fmt.Sprintf("%s%d", s, (b+1)%2)
		} else {
			s = fmt.Sprintf("%s%d", s, b)
		}
	}

	return parseBinToHex(s), nil
}

func Mac6(mac string) string {

	nMac := normalizeMACAddress(mac)

	s, err := invertBit7(nMac[:2])
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("fe80::%s%s:%sff:fe%s:%s%s", s, nMac[2:4], nMac[4:6], nMac[6:8], nMac[8:10], nMac[10:])
}
