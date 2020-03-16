package impl

import (
	"log"

	device "gitlab.com/RajaSrinivasan/edev/tools"
)

func Show(arg string) {
	log.Printf("Show %s", arg)
	resp := getAnswer("/d/show", arg, "GET")
	device.SetJSON(resp)
	device.ShowAll()
}
