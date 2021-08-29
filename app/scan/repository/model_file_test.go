package scan_repository_file

import (
	"github.com/lanzay/scan-server/models"
	"strings"
	"testing"
	"time"
)

func TestGetFileNameByJob(t *testing.T) {

	jobTest := models.JobHeader{
		ID:      "1111-2222-3333-4444",
		Name:    "",
		Comment: "",
		StartAt: time.Date(1979, 07, 02, 15, 01, 02, 3, MSK),
		EndAt:   time.Time{},
	}

	fileName := getFileNameByJob(jobTest)

	if !strings.EqualFold(fileName, "1979-02-07_15-01-02_1111-2222-3333-4444.json") {
		t.Fatal()
	}

}
