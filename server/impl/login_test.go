package impl

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

var fixedid uuid.UUID
var fixedtime time.Time

const adminPassword = "31c4197c7a90bfc75d19c0f30843df8e"
const timeLayout = "2006-01-02 15:04:05"
const fixedTimeStr = "2020-02-15 18:58:06"
const fixedAheadGoodStr = "2020-02-15 19:02:06"
const fixedAheadBadStr = "2020-02-15 19:05:06"

func init() {
	fixedid, _ = uuid.Parse("301245be-4e02-4036-bec4-ec20edbdaadd")
	fixedtime, _ = time.Parse(timeLayout, fixedTimeStr)
}
func TestGenerate(t *testing.T) {

	fixedid, _ := uuid.Parse("301245be-4e02-4036-bec4-ec20edbdaadd")
	fmt.Printf("Fixed ID: %s\n", fixedid.String())

	ft, _ := time.Parse(timeLayout, fixedTimeStr)
	fmt.Printf("Fixed date: %s\n", ft.Format(timeLayout))
	pw := Generate("admin", fixedid, ft)
	fmt.Printf("Password is: %s\n", pw)

	if strings.Compare(pw, adminPassword) == 0 {
		fmt.Printf("Password matches")
	} else {
		t.Error("Passwords do not match")
	}

}
func TestVerify(t *testing.T) {

	stat := Verify("admin", adminPassword, fixedid, fixedtime)
	fmt.Printf("Verification status %v\n", stat)

	goodAhead, _ := time.Parse(timeLayout, fixedAheadGoodStr)
	stat = Verify("admin", adminPassword, fixedid, goodAhead)
	fmt.Printf("Verification status Goodahead %v\n", stat)

	badAhead, _ := time.Parse(timeLayout, fixedAheadBadStr)
	stat = Verify("admin", adminPassword, fixedid, badAhead)
	fmt.Printf("Verification status Badahead %v\n", stat)
}
