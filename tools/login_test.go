package tools

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
)

var fixedid, altid uuid.UUID
var fixedtime, goodAhead, badAhead time.Time

const adminPassword = "31c4197c7a90bfc75d19c0f30843df8e"
const timeLayout = "2006-01-02 15:04:05"
const fixedTimeStr = "2020-02-15 18:58:06"
const fixedAheadGoodStr = "2020-02-15 19:02:06"
const fixedAheadBadStr = "2020-02-15 19:05:06"

const fixedidstr = "301245be-4e02-4036-bec4-ec20edbdaadd"
const altidstr = "f0018d8a-d221-46af-820e-53c9a7f44c64"

func init() {
	fixedid, _ = uuid.Parse(fixedidstr)
	fixedtime, _ = time.Parse(timeLayout, fixedTimeStr)

	goodAhead, _ = time.Parse(timeLayout, fixedAheadGoodStr)
	badAhead, _ = time.Parse(timeLayout, fixedAheadBadStr)

	altid := uuid.New()
	fmt.Printf("Alternate ID: %s", altid)
}

func TestGenerate(t *testing.T) {
	t.Logf("Fixed Id for admin-----------------------------------")
	pw := Generate("admin", fixedid, fixedtime)
	t.Logf("Good Password admin: %s\n", pw)
	pw = Generate("admin", fixedid, goodAhead)
	t.Logf("Good Ahead Password admin: %s\n", pw)
	pw = Generate("admin", fixedid, badAhead)
	t.Logf("Bad Ahead Password admin: %s\n", pw)

	t.Logf("Alternate Id for admin-----------------------------------")
	pw = Generate("admin", altid, fixedtime)
	t.Logf("Good Password admin: %s\n", pw)
	pw = Generate("admin", altid, goodAhead)
	t.Logf("Good Ahead Password admin: %s\n", pw)
	pw = Generate("admin", altid, badAhead)
	t.Logf("Bad Ahead Password admin: %s\n", pw)

	t.Logf("Fixed Id for user-----------------------------------")
	pw = Generate("user", fixedid, fixedtime)
	t.Logf("Good Password user: %s\n", pw)
	pw = Generate("user", fixedid, goodAhead)
	t.Logf("Good Ahead Password user: %s\n", pw)
	pw = Generate("user", fixedid, badAhead)
	t.Logf("Bad Ahead Password user: %s\n", pw)

}
func TestVerify(t *testing.T) {

	stat := Verify("admin", adminPassword, fixedid, fixedtime)
	t.Logf("Verification status %v\n", stat)

	stat = Verify("admin", adminPassword, fixedid, goodAhead)
	t.Logf("Verification status Goodahead %v\n", stat)

	stat = Verify("admin", adminPassword, fixedid, badAhead)
	t.Logf("Verification status Badahead %v\n", stat)
}

func TestUUID(t *testing.T) {
	for i := 0; i < 9; i++ {
		nextid := uuid.New()
		t.Logf("%s", nextid.String())
	}
}
