package models

import "time"

type Barcode struct {
	BarcodeRaw string    `json:"barcode_raw"`
	Barcode    string    `json:"barcode"`
	Count      int       `json:"count"`
	LastScanAt time.Time `json:"last_scan_at"`
}
