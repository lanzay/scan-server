package server_http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type server_http struct {
	server *http.Server
}

func New() *server_http {
	return &server_http{}
}

func (xx *server_http) Run(port string, mux *http.ServeMux) error {

	mux.HandleFunc("/ping", xx.ping)

	xx.server = &http.Server{
		Addr:           ":" + port,
		Handler:        mwCORS(mux),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("[D] webserver start on", "http://localhost"+xx.server.Addr)

	go func() {
		if err := xx.server.ListenAndServe(); err != nil {
			log.Panicln(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return xx.server.Shutdown(ctx)
}

func (xx *server_http) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"ok"}`))
}

func mwCORS(next http.Handler) http.Handler {

	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		next.ServeHTTP(w, r)
	})
}
