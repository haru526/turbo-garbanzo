package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// Parse command-line flags
	listen := flag.String("listen", ":8080", "Address to listen on (e.g., :8080)")
	target := flag.String("target", "http://localhost:3000", "Target backend URL to forward requests to")
	flag.Parse()

	// Parse the target URL
	targetURL, err := url.Parse(*target)
	if err != nil {
		log.Fatalf("Invalid target URL: %v", err)
	}

	// Create a reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Custom error handler to log proxy errors
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Proxy error for %s %s: %v", r.Method, r.RequestURI, err)
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Bad Gateway"))
	}

	// HTTP handler that logs requests and forwards them
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		proxy.ServeHTTP(w, r)
	})

	// Start the server
	log.Printf("Starting reverse proxy on %s, forwarding to %s", *listen, *target)
	if err := http.ListenAndServe(*listen, nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
