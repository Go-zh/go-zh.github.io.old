//usr/bin/env sudo setcap cap_net_bind_service=+ep proxy; exit

// Gerrit HTTP 服务的反向代理。
// 本程序监听 :80 端口，并将所有请求 301 到 :443。
// 同时本程序监听 :443 端口，根据请求的 Host 将请求分发：
// Gerrit 请求转发 localhost:8000，非 Gerrit 请求转发至 localhost:8080.
package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

const GerritHost = "go.minux.info"

func main() {
	go func() {
		server := &http.Server{
			Addr: ":80",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				url := url.URL{
					Scheme: "https",
					Opaque: "//" + r.Host + r.RequestURI,
				}
				log.Printf("redirects [%v] %s to %s\n", r.RemoteAddr, r.RequestURI, url.String())
				http.Redirect(w, r, url.String(), http.StatusMovedPermanently)
			}),
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 12,
		}
		log.Fatal(server.ListenAndServe())
	}()
	server := &http.Server{
		Addr: ":443",
		Handler: &httputil.ReverseProxy{Director: func(r *http.Request) {
			log.Printf("[%s] accessing %s", r.RemoteAddr, r.RequestURI)
			switch r.Host {
			case GerritHost:
				*r.URL = url.URL{
					Scheme: "http",
					Host:   "localhost:8000",
					Opaque: r.RequestURI,
				}
			default:
				*r.URL = url.URL{
					Scheme: "http",
					Host:   "localhost:8080",
					Opaque: r.RequestURI,
				}
			}
		}},
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 16,
	}
	log.Fatal(server.ListenAndServeTLS("$CERT$", "$KEY$"))
}
