package login

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"errors"

	"strings"
	"time"

	"github.com/google/uuid"
)

var driftMinutesAhead, driftMinutesBehind time.Duration

var saltBasis1 = []uint32{0x89866356, 0x04011986, 0x09171989, 0x10071956,
	0x07151954, 0x05221963}

var saltBasis = []uint32{27644437, // Bell Prime
	1046527, // Carol Primes
	16769023,
	1073676287,
	939193, // Circular Primes
	939391,
	993319,
	999331,
	26227, // Cuban Primes
	27361,
	33391,
	35317,
	39916801, // Factorial Primes
	479001599,
	28657, // Fibanocci Primes
	514229}

var uniqueIdTable map[string]uuid.UUID
var errorId uuid.UUID

func init() {
	driftMinutesAhead, _ = time.ParseDuration("-3m")
	driftMinutesBehind, _ = time.ParseDuration("3m")

	uniqueIdTable = make(map[string]uuid.UUID)
	id, _ := uuid.Parse("301245be-4e02-4036-bec4-ec20edbdaadd")
	uniqueIdTable["srini"] = id
	id, _ = uuid.Parse("f0018d8a-d221-46af-820e-53c9a7f44c64")
	uniqueIdTable["admin"] = id
	errorId = uuid.New()
	id, _ = uuid.Parse("9542a93a-82e7-403f-8857-4c55fd38391a")
	uniqueIdTable["AB3456"] = id
}

func lookupUniqueId(un string) (uuid.UUID, error) {
	val, err := uniqueIdTable[un]
	if !err {
		return errorId, errors.New("Undefined user " + un)
	}
	return val, nil
}

// Generate (un string, ud uuid.UUID, t time.Time) string
// generates a password based on the username, UUID and the timestamp.
// The time is truncated to the hour.
func Generate(un string, ud uuid.UUID, t time.Time) string {

	buf := new(bytes.Buffer)

	layout := "2006-01-02 15"
	ts := t.Format(layout)

	tsenc := make([]byte, 2*len(ts))
	hex.Encode(tsenc, []byte(ts))
	binary.Write(buf, binary.LittleEndian, tsenc)

	unenc := make([]byte, 2*len(un))
	hex.Encode(unenc, []byte(un))
	binary.Write(buf, binary.LittleEndian, unenc)

	for _, basis := range saltBasis {
		binary.Write(buf, binary.LittleEndian, basis)
	}

	binary.Write(buf, binary.LittleEndian, ud)

	finalpwd := md5.New()
	finalpwd.Write(buf.Bytes())
	result := finalpwd.Sum(nil)

	return hex.EncodeToString(result)
}

// Verify(username string, password string, ud uuid.UUID, ts ...time.Time) bool
// Verifies that the password provided matches with the generated password.
// Assuming the clocks drift, an allowance is imposed and the time dependent password
// is checked at 3 points in time.
// The real server will use the current time in UTC, Specific time overrides support
// unit testing.
func Verify(username string, password string, ud uuid.UUID, ts ...time.Time) bool {

	t1 := time.Now().UTC()
	if len(ts) > 0 {
		for _, t := range ts {
			t1 = t
		}
	}

	p1 := Generate(username, ud, t1)
	if strings.Compare(p1, password) == 0 {
		return true
	}

	t2 := t1.Add(driftMinutesAhead)
	p2 := Generate(username, ud, t2)
	if strings.Compare(p2, password) == 0 {
		return true
	}

	t3 := t1.Add(driftMinutesBehind)
	p3 := Generate(username, ud, t3)
	if strings.Compare(p3, password) == 0 {
		return true
	}

	return false
}
