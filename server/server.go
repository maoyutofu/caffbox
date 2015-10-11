package main

import (
	c "github.com/tjz101/caffbox/controller"
	"github.com/tjz101/caffmux"
)

var (
	docs   string
	rename bool
)

func init() {

}

func main() {
	caffmux.Debug("listen to :7001")

	app := caffmux.NewApplication()

	upload := &c.UploadController{}

	app.Router("^/u/?", upload)

	download := &c.DownloadController{}
	app.Router("^/d/?", download)

	view := &c.ViewController{}
	app.Router("^/v/:file(.*)", view)

	app.Run(":7001")
}
