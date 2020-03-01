package impl

import (
	"log"

	device "gitlab.com/RajaSrinivasan/edev/tools"
)

func Show(arg string) {
	log.Printf("Show %s", arg)
	top := "a"
	auth := genAuth()
	fullurl := "https://" + Server + "/" + top + "/show" + auth + "/" + arg
	resp := getAnswer(fullurl, "GET")
	device.SetJSON(resp)
	device.ShowAll()
}
