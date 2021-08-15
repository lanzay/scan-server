package models

import "time"

type Barcode struct {
	Barcode string    `json:"barcode"`
	Count   int       `json:"count"`
	ScanAt  time.Time `json:"scan_at"`
}
