package scan

import "github.com/lanzay/scan-server/models"

type ScanUseCaseI interface {
	StartJob(jobName string, comment string) (*models.Job, error) // 1. Создать Job
	GetJobs() ([]models.Job, error)                               // 2. Получить список Job
	GetBarcodesByJob(jobName string) ([]models.Barcode, error)    // 3. Получить состав Job
	EndJob(jobName string) (*models.Job, error)                   // 4. Удалять Job
	ScanBarcode(jobName string, barcode string, delta int) error  // 5. Сканируем ШК - добавить/удалить
}
