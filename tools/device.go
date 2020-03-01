package device

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"
	"time"
)

type Device struct {
	Name       string    `json:"name"`
	UniqueID   string    `json:"uuid"`
	Version    string    `json:"version"`
	Approved   bool      `json:"approved"`
	Publisher  bool      `json:"publisher"`
	Registered time.Time `json:"registered"`
}

var databaseName string
var KnownDevices []Device

func ShowAll() {
	for _, dev := range KnownDevices {
		log.Printf("Name : %s UUID %s Registered %v", dev.Name, dev.UniqueID, dev.Approved)
	}
}
func (dev *Device) Show() {
	log.Printf("Name : %s UUID %s Registered %v", dev.Name, dev.UniqueID, dev.Approved)
}

func SetJSON(devs string) error {
	json.Unmarshal([]byte(devs), &KnownDevices)
	return nil
}
func Load(fn string) error {
	ifile, err := os.Open(fn)
	if err != nil {
		log.Println(err)
		return err
	}
	defer ifile.Close()

	jload := json.NewDecoder(ifile)
	err = jload.Decode(&KnownDevices)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Loaded %s", fn)
	databaseName = fn
	return nil
}

func Save(fni ...string) error {
	fn := databaseName
	if len(fni) > 0 {
		fn = fni[0]
	}

	ofile, err := os.Create(fn)
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	defer ofile.Close()

	jcode := json.NewEncoder(ofile)
	jcode.SetIndent("    ", "  ")
	err = jcode.Encode(KnownDevices)
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	log.Printf("Saved device database to %s", fn)
	return nil
}

func Find(dn string) *Device {

	for _, dev := range KnownDevices {
		if strings.Compare(dn, dev.Name) == 0 {
			return &dev
		}
	}
	return nil
}

func (dev Device) Register() error {
	old := Find(dev.Name)
	if old == nil {
		KnownDevices = append(KnownDevices, dev)
		return nil
	}
	return errors.New("Already Registered")
}
