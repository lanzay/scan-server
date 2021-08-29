package scan_repository_file

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/lanzay/scan-server/app/scan"
	"github.com/lanzay/scan-server/models"
	"log"
	"os"
	"path"
	"time"
)

var _ scan.ScanRepoI = &scanRepo{}

type scanRepo struct {
	dir      string
	jobsOpen map[string]models.Job // [JobId]Job

}

func NewScanRepo(dir string) scan.ScanRepoI {

	info, err := os.Stat(dir)
	if err != nil || !info.IsDir() {
		errMk := os.MkdirAll(dir, 0666)
		if errMk != nil {
			panic(err)
		}
	}
	_ = os.MkdirAll(path.Join(dir, "open"), 0666)
	_ = os.MkdirAll(path.Join(dir, "close"), 0666)

	repo := &scanRepo{
		dir:      dir,
		jobsOpen: map[string]models.Job{},
	}
	err = repo.loadJobsOpenOnly()

	return repo
}

func (s scanRepo) loadJobsOpenOnly() error {

	status := "open"
	dirs, err := os.ReadDir(path.Join(s.dir, status))
	for _, fileName := range dirs {
		if fileName.IsDir() {
			continue
		}
		//fileName.Name()

		jobData, err := loadJobFromFile(s.dir, status, fileName.Name())
		if err != nil {
			log.Panicln("[E]", err)
		}
		s.jobsOpen[jobData.ID] = *jobData
	}
	return err
}

func (s scanRepo) CreateJob(ctx context.Context, jobHeader models.JobHeader) (*models.Job, error) {

	jobHeader.ID = time.Now().In(MSK).Format("2006-01-02_15-04-05") + "_" + uuid.New().String()[:8]

	//uuid4, _ := uuid.NewUUID()
	//job.ID = uuid4.String()

	job := models.Job{
		JobHeader: jobHeader,
		Barcodes:  nil,
	}

	err := SaveJob(s.dir, job)
	s.jobsOpen[jobHeader.ID] = job

	return &job, err
}

func (s scanRepo) CloseJob(ctx context.Context, jobId string) (*models.Job, error) {

	job, err := LoadJobById(s.dir, jobId)
	job.EndAt = time.Now().In(MSK)
	err = SaveJob(s.dir, *job)

	{ // index del
		delete(s.jobsOpen, job.ID)
	}

	return job, err
}

func (s scanRepo) GetJobs(ctx context.Context) ([]models.JobHeader, error) {

	var jobsHeader []models.JobHeader
	for _, jobsOpen := range s.jobsOpen {
		jobsHeader = append(jobsHeader, jobsOpen.JobHeader)

	}
	return jobsHeader, nil
}

func (s scanRepo) GetJob(ctx context.Context, jobId string) (*models.Job, error) {

	job, ok := s.jobsOpen[jobId]
	if !ok {
		return nil, errors.New("Cannot open job: " + jobId)
	}

	return &job, nil
}

func (s scanRepo) AddBarcode(ctx context.Context, jobId, barcodeRaw, barcode string, delta int) (*models.Job, error) {

	job, ok := s.jobsOpen[jobId]
	if !ok {
		return nil, errors.New("Cannot open job: " + jobId)
	}

	rec := models.Barcode{
		BarcodeRaw: barcodeRaw,
		Barcode:    barcode,
		Count:      delta,
		LastScanAt: time.Now().In(MSK),
	}
	job.BarcodesDetail = append(job.BarcodesDetail, rec)

	skip := false
	for i := range job.Barcodes {
		if job.Barcodes[i].Barcode == barcode {
			if len(barcode) == 31 {
				skip = true
				continue
			}
			job.Barcodes[i].Count += delta
			job.Barcodes[i].LastScanAt = time.Now().In(MSK)
			skip = true
		}
	}
	if !skip {
		rec.BarcodeRaw = ""
		job.Barcodes = append(job.Barcodes, rec)
	}

	s.jobsOpen[jobId] = job
	err := SaveJob(s.dir, job)

	return &job, err
}
