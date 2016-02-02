package main

import (
	"fmt"
	. "github.com/tjz101/caffbox"
	c "github.com/tjz101/caffbox/controller"
	"github.com/tjz101/caffmux"
	"os"
)

func command() {
	arg_num := len(os.Args)
	if arg_num > 1 {
		for i := 1; i < arg_num; i++ {
			cmd := os.Args[i]
			if cmd == "-v" {
				fmt.Println("caffbox version: caffbox/1.0.4")
			}
		}
		os.Exit(0)
	}
}

func main() {
	command()
	ParseConf()
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
