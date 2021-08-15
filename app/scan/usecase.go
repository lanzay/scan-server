package scan

import (
	"context"
	"github.com/lanzay/scan-server/models"
)

type ScanUseCaseI interface {
	StartJob(ctx context.Context, jobName string, comment string) (*models.Job, error) // 1. Создать Job
	GetJobs(ctx context.Context) ([]models.Job, error)                                 // 2. Получить список Job
	GetBarcodesByJob(ctx context.Context, jobName string) ([]models.Barcode, error)    // 3. Получить состав Job
	EndJob(ctx context.Context, jobName string) (*models.Job, error)                   // 4. Удалять Job
	ScanBarcode(ctx context.Context, jobName string, barcode string, delta int) error  // 5. Сканируем ШК - добавить/удалить
}
