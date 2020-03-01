package tools

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestLoad(t *testing.T) {
	err := Load("../config/registry.json")
	if err == nil {
		ShowAll()
	}
}

func TestSave(t *testing.T) {
	Load("../config/registry.json")
	err := Save("../config/registry.out.json")
	if err != nil {
		t.Error(err)
	}
}

func TestFind(t *testing.T) {
	Load("../config/registry.json")
	d := Find("AB00099")
	if d == nil {
		log.Println("Cannot find AB00099")
	}

	d = Find("AB0001")
	if d != nil {
		log.Println("Found AB0001")
		d.Show()
	}
}

func TestRegister(t *testing.T) {
	Load("../config/registry.json")
	d := Find("AB0001")
	d.UniqueID = "newid"
	err := (*d).Register()
	if err != nil {
		log.Printf("%s", err)
		t.Fail()
	}

	newdev := Device{Name: "RL00001", UniqueID: "33-44--55-66", Publisher: false, Approved: false}
	err = newdev.Register()
	if err == nil {
		log.Printf("Registered RL00001")
	}

	Save("../config/registry.out.json")
}

func TestSetJSON(t *testing.T) {
	conts, _ := ioutil.ReadFile("../../config/registry.out.json")
	SetJSON(string(conts))
	ShowAll()
}
