package login

import (
	"log"
	"testing"
)

func TestLogin(t *testing.T) {
	MyName = "AB3456"
	MyUniqueId = "9542a93a-82e7-403f-8857-4c55fd38391a"
	stat := Login()
	log.Printf("Status %v", stat)

	MyName = "srini"
	MyUniqueId = "9542a93a-82e7-403f-8857-4c55fd38391a"
	stat = Login()
	log.Printf("Status %v", stat)

}
