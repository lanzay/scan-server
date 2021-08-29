package models

import "time"

type (
	Job struct {
		JobHeader
		Barcodes       []Barcode `json:"barcodes,omitempty"`        // Count total by SCU
		BarcodesDetail []Barcode `json:"barcodes_detail,omitempty"` // line by line with time
	}

	JobHeader struct {
		ID      string    `json:"id"`
		Name    string    `json:"name"`
		Comment string    `json:"comment"`
		StartAt time.Time `json:"start_at"`
		EndAt   time.Time `json:"end_at,omitempty"`
	}
)
