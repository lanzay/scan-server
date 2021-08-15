package scan

import (
	"context"
	"github.com/lanzay/scan-server/models"
)

type ScanRepoI interface {
	CreateJob(ctx context.Context, job models.Job) (*models.Job, error)
	CloseJob(ctx context.Context, job models.Job) (*models.Job, error)
	GetJobs(ctx context.Context) ([]models.Job, error)
	GetBarcodesByJob(ctx context.Context, jobName string) ([]models.Barcode, error)
	AddBarcode(ctx context.Context, jobName, barcode string, delta int) error
}
