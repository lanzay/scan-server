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
	jobsOpen map[string][]models.Job // [JobName][]Jobs

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
		jobsOpen: map[string][]models.Job{},
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
		fileName.Name()

		jobData, err := loadJobFromFile(s.dir, status, fileName.Name())
		if err != nil {
			log.Panicln("[E]", err)
		}
		s.jobsOpen[jobData.Job.Name] = append(s.jobsOpen[jobData.Job.Name], jobData.Job)
	}
	return err
}

func (s scanRepo) CreateJob(ctx context.Context, job models.Job) (*models.Job, error) {

	uuid4, _ := uuid.NewUUID()
	job.ID = uuid4.String()
	jobData := jobDataStruct{
		Job:      job,
		Barcodes: nil,
	}
	err := SaveJob(s.dir, jobData)

	s.jobsOpen[job.Name] = append(s.jobsOpen[job.Name], job)

	return &job, err
}

func (s scanRepo) CloseJob(ctx context.Context, job models.Job) (*models.Job, error) {

	jobData, err := LoadJob(s.dir, job)
	jobData.Job.EndAt = time.Now().In(MSK)
	err = SaveJob(s.dir, *jobData)

	{ // index del
		var jobsNew []models.Job
		for _, jobOld := range s.jobsOpen[job.Name] {
			if jobOld.ID == job.ID {
				continue
			}
			jobsNew = append(jobsNew)
		}
		s.jobsOpen[job.Name] = jobsNew
	}

	return &jobData.Job, err
}

func (s scanRepo) GetJobs(ctx context.Context) ([]models.Job, error) {

	jobs := []models.Job{}
	for _, jobsOpen := range s.jobsOpen {
		for _, job := range jobsOpen {
			jobs = append(jobs, job)
		}
	}
	return jobs, nil
}

func (s scanRepo) GetBarcodesByJob(ctx context.Context, jobName string) ([]models.Barcode, error) {

	jobs, ok := s.jobsOpen[jobName]
	if !ok {
		return nil, errors.New("Cannot open job: " + jobName)
	}
	if len(jobs) != 1 {
		return nil, errors.New("More then 1 open job by name: " + jobName)
	}

	jobData, err := loadJobFromFile(s.dir, "open", getFileNameByJob(jobs[0]))
	return jobData.Barcodes, err
}

func (s scanRepo) AddBarcode(ctx context.Context, jobName, barcode string, delta int) error {

	jobs, ok := s.jobsOpen[jobName]
	if !ok {
		return errors.New("Cannot open job: " + jobName)
	}
	if len(jobs) != 1 {
		return errors.New("More then 1 open job by name: " + jobName)
	}

	barcpdes, _ := s.GetBarcodesByJob(ctx, jobName)
	barcpdes = append(barcpdes, models.Barcode{
		Barcode: barcode,
		Count:   delta,
		ScanAt:  time.Now().In(MSK),
	})

	err := SaveJob(s.dir, jobDataStruct{
		Job:      jobs[0],
		Barcodes: barcpdes,
	})

	return err
}
