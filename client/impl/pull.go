package impl

import (
	"fmt"
	"log"
)

func Pull(list bool, reffile string, devspec []string, outdir string) error {

	var postfix string
	if list {
		postfix = fmt.Sprintf("list:/", reffile)
	} else {
		postfix = fmt.Sprintf("f/%s", devspec)
	}
	pullresp := getAnswer("/d/pull", postfix, "GET")
	log.Printf("Pull postfix %s response: %s", postfix, pullresp)
	return nil
}
