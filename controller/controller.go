package controller

import (
	"encoding/json"
	"fmt"
	"github.com/tjz101/caffbox"
	"github.com/tjz101/caffbox/util"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func writeFile(file multipart.File, header *multipart.FileHeader) (caffbox.Response, error) {
	resp, err := writeFileToPath("", file, header)
	return resp, err
}

func writeFileToPath(path string, file multipart.File, header *multipart.FileHeader) (caffbox.Response, error) {
	if path == "" {
		path = util.PathFromTime()
	}
	physicalPath := util.GetPhysicalPath(path)

	err := util.Mkdir(physicalPath)
	if err != nil {
		return caffbox.Response{ErrCode: caffbox.CODE_DIR_CREATE_FAILED, ErrMsg: caffbox.MSG_DIR_CREATE_FAILED}, err
	}

	destFilename := header.Filename
	if caffbox.Sett.Rename {
		destFilename = util.RandomString()
		destFilename = destFilename + util.GetFileSuffix(header.Filename)
	}
	filePhysicalPath := filepath.Join(physicalPath, destFilename)
	fla, err := util.FileExists(filePhysicalPath)
	if fla && err == nil {
		newFilename, err := util.IncreaseFilename(physicalPath, destFilename)
		if newFilename == destFilename || err != nil {
			return caffbox.Response{ErrCode: caffbox.CODE_FILE_ALREADY_EXISTS, ErrMsg: caffbox.MSG_FILE_ALREADY_EXISTS}, err
		} else {
			destFilename = newFilename
		}
		filePhysicalPath = filepath.Join(physicalPath, newFilename)
	}
	destFile, err := os.OpenFile(filePhysicalPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return caffbox.Response{ErrCode: caffbox.CODE_FILE_OPEN_FAILED, ErrMsg: caffbox.MSG_FILE_OPEN_FAILED}, err
	}
	defer destFile.Close()
	_, err = io.Copy(destFile, file)
	if err != nil {
		return caffbox.Response{ErrCode: caffbox.CODE_FILE_SAVE_FAILED, ErrMsg: caffbox.MSG_FILE_SAVE_FAILED}, err
	}
	absPath, _ := util.GetAbsPath(physicalPath)
	return caffbox.Response{ErrCode: caffbox.CODE_SUCCESS, ErrMsg: caffbox.MSG_SUCCESS, Data: caffbox.CaffFile{Name: destFilename, Path: absPath}}, nil
}

func writeMsg(w http.ResponseWriter, response caffbox.Response) {
	buff, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(w, string(buff))
}
