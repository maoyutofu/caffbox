package controller

import (
	//"bytes"
	"errors"
	"fmt"
	"github.com/tjz101/caffbox/util"
	"github.com/tjz101/caffmux"
	//"io/ioutil"
	"net/http"
	//"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

type ViewController struct {
	caffmux.Controller
}

func getResize(r *http.Request) (string, error) {
	index := strings.Index(r.RequestURI, "?")
	if index >= 0 {
		param := string(r.RequestURI[index+1:])
		params := strings.Split(param, "/")
		lengths := len(params)
		if lengths%2 == 0 {
			kv := make(map[string]string)
			for i := 0; i < lengths; i += 2 {
				kv[params[i]] = params[i+1]
			}
			if kv["resize"] != "" {
				resize := kv["resize"]
				reg, err := regexp.Compile(`\d{1,}x\d{1,}`)
				if err != nil {
					return "", err
				}
				if reg.MatchString(resize) {
					return resize, nil
				} else {
					return "", errors.New("Invalid parameter")
				}
			} else {
				return "", errors.New("Invalid parameter")
			}
		} else {
			return "", errors.New("The number of parameters is not correct")
		}
	}
	return "", errors.New("Do not specify the resize parameter")
}

func thumbnail(filePath string, resize string) (string, error) {
	preFilePath := filePath
	filePath = filepath.Join(filepath.Dir(preFilePath), resize, filepath.Base(preFilePath))
	if preFilePath != filePath {
		m := &sync.Mutex{}
		m.Lock()
		defer m.Unlock()
		fla, _ := util.FileExists(filePath)
		if !fla {
			newFilePath := filepath.Dir(filePath)
			fla, _ := util.DirExists(newFilePath)
			if !fla {
				err := util.Mkdir(newFilePath)
				if err != nil {
					return preFilePath, err
				}
			}
			cmdFormat := "convert %s -resize %s %s"
			cmd := fmt.Sprintf(cmdFormat, preFilePath, resize, filePath)
			list := strings.Split(cmd, " ")
			c := exec.Command(list[0], list[1:]...)
			err := c.Run()
			if err != nil {
				return preFilePath, err
			}
			return filePath, nil
		} else {
			return filePath, nil
		}
	}
	return preFilePath, nil
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

	suffix := filePath[strings.LastIndex(filePath, "."):]
	if util.IsPic(filePath) {
		resize, err := getResize(r)
		if err != nil && resize == "" {
			caffmux.Debug(err)
		} else {
			filePath, err = thumbnail(filePath, resize)
			if err != nil {
				caffmux.Debug(err)
			}
			caffmux.Debug(filePath)
		}
	}

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
	case ".json":
		contentType = "application/json"
	default:
	}

	if contentType == "" {
		http.NotFound(w, r)
		return
	}
	/*
		f, err := os.Open(filePath)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer f.Close()
		d, err := f.Stat()
		if err != nil {
			http.NotFound(w, r)
			return
		}
		buff, err := ioutil.ReadAll(f)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", contentType)
		w.Write(buff)
		http.ServeContent(w, r, d.Name(), d.ModTime(), bytes.NewReader(buff))
	*/
	http.ServeFile(w, r, filePath)
}
