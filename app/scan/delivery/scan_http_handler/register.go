package scan_http_handler

import (
	"github.com/lanzay/scan-server/app/scan"
	"log"
	"net/http"
)

func HandlerRegister(mux *http.ServeMux, uc scan.ScanUseCaseI) {

	h := NewHttp(uc)
	group := "/api/v1.0"

	group = "/api/v1.0/jobs"
	mux.HandleFunc(group+"", h.GetJobs)
	mux.HandleFunc(group+"/", h.GetJobs)
	log.Println("[I] ROUTER GET ", group)

	group = "/api/v1.0/job"
	mux.HandleFunc(group+"", h.JobAny)
	mux.HandleFunc(group+"/", h.JobAny)
	log.Println("[I] ROUTER POST", group+"/new {job_name, comment}")
	log.Println("[I] ROUTER GET ", group+"/{id}")
	log.Println("[I] ROUTER PUT ", group+"/{id}/scan/{barcode}/{count}")
	log.Println("[I] ROUTER GET ", group+"/{id}/close")
}
