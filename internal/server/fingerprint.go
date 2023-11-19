package server

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"

	"github.com/gin-gonic/gin"
)

func getUserAgent(r *gin.Context) string {
	return string(r.Request.UserAgent())
}

func getUserIp(r *gin.Context) string {
	ip := r.Request.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = r.Request.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = r.Request.RemoteAddr
	}
	return ip
}

func getUserId(r *gin.Context) string {
	return string(r.Request.URL.Query().Get("user_id"))
}

func getOs(u string) string {
	os_ref := make(map[string]*regexp.Regexp)
	os_ref["Windows 10"] = regexp.MustCompile(`(?i)windows nt 10`)
	os_ref["Windows 8.1"] = regexp.MustCompile(`(?i)windows nt 6.3`)
	os_ref["Windows 8"] = regexp.MustCompile(`(?i)windows nt 6.2`)
	os_ref["Windows 7"] = regexp.MustCompile(`(?i)windows nt 6.1`)
	os_ref["Windows Vista"] = regexp.MustCompile(`(?i)windows nt 6.0`)
	os_ref["Windows Server 2003/XP x64"] = regexp.MustCompile(`(?i)windows nt 5.2`)
	os_ref["Windows XP"] = regexp.MustCompile(`(?i)windows nt 5.1`)
	os_ref["Windows 2000"] = regexp.MustCompile(`(?i)windows nt 5.0`)
	os_ref["Windows ME"] = regexp.MustCompile(`(?i)windows me`)
	os_ref["Windows 98"] = regexp.MustCompile(`(?i)win98`)
	os_ref["Windows 95"] = regexp.MustCompile(`(?i)win95`)
	os_ref["Windows 3.11"] = regexp.MustCompile(`(?i)win16`)
	os_ref["Mac OS X"] = regexp.MustCompile(`(?i)macintosh|mac os x`)
	os_ref["Mac OS 9"] = regexp.MustCompile(`(?i)mac_powerpc`)
	os_ref["Linux"] = regexp.MustCompile(`(?i)linux`)
	os_ref["Ubuntu"] = regexp.MustCompile(`(?i)ubuntu`)
	os_ref["iPhone"] = regexp.MustCompile(`(?i)iphone`)
	os_ref["iPod"] = regexp.MustCompile(`(?i)ipod`)
	os_ref["iPad"] = regexp.MustCompile(`(?i)ipad`)
	os_ref["Android"] = regexp.MustCompile(`(?i)android`)
	os_ref["BlackBerry"] = regexp.MustCompile(`(?i)blackberry`)
	os_ref["SymbianOS"] = regexp.MustCompile(`(?i)symbianos`)
	os_ref["Mobile"] = regexp.MustCompile(`(?i)webos`)
	for key, regex := range os_ref {
		match := regex.FindStringIndex(u)
		if match != nil {
			return key
		}
	}
	return ""
}

func getBrowser(u string) string {
	browser_ref := make(map[string]*regexp.Regexp)
	browser_ref["Internet Explorer"] = regexp.MustCompile(`(?i)msie`)
	browser_ref["Firefox"] = regexp.MustCompile(`(?i)firefox`)
	browser_ref["Chrome"] = regexp.MustCompile(`(?i)chrome`)
	browser_ref["Safari"] = regexp.MustCompile(`(?i)safari`)
	browser_ref["Opera"] = regexp.MustCompile(`(?i)opera`)
	browser_ref["Netscape"] = regexp.MustCompile(`(?i)netscape`)
	browser_ref["Maxthon"] = regexp.MustCompile(`(?i)maxthon`)
	browser_ref["Konqueror"] = regexp.MustCompile(`(?i)konqueror`)
	browser_ref["Edge"] = regexp.MustCompile(`(?i)edge`)
	browser_ref["Other"] = regexp.MustCompile(`(?i)other`)
	browser_ref["Handheld Browser"] = regexp.MustCompile(`(?i)mobile`)
	for key, regex := range browser_ref {
		match := regex.FindStringIndex(u)
		if match != nil {
			return key
		}
	}
	return ""
}

func createHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	bs := hex.EncodeToString(h.Sum(nil))
	return bs
}

func fingerPrint(r *gin.Context) string {
	id := getUserId(r)

	ua := getUserAgent(r)
	ip := getUserIp(r)
	os := getOs(ua)
	br := getBrowser(ua)

	mix := createHash(ip + os + br)
	return id + mix
}
