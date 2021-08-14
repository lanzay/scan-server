package scan

import "github.com/lanzay/scan-server/models"

type ScanRepoI interface {
	CreateJob(job models.Job) (*models.Job, error)
	CloseJob(job models.Job) (*models.Job, error)
	GetJobs() ([]models.Job, error)
	GetBarcodesByJob(jobName string) ([]models.Barcode, error)
	AddBarcode(jobName, barcode string, delta int) error
}
