package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func New(target string) http.Handler {

	u, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(u)
}
