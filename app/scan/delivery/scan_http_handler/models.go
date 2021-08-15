package scan_http_handler

type scanReq struct {
	Job     string `json:"job"`
	Barcode string `json:"barcode"`
}

type scanRes struct {
	Job      string   `json:"job"`
	Barcodes []string `json:"barcodes"`
}
type newJobReq struct {
	JobName string `json:"job_name"`
	Comment string `json:"comment"`
}
