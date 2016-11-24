package utils

import (
	"net/http"
)

func GetClientIP(req *http.Request) string {
	if ip := req.Header.Get("X-Real-Ip"); ip != "" {
		return ip
	}

	if ip := req.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}

	if ip := req.Header.Get("RemoteAddr"); ip != "" {
		return ip
	}

	return ""
}
