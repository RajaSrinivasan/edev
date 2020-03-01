package tools

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type FilePayLoad struct {
	Name      string
	Size      int64
	Mode      uint32
	Contents  string
	Signature string
}

func ToJSON(fn string) (string, error) {

	s, err := os.Stat(fn)
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	fb, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	fsig := md5.New()
	fsig.Write(fb)

	payload := FilePayLoad{
		Name:      fn,
		Size:      s.Size(),
		Mode:      uint32(s.Mode()),
		Contents:  base64.StdEncoding.EncodeToString(fb),
		Signature: hex.EncodeToString(fsig.Sum(nil)),
	}

	pj, err := json.Marshal(payload)
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	return string(pj), nil
}

func FromJSON(conts string) (string, string, error) {
	var payload FilePayLoad

	err := json.Unmarshal([]byte(conts), &payload)
	if err != nil {
		log.Printf("%s", err)
		return "", "", err
	}

	fb, err := base64.StdEncoding.DecodeString(payload.Contents)
	if err != nil {
		log.Printf("%s", err)
		return "", "", err
	}

	fsig := md5.New()
	fsig.Write(fb)
	fsigstr := hex.EncodeToString(fsig.Sum(nil))

	if strings.Compare(fsigstr, payload.Signature) != 0 {
		log.Printf("Signature Comparison failed. Got %s Expecting %s\n", fsigstr, payload.Signature)
		return "", "", errors.New("Signature Comparison failed")
	} else {
		log.Printf("File Signature %s\n", fsigstr)
	}

	of, err := ioutil.TempFile("", payload.Name)
	if err != nil {
		log.Printf("%s", err)
		return "", "", err
	}
	defer of.Close()

	ofn := of.Name()
	of.Write(fb)
	return payload.Name, ofn, nil
}
