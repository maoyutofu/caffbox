package controller

import (
	. "github.com/tjz101/caffbox"
	"github.com/tjz101/caffmux"
)

type UploadController struct {
	caffmux.Controller
}

func (c *UploadController) Post() {
	r := c.Content.Request
	w := c.Content.ResponseWriter
	path := r.FormValue("path")
	watermark := r.FormValue("watermark")
	file, header, err := r.FormFile("file")
	if err != nil {
		writeMsg(w, Response{ErrCode: CODE_NOT_FOUND, ErrMsg: MSG_NOT_FOUND})
		return
	}
	defer file.Close()
	resp, err := writeFileToPath(path, file, watermark, header)
	if err != nil {
		writeMsg(w, resp)
		return
	}
	writeMsg(w, resp)
}
