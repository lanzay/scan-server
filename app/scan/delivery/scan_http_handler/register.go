package scan_http_handler

import (
	"github.com/lanzay/scan-server/app/scan"
	"log"
	"net/http"
)

func HandlerRegister(mux *http.ServeMux, uc scan.ScanUseCaseI) {

	h := NewHttp(uc)
	group := "/api/v1.0/scan"

	mux.HandleFunc(group+"", h.ScanAny)
	log.Println("[I] ROUTER", group+"?job=default")
	log.Println("[I] ROUTER POST  ", group+"?job=default")
	log.Println("[I] ROUTER GET   ", group+"?job=default")
	log.Println("[I] ROUTER DELETE", group+"?job=default")

	group = "/api/v1.0/jobs"
	mux.HandleFunc(group+"", h.JobsAny)
	mux.HandleFunc(group+"/", h.JobsAny)
	log.Println("[I] ROUTER", group+"?days=1")

}
