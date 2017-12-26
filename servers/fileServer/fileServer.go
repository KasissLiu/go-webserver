package fileServer

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

type FileServer struct {
	BaseDir string
	Index   string
	Exts    []string
	ExtType map[string]string
}

//静态文件服务器入口
func (this *FileServer) FileOutput(w http.ResponseWriter, r *http.Request) {

	switch path.Ext(r.URL.Path) {
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "text/javascript")
	case ".jpg":
		fallthrough
	case ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".gif":
		w.Header().Set("Content-Type", "image/gif")
	case ".html":
		fallthrough
	case ".htm":
		w.Header().Set("Content-Type", "text/html")
	default:
		w.Header().Set("Content-Type", "text/plain")
	}

	content, err := this.readFile(this.BaseDir + r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		io.WriteString(w, "------------ 504 Bad GateWay -------------")
		return
	}
	io.WriteString(w, string(content))

}

func (this *FileServer) IndexOutput(w http.ResponseWriter) {
	content, err := this.readFile(this.BaseDir + "/" + this.Index)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		io.WriteString(w, "------------ 504 Bad GateWay -------------")
		return
	}
	io.WriteString(w, string(content))
}

func (this *FileServer) readFile(path string) ([]byte, error) {
	file, err := os.Open(path)

	defer file.Close()

	if err != nil {
		return make([]byte, 0, 0), err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return make([]byte, 0, 0), err
	}
	return content, nil
}

func New(baseDir string, index string, exts []string, extType map[string]string) *FileServer {
	return &FileServer{baseDir, index, exts, extType}
}
