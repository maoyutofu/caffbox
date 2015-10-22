package main

import (
	. "github.com/tjz101/caffbox"
	c "github.com/tjz101/caffbox/controller"
	"github.com/tjz101/caffbox/util"
	"github.com/tjz101/caffmux"
)

func main() {
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
