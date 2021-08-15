package scan_http_handler

import (
	"encoding/json"
	"github.com/lanzay/scan-server/app/scan"
	"io"
	"log"
	"net/http"
)

type handler struct {
	uc scan.ScanUseCaseI
}

func NewHttp(uc scan.ScanUseCaseI) *handler {
	return &handler{
		uc: uc,
	}
}

func (h *handler) ScanAny(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		h.ScanPOST(w, r)
	case http.MethodGet:
		h.ScanGET(w, r)
	case http.MethodDelete:
		h.DelByJob(w, r)
	}
}

func (h *handler) ScanPOST(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)
	r.Body.Close()
	var scanReq *scanReq
	err := json.Unmarshal(body, &scanReq)
	if err != nil {
		panic(err)
	}

	err = h.uc.ScanBarcode(r.Context(), scanReq.Job, scanReq.Barcode, 1)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		log.Println("[E]", err)
		return
	}
	// TODO RES

}

func (h *handler) ScanGET(w http.ResponseWriter, r *http.Request) {

	job := r.FormValue("job")
	scans, _ := h.uc.GetBarcodesByJob(r.Context(), job)

	res := scanRes{Job: job}
	for _, scan := range scans {
		res.Barcodes = append(res.Barcodes, scan.Barcode)
	}

	body, _ := json.Marshal(res)
	w.Write(body)

}

func (h *handler) JobsAny(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		h.NewJob(w, r)
	case http.MethodGet:
		h.GetJobs(w, r)
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

	job, err := h.uc.StartJob(r.Context(), newJobReq.JobName, newJobReq.Comment)
	bodyRes, _ := json.Marshal(job)
	_, _ = w.Write(bodyRes)
}

func (h *handler) GetJobs(w http.ResponseWriter, r *http.Request) {

	jobs, _ := h.uc.GetJobs(r.Context())
	body, _ := json.Marshal(jobs)
	_, _ = w.Write(body)
}

func (h *handler) DelByJob(w http.ResponseWriter, r *http.Request) {

	job := r.FormValue("job")
	_, _ = h.uc.EndJob(r.Context(), job)
}
