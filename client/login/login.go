package login

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	slogin "gitlab.com/RajaSrinivasan/edev/tools"
)

var MyUniqueId string
var MyName string

func Login() bool {

	myuuid, _ := uuid.Parse(MyUniqueId)

	pw := slogin.Generate(MyName, myuuid, time.Now().UTC())

	// TODO: This is insecure; use only in dev environments.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	url := "https://localhost:8689/a/login/" + MyName + "/" + pw
	log.Printf("Dialing %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("%v", req)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer resp.Body.Close()
	respbytes, _ := ioutil.ReadAll(resp.Body)
	log.Printf("Response: %s", string(respbytes))
	return true
}
