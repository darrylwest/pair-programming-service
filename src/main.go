//
// pair programming service
//
// @author darryl.west <darwest@ebay.com>
// @created 2018-04-08 09:26:59

package main

import (
	"./edit"
)

func main() {
	edit.CreateLogger()
	cfg := edit.ParseArgs()
	if cfg == nil {
		edit.ShowHelp()
		return
	}

	service, err := edit.NewService(cfg)
	if err != nil {
		panic(err)
	}

	err = service.Start()
	if err != nil {
		println(err.Error())
	}
}
