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
var Name string

func genAuth() string {
	myuuid, _ := uuid.Parse(UniqueID)
	pw := slogin.Generate(Name, myuuid, time.Now().UTC())
	return Name + "/" + pw
}

func getAnswer(prefix string, postfix string, method string) string {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	auth := genAuth()
	fullurl := "https://" + Server + prefix + "/" + auth + "/" + postfix

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
	log.Printf("Response: %s", respstr)
	return respstr
}
