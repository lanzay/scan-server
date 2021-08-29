package www

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
)

type www struct{}

func New() *www {
	return &www{}
}

func (xx *www) HandlerRegister(mux *http.ServeMux) {

	mux.HandleFunc("/", xx.handlerIndex)
	log.Println("[I] ROUTER", "/")
}

func (xx *www) handlerIndex(w http.ResponseWriter, r *http.Request) {

	statDir := "./servers/www/site_v1/build"
	index := statDir + "/index.html"
	//index := filepath.Join(statDir, "/index.html")

	path := filepath.Join(statDir, r.URL.Path)
	_, err := os.Stat(path)
	var body []byte
	if err == nil {
		body, err = os.ReadFile(path)
	}
	if err != nil {
		body, err = os.ReadFile(index)
	}
	_, _ = w.Write(body)
}

func (xx *www) handlerOk(w http.ResponseWriter, r *http.Request) {

	log.Println("[D] req", r.URL)

	//ctx := r.Context()
	//uid := ctx.Value(COOKIE_NAME).(string)
	//log.Println("[I] uid", uid)

	rBody, _ := httputil.DumpRequest(r, true)
	log.Println("[D] req Ok", string(rBody))

	_, _ = w.Write([]byte(`"status":"ok"`))
}
