package impl

import "log"

var UniqueID string

func Register(pub bool) error {
	var postfix string
	if pub {
		postfix = UniqueID + "/pub"
	} else {
		postfix = UniqueID
	}

	regresp := getAnswer("/d/reg", postfix, "POST")
	log.Printf("Registration response: %s", regresp)
	return nil
}
