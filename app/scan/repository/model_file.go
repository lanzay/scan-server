package scan_repository_file

import (
	"encoding/json"
	"github.com/lanzay/scan-server/models"
	"io"
	"log"
	"os"
	"path"
	"time"
)

type jobDataStruct struct {
	Job      models.Job       `json:"job"`
	Barcodes []models.Barcode `json:"barcodes,omitempty"`
}

const FILE_NAME_DATE = "2006-02-01_15-04-05"

var MSK = time.FixedZone("MSK", 3*60*60)

func getFileNameByJob(job models.Job) string {

	fileName := job.StartAt.In(MSK).Format(FILE_NAME_DATE)
	fileName += "_" + job.ID
	fileName += ".json"
	return fileName
}

func loadJobFromFile(dir, status, fileName string) (*jobDataStruct, error) {

	fn, err := os.Open(path.Join(dir, status, fileName))
	if err != nil {
		log.Panicln("[E] ", err)
	}

	body, err := io.ReadAll(fn)
	if err != nil {
		log.Panicln("[E]", err)
	}
	_ = fn.Close()

	//
	var jobData *jobDataStruct
	err = json.Unmarshal(body, &jobData)
	if err != nil {
		log.Panicln("[E]", err)
	}

	return jobData, err
}

func LoadJob(dir string, job models.Job) (*jobDataStruct, error) {

	fileName := getFileNameByJob(job)
	status := "close"
	if job.EndAt.IsZero() {
		status = "open"
	}
	return loadJobFromFile(dir, status, fileName)
}

func SaveJob(dir string, jobData jobDataStruct) error {

	fileName := getFileNameByJob(jobData.Job)
	status := ""
	if jobData.Job.EndAt.IsZero() {
		status = "open"
	}

	// create
	fn, err := os.OpenFile(path.Join(dir, status, fileName), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Panicln("[E] ", err)
	}

	// save
	body, _ := json.MarshalIndent(jobData, "", "    ")
	_, err = fn.Write(body)
	if err != nil {
		log.Panicln("[E]", err)
	}
	_ = fn.Close()

	// remove not End Job
	if !jobData.Job.EndAt.IsZero() {
		os.Remove(path.Join("open", fileName))
	}

	return err
}
