package controller

import (
	. "github.com/tjz101/caffbox"
	"github.com/tjz101/caffbox/util"
	"github.com/tjz101/caffmux"
	"io/ioutil"
	"os"
	"path/filepath"
)

type DownloadController struct {
	caffmux.Controller
}

func (c *DownloadController) Get() {
	r := c.Content.Request
	w := c.Content.ResponseWriter

	file := r.FormValue("file")
	if file == "" {
		writeMsg(w, Response{ErrCode: CODE_NOT_FOUND, ErrMsg: MSG_NOT_FOUND})
		return
	}
	filePath := util.GetPhysicalPath(file)
	fla, _ := util.FileExists(filePath)
	if !fla {
		writeMsg(w, Response{ErrCode: CODE_NOT_FOUND, ErrMsg: MSG_NOT_FOUND})
		return
	}
	display := r.FormValue("display")
	if display == "" {
		display = filepath.Base(file)
	}

	f, err := os.Open(filePath)
	if err != nil {
		writeMsg(w, Response{ErrCode: CODE_FILE_OPEN_FAILED, ErrMsg: MSG_FILE_OPEN_FAILED})
		return
	}
	defer f.Close()
	buff, err := ioutil.ReadAll(f)
	if err != nil {
		writeMsg(w, Response{ErrCode: CODE_FILE_READ_FAILED, ErrMsg: MSG_FILE_READ_FAILED})
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+display+"\"")
	w.Write(buff)
}
