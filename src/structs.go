package src

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

type Guilded struct {
	Proxy string `json:"proxy"`
}

var (
	x      = Guilded{}
	p, _   = url.Parse("http://" + x.Config().Proxy)
	Client = http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MaxVersion: tls.VersionTLS13,
			},
			Proxy: http.ProxyURL(p),
		},
	}
	c, r, g, bg, rb, gr, u, clr, yellow = "\u001b[30;1m", "\033[39m", "\033[32m", "\u001b[34;1m", "\u001b[0m", "\u001b[30;1m", "\u001b[4m", "\u001b[38;5;8m", "\033[33m"
)
