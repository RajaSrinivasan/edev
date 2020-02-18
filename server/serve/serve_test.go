package serve

import (
	"log"
	"testing"
)

func TestProvideServiceRaw(t *testing.T) {

	log.SetPrefix(t.Name())
	ProvideServiceRaw("../install/certfile", "/Users/rajasrinivasan/.ssh/id_rsa", "localhost:9999")
}

func TestProvideService(t *testing.T) {

	log.SetPrefix(t.Name())
	ProvideService("../../config/certfile", "/Users/rajasrinivasan/.ssh/id_rsa", "localhost:9443", "../../config/html")
}
