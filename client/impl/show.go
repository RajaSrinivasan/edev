package impl

import (
	"log"

	"gitlab.com/RajaSrinivasan/edev/server/device"
)

func Show(arg string) {
	log.Printf("Show %s", arg)
	top := "a"
	auth := genAuth()
	fullurl := "https://" + Server + "/" + top + "/show" + auth + "/" + arg
	resp := getAnswer(fullurl, "GET")
	device.Set(resp)
	device.ShowAll()
}
