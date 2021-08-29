package scan_usecase

import (
	"strings"
	"testing"
)

func TestBarcodeNormalisation(t *testing.T) {

	barcodes := []struct {
		raw string
		ok  string
	}{
		{"046801340477040121114103001", "0104680134047704210121114103001"},
		{"(01)04680134047704(21)0121114103001", "0104680134047704210121114103001"},
		{"4680134047704", "4680134047704"},
		{"0104680134047704210121114103001", "0104680134047704210121114103001"},
		{"010468013404770421012111410300191Ejhvx", "0104680134047704210121114103001"},
		{"(01)04680134047704(21)012111410300191Ejhvx", "0104680134047704210121114103001"},
	}

	for _, barcode := range barcodes {
		norm := BarcodeNormalisation(barcode.raw)
		if !strings.EqualFold(barcode.ok, norm) {
			t.Fatal()
		}
	}
}
