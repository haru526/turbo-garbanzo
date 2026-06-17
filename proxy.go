package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func main() {
	// フラグの定義
	listenAddr := flag.String("listen", ":8080", "Listen address (e.g., :8080)")
	targetURL := flag.String("target", "http://localhost:3000", "Target backend URL (e.g., http://localhost:3000)")
	flag.Parse()

	// ターゲットURLをパース
	target, err := url.Parse(*targetURL)
	if err != nil {
		log.Fatalf("Invalid target URL: %v", err)
	}

	// リバースプロキシを作成
	proxy := httputil.NewSingleHostReverseProxy(target)

	// エラーハンドラーをカスタマイズ
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Proxy error: %v (path: %s, method: %s)", err, r.RequestURI, r.Method)
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "Bad Gateway: %v\n", err)
	}

	// リクエストハンドラーを定義
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request: %s %s (client: %s)", r.Method, r.RequestURI, getClientIP(r))
		proxy.ServeHTTP(w, r)
	}).ServeHTTP(w, r)

	// HTTPサーバーの設定
	server := &http.Server{
		Addr: *listenAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Incoming request: %s %s (client: %s)", r.Method, r.RequestURI, getClientIP(r))
			proxy.ServeHTTP(w, r)
		}),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// サーバー起動ログ
	log.Printf("Starting HTTP reverse proxy")
	log.Printf("  Listen: %s", *listenAddr)
	log.Printf("  Target: %s", target.String())
	log.Printf("")

	// サーバーを起動
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}

// クライアントのIPアドレスを取得
func getClientIP(r *http.Request) string {
	// X-Forwarded-Forヘッダーを確��
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		return forwarded
	}

	// RemoteAddrから取得
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	return ip
}
