package scan_http_handler

import (
	"encoding/json"
	"github.com/lanzay/scan-server/app/scan"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type handler struct {
	uc scan.ScanUseCaseI
}

func NewHttp(uc scan.ScanUseCaseI) *handler {
	return &handler{
		uc: uc,
	}
}

func (h *handler) ScanBarcode(w http.ResponseWriter, r *http.Request) {

	//PUT  /api/v1.0/job/{id}/scan/{barcode}/{count}
	//    0  1    2   3   4    5      6         7

	var err error
	req := &scanReq{}

	paths := strings.Split(r.URL.Path, "/")
	if len(paths) < 6 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	req.JobId = paths[4]
	req.Barcode = paths[6]

	req.Count = 1
	if len(paths) < 7 {
		req.Count, _ = strconv.Atoi(paths[7])
	}

	job, err := h.uc.ScanBarcode(r.Context(), req.JobId, req.Barcode, req.Count)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		log.Println("[E]", err)
		return
	}

	body, _ := json.MarshalIndent(job, "", "    ")
	_, _ = w.Write(body)

}

func (h *handler) JobAny(w http.ResponseWriter, r *http.Request) {

	//POST /job/new {job_name, comment}
	//GET  /job/{id}
	//PUT  /job/{id}/scan/{barcode}/{count}
	//GET  /job/{id}/close

	//  /api/v1.0/job/new
	//  /api/v1.0/job/2021-08-28_22-40-50_9f9b2be4
	//  /api/v1.0/job/2021-08-28_22-40-50_9f9b2be4/close
	// 0  1    2   3   4                            5

	paths := strings.Split(r.URL.Path, "/")

	switch r.Method {
	case http.MethodPost:
		h.NewJob(w, r)
	case http.MethodGet:
		if len(paths) == 6 && strings.EqualFold(paths[5], "close") {
			h.CloseJob(w, r)
			return
		}
		h.GetJob(w, r)

	case http.MethodPut:
		h.ScanBarcode(w, r)
	}
}

func (h *handler) NewJob(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)
	r.Body.Close()
	var newJobReq *newJobReq
	err := json.Unmarshal(body, &newJobReq)
	if err != nil {
		panic(err)
	}

	job, err := h.uc.NewJob(r.Context(), newJobReq.JobName, newJobReq.Comment)
	bodyRes, _ := json.Marshal(job)
	_, _ = w.Write(bodyRes)
}

func (h *handler) GetJobs(w http.ResponseWriter, r *http.Request) {

	jobs, _ := h.uc.GetJobs(r.Context())
	body, _ := json.Marshal(jobs)
	_, _ = w.Write(body)
}

func (h *handler) GetJob(w http.ResponseWriter, r *http.Request) {

	jobId := strings.Split(r.URL.Path, "/")[4]
	job, _ := h.uc.GetJob(r.Context(), jobId)
	body, _ := json.Marshal(job)
	_, _ = w.Write(body)
}

func (h *handler) CloseJob(w http.ResponseWriter, r *http.Request) {

	job := strings.Split(r.URL.Path, "/")[4]
	_, _ = h.uc.CloseJob(r.Context(), job)
}
