package scan_repository_csv

import (
	"encoding/json"
	"github.com/lanzay/scan-server/app/scan"
	"github.com/lanzay/scan-server/models"
	"log"
	"os"
	"github.com/google/uuid"
)

var _ scan.ScanRepoI = &scanRepo{}

type scanRepo struct {
	dir string
}

func NewScanRepo(dir string) scan.ScanRepoI {
	return &scanRepo{
		dir: dir,
	}
}

func (s scanRepo) CreateJob(job models.Job) (*models.Job, error) {

	job.ID = uuid.New().String()
	fileName := "_" + getFileNameByJob(job)
	fn, err := os.Create(fileName)
	if err != nil {
		log.Panicln("[E] ", err)
	}

	data := fileStruct{
		Job:      job,
		Barcodes: nil,
	}
	body, _ := json.MarshalIndent(data, "", "    ")
	_, err = fn.Write(body)
	_ = fn.Close()

	return &job, err
}

func (s scanRepo) CloseJob(job models.Job) (*models.Job, error) {
	panic("implement me")
}

func (s scanRepo) GetJobs() ([]models.Job, error) {
	panic("implement me")
}

func (s scanRepo) GetBarcodesByJob(jobName string) ([]models.Barcode, error) {
	panic("implement me")
}

func (s scanRepo) AddBarcode(jobName, barcode string, delta int) error {
	panic("implement me")
}
