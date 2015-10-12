package caffbox

import (
	"github.com/tjz101/goprop"
	"path/filepath"
	"strconv"
)

const (
	CODE_SUCCESS             int = 0
	CODE_NOT_FOUND           int = -1
	CODE_PERMISSION_DENIED   int = -2
	CODE_FILE_OPEN_FAILED    int = -3
	CODE_FILE_SAVE_FAILED    int = -4
	CODE_FILE_ALREADY_EXISTS int = -5
	CODE_DIR_CREATE_FAILED   int = -6
	CODE_FILE_READ_FAILED    int = -7

	MSG_SUCCESS             string = "success"
	MSG_NOT_FOUND           string = "no such file"
	MSG_PERMISSION_DENIED   string = "permission denied"
	MSG_FILE_OPEN_FAILED    string = "file open failed"
	MSG_FILE_SAVE_FAILED    string = "file save failed"
	MSG_FILE_ALREADY_EXISTS string = "file already exists"
	MSG_DIR_CREATE_FAILED   string = "directory to create failure"
	MSG_FILE_READ_FAILED    string = "file read failure"

	ROOT_DIR string = "/"
)

var (
	Sett             *Setting
	RootPhysicalPath string
)

func init() {
	prop := goprop.NewProp()
	prop.Read("./conf.properties")
	addr := prop.Get("addr")
	docs := prop.Get("docs")
	rename, err := strconv.ParseBool(prop.Get("rename"))
	if err != nil {
		rename = false
	}
	Sett = &Setting{Addr: addr, Rename: rename}
	RootPhysicalPath = filepath.Join(docs, ROOT_DIR)
}

type Setting struct {
	Addr   string
	Rename bool
}

type Response struct {
	ErrCode int         `json:"errcode"`
	ErrMsg  string      `json:"errmsg"`
	Data    interface{} `json:"data,omitempty"`
}

type CaffFile struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Filename string `json:"filename"`
}
