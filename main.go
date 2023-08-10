package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	listen string
	addr   string
)

func init() {
	flag.StringVar(&listen, "l", ":8080", "listen on ip:port")
	flag.StringVar(&addr, "r", "", "reverse proxy addr")
}

func main() {
	flag.Parse()
	if addr == "" {
		fmt.Println("reverse proxy addr cannot be empty")
		return
	}
	target, err := url.Parse(addr)
	if err != nil {
		panic(err)
	}
	proxy := NewReverseProxy(target)
	http.Handle("/", proxy)
	http.ListenAndServe(listen, nil)
}

func NewReverseProxy(target *url.URL) *httputil.ReverseProxy {
	rewriteFunc := func(r *httputil.ProxyRequest) {
		r.Out.Host = ""
		r.Out.URL.Scheme = "https"
		r.Out.URL.Host = target.Host
		r.SetXForwarded()
	}

	errFunc := func(w http.ResponseWriter, r *http.Request, err error) {
		fmt.Printf("http proxy error: %v\nreq: %+v", err, target)
		w.WriteHeader(http.StatusBadGateway)
	}

	return &httputil.ReverseProxy{
		Rewrite:      rewriteFunc,
		ErrorHandler: errFunc,
	}
}
