package scan_usecase

import (
	"context"
	"github.com/lanzay/scan-server/app/scan"
	"github.com/lanzay/scan-server/models"
	"time"
)

var _ scan.ScanUseCaseI = &scanUseCase{}

type scanUseCase struct {
	repo scan.ScanRepoI
}

func NewScanUseCase(repo scan.ScanRepoI) scan.ScanUseCaseI {
	return &scanUseCase{
		repo: repo,
	}
}

func (uc scanUseCase) StartJob(ctx context.Context, jobName string, comment string) (*models.Job, error) {

	job, err := uc.repo.CreateJob(ctx, models.Job{
		Name:    jobName,
		Comment: comment,
		StartAt: time.Now(),
	})
	return job, err
}

func (uc scanUseCase) GetJobs(ctx context.Context) ([]models.Job, error) {
	return uc.repo.GetJobs(ctx)
}

func (uc scanUseCase) GetBarcodesByJob(ctx context.Context, jobName string) ([]models.Barcode, error) {
	return uc.repo.GetBarcodesByJob(ctx, jobName)
}

func (uc scanUseCase) EndJob(ctx context.Context, jobName string) (*models.Job, error) {

	job, err := uc.repo.CloseJob(ctx, models.Job{
		Name:  jobName,
		EndAt: time.Now(),
	})
	return job, err
}

func (uc scanUseCase) ScanBarcode(ctx context.Context, jobName, barcode string, delta int) error {

	err := uc.repo.AddBarcode(ctx, jobName, barcode, delta)
	return err
}
