package controller

import (
	"github.com/tjz101/caffbox/util"
	"github.com/tjz101/caffmux"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ViewController struct {
	caffmux.Controller
}

func (c *ViewController) Get() {
	r := c.Content.Request
	w := c.Content.ResponseWriter

	file := c.Content.Params["file"]
	if file == "" {
		http.NotFound(w, r)
		return
	}
	filePath := util.GetPhysicalPath(file)
	fla, _ := util.FileExists(filePath)
	if !fla {
		http.NotFound(w, r)
		return
	}
	displayname := r.FormValue("displayname")

	if displayname == "" {
		displayname = filepath.Base(file)
	}

	f, err := os.Open(filePath)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer f.Close()
	buff, err := ioutil.ReadAll(f)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	suffix := filePath[strings.LastIndex(filePath, "."):]
	contentType := ""
	switch suffix {
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "text/javascript"
	case ".html":
		contentType = "text/html"
	case ".txt":
		contentType = "text/plain"
	case ".jpeg":
		contentType = "image/jpeg"
	case ".jpg":
		contentType = "image/jpeg"
	case ".gif":
		contentType = "image/gif"
	case ".png":
		contentType = "image/png"
	case ".bmp":
		contentType = "image/bmp"
	case ".xml":
		contentType = "text/xml"
	default:
	}
	if contentType == "" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", contentType)
	w.Write(buff)
}
