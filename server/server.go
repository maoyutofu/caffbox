package main

import (
	. "github.com/tjz101/caffbox"
	c "github.com/tjz101/caffbox/controller"
	"github.com/tjz101/caffbox/util"
	"github.com/tjz101/caffmux"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		caffmux.Debug(err.Error())
		os.Exit(1)
	}
	caffmux.Debug(file)
	path, err := filepath.Abs(file)
	if err != nil {
		caffmux.Debug(err.Error())
		os.Exit(1)
	}
	path = filepath.Dir(path)
	caffmux.Debug(path)
	ParseConf(path)
	util.WritePid()
	caffmux.Debug("listen to " + Sett.Addr)

	app := caffmux.NewApplication()

	upload := &c.UploadController{}

	app.Router("^/u/?", upload)

	download := &c.DownloadController{}
	app.Router("^/d/?", download)

	view := &c.ViewController{}
	app.Router("^/v/:file(.*)", view)

	app.Run(Sett.Addr)
}
