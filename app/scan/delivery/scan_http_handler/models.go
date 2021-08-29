package scan_http_handler

type scanReq struct {
	JobId   string `json:"job_id"`
	Barcode string `json:"barcode"`
	Count   int    `json:"count"`
}

type newJobReq struct {
	JobName string `json:"job_name"`
	Comment string `json:"comment"`
}
