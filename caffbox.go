package caffbox

import (
	"fmt"
	"github.com/tjz101/goprop"
	"os"
	"os/exec"
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

func ParseConf() {
	prop := goprop.NewProp()
	path, err := GetExecPath()
	if err != nil {
		os.Exit(1)
	}
	confFilename := fmt.Sprintf("%s/conf/conf.properties", path)
	prop.Read(confFilename)
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
	Name         string `json:"name"`
	Path         string `json:"path"`
	OriginalName string `json:"originalName"`
}

func GetExecPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	return filepath.Dir(path), nil
}
func WritePid() error {
	path, err := GetExecPath()
	if err != nil {
		return err
	}
	pidFilename := fmt.Sprintf("%s/logs/caffbox.pid", path)
	pid := os.Getpid()
	f, err := os.OpenFile(pidFilename, os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(strconv.Itoa(pid))
	return err
}
