package scan_repository_csv

import "github.com/lanzay/scan-server/models"

type fileStruct struct {
	Job      models.Job       `json:"job"`
	Barcodes []models.Barcode `json:"barcodes"`
}

const FILE_NAME_DATE = "2006-02-01_15-04-05"

func getFileNameByJob(job models.Job) string {

	fileName := job.StartAt.Format(FILE_NAME_DATE)
	fileName += "_" + job.ID
	return fileName
}
