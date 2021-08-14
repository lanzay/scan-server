package scan_usecase

import (
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

func (uc scanUseCase) StartJob(jobName string, comment string) (*models.Job, error) {

	job, err := uc.repo.CreateJob(models.Job{
		Name:    jobName,
		Comment: comment,
		StartAt: time.Now(),
	})
	return job, err
}

func (uc scanUseCase) GetJobs() ([]models.Job, error) {
	return uc.repo.GetJobs()
}

func (uc scanUseCase) GetBarcodesByJob(jobName string) ([]models.Barcode, error) {
	return uc.repo.GetBarcodesByJob(jobName)
}

func (uc scanUseCase) EndJob(jobName string) (*models.Job, error) {

	job, err := uc.repo.CloseJob(models.Job{
		Name:  jobName,
		EndAt: time.Now(),
	})
	return job, err
}

func (uc scanUseCase) ScanBarcode(jobName, barcode string, delta int) error {

	err := uc.repo.AddBarcode(jobName, barcode, delta)
	return err
}
