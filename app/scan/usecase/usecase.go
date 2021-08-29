package scan_usecase

import (
	"context"
	"github.com/lanzay/scan-server/app/scan"
	"github.com/lanzay/scan-server/models"
	"log"
	"strings"
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

func (uc scanUseCase) NewJob(ctx context.Context, jobName string, comment string) (*models.Job, error) {

	job, err := uc.repo.CreateJob(ctx, models.JobHeader{
		Name:    jobName,
		Comment: comment,
		StartAt: time.Now(),
	})

	log.Println("[D] JOB.NEW", job.ID, job.Name, job.Comment)
	return job, err
}

func (uc scanUseCase) GetJobs(ctx context.Context) ([]models.JobHeader, error) {
	return uc.repo.GetJobs(ctx)
}

func (uc scanUseCase) GetJob(ctx context.Context, jobId string) (*models.Job, error) {
	return uc.repo.GetJob(ctx, jobId)
}

func (uc scanUseCase) CloseJob(ctx context.Context, jobId string) (*models.Job, error) {

	job, err := uc.repo.CloseJob(ctx, jobId)
	return job, err
}

func (uc scanUseCase) ScanBarcode(ctx context.Context, jobId, barcodeRaw string, delta int) (*models.Job, error) {

	barcode := BarcodeNormalisation(barcodeRaw)

	job, err := uc.repo.AddBarcode(ctx, jobId, barcodeRaw, barcode, delta)
	if err != nil {
		jobNew, _ := uc.NewJob(ctx, jobId, "auto create")
		job, err = uc.repo.AddBarcode(ctx, jobNew.ID, barcodeRaw, barcode, delta)
	}

	log.Println("[D] SCAN", jobId, barcode, barcodeRaw)
	return job, err
}

func BarcodeNormalisation(barcodeRaw string) string {

	barcode := barcodeRaw

	barcode = strings.Replace(barcode, ")", "", -1)
	barcode = strings.Replace(barcode, "(", "", -1)
	l := len(barcode)

	if l == 13 {
		return barcode
	}

	if l == 31 && (strings.HasPrefix(barcode, "01") ||
		strings.HasPrefix(barcode, "02")) {
		return barcode
	}

	if l == 27 {
		barcode = "01" + barcode[:14] + "21" + barcode[14:]
	}

	if l > 31 {
		barcode = barcode[:31]
	}

	return barcode
}
