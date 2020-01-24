package mac

import (
	"bytes"
	"fmt"
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

// NormalizeMACAddress takes the input MAC address and remove every non hexa symbol
// and lowercase everything else
func NormalizeMACAddress(mac string) (string, error) {
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

	if buffer.Len() != 12 {
		return "", fmt.Errorf("Invalid MAC address")
	}

	return buffer.String(), nil
}
