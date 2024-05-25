package webhook

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/majoramari/visitor-logger/visitor"
)

type WebhookPayload struct {
	Embeds []Embed `json:"embeds,omitempty"`
}

type Embed struct {
	Content     string  `json:"content,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Color       int     `json:"color,omitempty"`
	Fields      []Field `json:"fields,omitempty"`
}

type Field struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

type IPGeolocationResponse struct {
	ContinentName string `json:"continent_name"`
	CountryName   string `json:"country_name"`
	City          string `json:"city"`
}

func CreatePayload(visitorInfo visitor.VisitorInfo) WebhookPayload {
	country := getCountryFromIP(visitorInfo.IP)

	embed := Embed{
		Title: "Logged Visitor Information",
		Color: 0x4fa9f0,
		Fields: []Field{
			{Name: "IP", Value: visitorInfo.IP, Inline: true},
			{Name: "Country", Value: country, Inline: true},
			{Name: "URL", Value: visitorInfo.URL, Inline: true},
			{Name: "Referrer", Value: visitorInfo.Referrer, Inline: true},
			{Name: "User Agent", Value: visitorInfo.UserAgent, Inline: true},
			{Name: "Platform", Value: visitorInfo.Platform, Inline: true},
			{Name: "Language", Value: visitorInfo.Language, Inline: true},
			{Name: "Screen Size", Value: fmt.Sprintf("%d x %d", visitorInfo.ScreenWidth, visitorInfo.ScreenHeight), Inline: true},
			{Name: "Viewport Size", Value: fmt.Sprintf("%d x %d", visitorInfo.ViewportWidth, visitorInfo.ViewportHeight), Inline: true},
		},
	}

	payload := WebhookPayload{
		Embeds: []Embed{embed},
	}

	return payload
}

func getCountryFromIP(ip string) string {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("IP geolocation API key not found in environment variables")
		return "Unknown"
	}

	resp, err := http.Get("https://api.ipgeolocation.io/ipgeo?apiKey=" + apiKey + "&ip=" + ip + "&fields=continent_name,country_name,city")
	if err != nil {
		fmt.Println("Error retrieving country:", err)
		return "Unknown"
	}
	defer resp.Body.Close()

	// Decode JSON response
	var ipGeolocationResponse IPGeolocationResponse
	err = json.NewDecoder(resp.Body).Decode(&ipGeolocationResponse)
	if err != nil {
		fmt.Println("Error decoding country response:", err)
		return "Unknown"
	}

	countryInfo := fmt.Sprintf("%s, %s - %s", ipGeolocationResponse.City, ipGeolocationResponse.CountryName, ipGeolocationResponse.ContinentName)

	return countryInfo
}
