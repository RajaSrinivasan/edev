package impl

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	slogin "gitlab.com/RajaSrinivasan/edev/tools"
)

var Server string
var UniqueID string
var Name string

// TODO: This is insecure; use only in dev environments.

func genAuth() string {
	myuuid, _ := uuid.Parse(UniqueID)
	pw := slogin.Generate(Name, myuuid, time.Now().UTC())
	return "/" + Name + "/" + pw
}

func getAnswer(fullurl string, method string) string {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	log.Printf("Dialing %s", fullurl)
	req, err := http.NewRequest(method, fullurl, nil)
	if err != nil {
		log.Fatalf("%v", req)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer resp.Body.Close()
	respbytes, _ := ioutil.ReadAll(resp.Body)
	respstr := string(respbytes)
	log.Printf("Register Response: %s", respstr)
	return respstr
}

func Register(pub bool) error {
	log.Printf("Register %s Publisher %v", Name, pub)
	top := "d"
	if pub {
		top = "a"
	}
	auth := genAuth()
	fullurl := "https://" + Server + "/" + top + "/reg" + auth + "/" + UniqueID + "/A"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	log.Printf("Dialing %s", fullurl)
	req, err := http.NewRequest("POST", fullurl, nil)
	if err != nil {
		log.Fatalf("%v", req)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer resp.Body.Close()
	respbytes, _ := ioutil.ReadAll(resp.Body)
	log.Printf("Register Response: %s", string(respbytes))
	return nil
}
