package visitor

import (
	"net/http"
)

type VisitorInfo struct {
	IP             string `json:"ip"`
	URL            string `json:"url"`
	Referrer       string `json:"referrer"`
	UserAgent      string `json:"user_agent"`
	Platform       string `json:"platform"`
	Language       string `json:"language"`
	ScreenWidth    int    `json:"screen_width"`
	ScreenHeight   int    `json:"screen_height"`
	ViewportWidth  int    `json:"viewport_width"`
	ViewportHeight int    `json:"viewport_height"`
}

func ParseVisitorInfo(r *http.Request, data map[string]interface{}) VisitorInfo {
	ip := getIPAddress(r)

	screenWidth := int(data["screen_width"].(float64))
	screenHeight := int(data["screen_height"].(float64))
	viewportWidth := int(data["viewport_width"].(float64))
	viewportHeight := int(data["viewport_height"].(float64))

	return VisitorInfo{
		IP:             ip,
		URL:            r.URL.String(),
		Referrer:       r.Referer(),
		UserAgent:      r.UserAgent(),
		Platform:       r.Header.Get("User-Agent"),
		Language:       r.Header.Get("Accept-Language"),
		ScreenWidth:    screenWidth,
		ScreenHeight:   screenHeight,
		ViewportWidth:  viewportWidth,
		ViewportHeight: viewportHeight,
	}
}

func getIPAddress(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
}
