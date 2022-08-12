package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net"
	"regexp"
)

func IsIP(ipv4 string) bool {
	if ip := net.ParseIP(ipv4); ip == nil {
		return false
	}
	return true
}

func IsCIDR(cidr string) bool {
	p := regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$|^s*((([0-9A-Fa-f]{1,4}:){7}(:|([0-9A-Fa-f]{1,4})))|(([0-9A-Fa-f]{1,4}:){6}:([0-9A-Fa-f]{1,4})?)|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){0,1}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){0,2}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){0,3}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){0,4}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){0,5}):([0-9A-Fa-f]{1,4})?))|(:(:|((:[0-9A-Fa-f]{1,4}){1,7}))))(%.+)?s*/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`)
	return p.MatchString(cidr)
}

func GetCookieFromGin(ctx *gin.Context, key string) (value string) {
	value, _ = ctx.Cookie(key)
	return
}

func NewMd5(str ...string) string {
	h := md5.New()
	for _, v := range str {
		h.Write([]byte(v))
	}
	return hex.EncodeToString(h.Sum(nil))
}

func Base64Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}
