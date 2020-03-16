package impl

import (
	"log"
)

func makeDevSpec(devs []string) string {
	//devspec, _ := json.Marshal(devs)
	return devs[0]
}

func Authorize(revoke bool, list bool, devices []string) error {

	prefix := "/d/auth"
	if list {
		authResp := getAnswer(prefix, "", "GET")
		log.Printf("Auth (list) Response : %s", authResp)
		return nil
	}

	devspec := makeDevSpec(devices)

	if revoke {
		authResp := getAnswer(prefix, "rev/"+devspec, "PUT")
		log.Printf("Auth (revoke) Response: %s", authResp)
		return nil
	}

	authResp := getAnswer(prefix, "app/"+devspec, "PUT")
	log.Printf("Auth (approve) Response: %s", authResp)

	return nil
}
