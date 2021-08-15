package main

import (
	"github.com/lanzay/scan-server/app/scan/delivery/scan_http_handler"
	scan_repo_files "github.com/lanzay/scan-server/app/scan/repository"
	scan_usecase "github.com/lanzay/scan-server/app/scan/usecase"
	server_http "github.com/lanzay/scan-server/servers"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	scanRepo := scan_repo_files.NewScanRepo("./data")
	scanUC := scan_usecase.NewScanUseCase(scanRepo)
	scan_http_handler.HandlerRegister(mux, scanUC)

	server := server_http.New()
	err := server.Run("3030", mux)
	if err != nil {
		panic(err)
	}
}
