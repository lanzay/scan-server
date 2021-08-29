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

const FILE_NAME_DATE = "2006-02-01_15-04-05"

var MSK = time.FixedZone("MSK", 3*60*60)

func getFileNameByJob(jobHeader models.JobHeader) string {

	// UUID time based
	fileName := jobHeader.ID + ".json"
	return fileName
}

func loadJobFromFile(dir, status, fileName string) (*models.Job, error) {

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
	var job *models.Job
	err = json.Unmarshal(body, &job)
	if err != nil {
		log.Panicln("[E]", err)
	}

	return job, err
}

func LoadJob(dir string, jobHeader models.JobHeader) (*models.Job, error) {

	fileName := getFileNameByJob(jobHeader)
	status := "close"
	if jobHeader.EndAt.IsZero() {
		status = "open"
	}
	return loadJobFromFile(dir, status, fileName)
}

func LoadJobById(dir string, jobId string) (*models.Job, error) {

	// TODO check from close
	fileName := jobId + ".json"
	status := "open"
	return loadJobFromFile(dir, status, fileName)
}

func SaveJob(dir string, job models.Job) error {

	fileName := getFileNameByJob(job.JobHeader)
	status := "close"
	if job.EndAt.IsZero() {
		status = "open"
	}

	// create
	fileN := path.Join(dir, status, fileName)
	fn, err := os.OpenFile(fileN, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Panicln("[E] ", err)
	}

	// save
	body, _ := json.MarshalIndent(job, "", "    ")
	_, err = fn.Write(body)
	if err != nil {
		log.Panicln("[E]", err)
	}
	_ = fn.Close()
	log.Println("[D] FILE.CREATE", fileN)

	// remove not End JobHeader
	if !job.EndAt.IsZero() {
		fileN := path.Join(dir, "open", fileName)
		err = os.Remove(fileN)
		if err != nil {
			log.Panicln(err)
		}
		log.Println("[D] FILE.DEL   ", fileN)
	}

	return err
}
