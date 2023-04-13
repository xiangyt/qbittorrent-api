package magnet

import (
	"encoding/base32"
	"encoding/hex"
	"strings"
)

const Prefix = "magnet:?xt=urn:btih:"

func Base32ToHex(src string) string {
	switch len(src) {
	case 52, 32:
		if strings.HasPrefix(src, Prefix) {
			return Prefix + base32ToHex(src[20:])
		} else {
			return base32ToHex(src)
		}
	default:
		return ""
	}

}

func base32ToHex(src string) string {
	bytes, err := base32.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func HexToBase32(src string) string {
	switch len(src) {
	case 60, 40:
		if strings.HasPrefix(src, Prefix) {
			return Prefix + hexToBase32(src[20:])
		} else {
			return hexToBase32(src)
		}
	default:
		return ""
	}

}

func hexToBase32(src string) string {
	bytes, err := hex.DecodeString(src)
	if err != nil {
		return ""
	}
	return base32.StdEncoding.EncodeToString(bytes)
}

func GetHash(src string) string {
	switch len(src) {
	case 32:
		return base32ToHex(src)
	case 40:
		return src
	case 52:
		return base32ToHex(strings.TrimPrefix(src, Prefix))
	case 60:
		return strings.TrimPrefix(src, Prefix)
	default:
		return ""
	}
}
