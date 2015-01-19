//usr/bin/env sudo setcap cap_net_bind_service=+ep proxy; exit

// Gerrit HTTP 服务的反向代理。
// 本程序监听 :80 端口，将有 SSL 证书的请求 301 到 :443。
// 同时本程序监听 :443 端口，根据请求的 Host 将请求分发。
package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// maps hostname to hostname
var ReverseProxyTable = map[string]string{
	"review.go-zh.org": "localhost:8000", // gerrit
	"go-zh.org":        "localhost:8080", // godoc
	"blog.go-zh.org":   "localhost:8081", // blog
	"talks.go-zh.org":  "localhost:8082", // present
}

// maps hostname to whether we have SSL certificate for that hostname
var WantHTTPS = map[string]bool{
	"review.go-zh.org": true,
	"go-zh.org":        true,
}

func main() {
	proxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
		log.Printf("%s accessing %s", req.RemoteAddr, req.RequestURI)
		host := ReverseProxyTable[req.Host]
		if host != "" {
			*req.URL = url.URL{
				Scheme: "http",
				Host:   host,
				Opaque: req.RequestURI,
			}
		} else {
			*req.URL = url.URL{
				Scheme: "http",
				Host:   "localhost:8080",
				Opaque: req.RequestURI,
			}
		}
	}}
	go func() {
		server := &http.Server{
			Addr: ":80",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if WantHTTPS[r.Host] {
					url := url.URL{
						Scheme: "https",
						Opaque: "//" + r.Host + r.RequestURI,
					}
					log.Printf("redirects [%v] %s to %s\n", r.RemoteAddr, r.RequestURI, url.String())
					http.Redirect(w, r, url.String(), http.StatusMovedPermanently)
				} else {
					proxy.ServeHTTP(w, r)
				}
			}),
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 12,
		}
		log.Fatal(server.ListenAndServe())
	}()
	server := &http.Server{
		Addr:           ":443",
		Handler:        proxy,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 16,
	}
	log.Fatal(server.ListenAndServeTLS("$CERT$", "$KEY$"))
}
