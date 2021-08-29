package scan_repository_file

import (
	"github.com/lanzay/scan-server/models"
	"os"
	"path"
	"testing"
	"time"
)

var (
	dirTest = "./repo_test"
	jobTest = models.JobHeader{
		ID:      "1111-2222-3333-4444",
		Name:    "",
		Comment: "",
		StartAt: time.Date(1979, 07, 02, 15, 01, 02, 3, MSK),
		EndAt:   time.Time{},
	}
)

func TestScanRepo_CreateJob(t *testing.T) {

	rep := NewScanRepo(dirTest)
	jobNew, err := rep.CreateJob(jobTest)
	if err != nil {
		t.FailNow()
	}

	if jobNew.ID == jobTest.ID {
		t.FailNow()
	}

	info, err := os.Stat(path.Join(dirTest, "open", getFileNameByJob(*jobNew)))
	if err != nil || info.IsDir() {
		t.FailNow()
	}
}

func TestScanRepo_CloseJob(t *testing.T) {

	rep := NewScanRepo(dirTest)
	jobNew, err := rep.CreateJob(jobTest)
	if err != nil {
		t.FailNow()
	}

	jobClose, err := rep.CloseJob(*jobNew)
	if err != nil {
		t.FailNow()
	}

	if jobClose.EndAt.IsZero() {
		t.FailNow()
	}
}

func TestScanRepo_GetJobs(t *testing.T) {

	dirTest := path.Join(dirTest, "open_test")
	defer os.RemoveAll(dirTest)

	rep := NewScanRepo(dirTest)

	jobTest.Name = "111"
	rep.CreateJob(jobTest)
	jobTest.Name = "222"
	rep.CreateJob(jobTest)
	jobTest.Comment = "333"
	rep.CreateJob(jobTest)

	jobs, err := rep.GetJobs()
	if err != nil {
		t.FailNow()
	}

	if len(jobs) != 3 {
		t.FailNow()
	}

	for _, job := range jobs {
		if !(job.Name == "111" || job.Name == "222" || job.Comment == "333") {
			t.FailNow()
		}
	}
}

func TestScanRepo_AddBarcode(t *testing.T) {

	rep := NewScanRepo(dirTest)
	jobTest.Name = "test"
	_, err := rep.CreateJob(jobTest)
	if err != nil {
		t.FailNow()
	}

	err = rep.AddBarcode("test", "111", 1)
	if err != nil {
		t.FailNow()
	}
}
