package scan

import (
	"context"
	"github.com/lanzay/scan-server/models"
)

type ScanRepoI interface {
	CreateJob(ctx context.Context, job models.JobHeader) (*models.Job, error)
	CloseJob(ctx context.Context, jobId string) (*models.Job, error)
	GetJobs(ctx context.Context) ([]models.JobHeader, error)
	GetJob(ctx context.Context, jobId string) (*models.Job, error)
	AddBarcode(ctx context.Context, jobId, barcodeRaw, barcode string, delta int) (*models.Job, error)
}
