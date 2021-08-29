package scan

import (
	"context"
	"github.com/lanzay/scan-server/models"
)

type ScanUseCaseI interface {
	NewJob(ctx context.Context, jobName string, comment string) (*models.Job, error)               // 1. Создать JobHeader
	GetJobs(ctx context.Context) ([]models.JobHeader, error)                                       // 2. Получить список JobHeader
	GetJob(ctx context.Context, jobId string) (*models.Job, error)                                 // 3. Получить состав JobHeader
	CloseJob(ctx context.Context, jobId string) (*models.Job, error)                               // 4. Удалять JobHeader
	ScanBarcode(ctx context.Context, jobId string, barcode string, delta int) (*models.Job, error) // 5. Сканируем ШК - добавить/удалить
}
