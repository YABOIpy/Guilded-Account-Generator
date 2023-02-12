package src

import "net/http"

func (Gu *Guilded) HeaderAuth(req *http.Request, hex int) {
	ID := Gu.RandHex(hex)
	for x, o := range map[string]string{
		"accept-encoding":     "gzip, deflate, br",
		"accept-language":     "en-GB,en-US;q=0.9,en;q=0.8",
		"content-type":        "application/json",
		"guilded-device-id":   ID,
		"guilded-device-type": "desktop",
		"guilded-stag":        "c4740afd3f3e4d63d365d826139de166",
		"origin":              "https://www.guilded.gg",
		"referer":             "https://www.guilded.gg/",
		"sec-fetch-dest":      "empty",
		"sec-fetch-mode":      "cors",
		"sec-fetch-site":      "same-origin",
		"user-agent":          "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
		"x-requested-with":    "XMLHttpRequest",
	} {
		req.Header.Set(x, o)
	}
}
