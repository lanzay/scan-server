package models

import "time"

type Barcode struct {
	Barcode string
	ScanAt  time.Time
}
