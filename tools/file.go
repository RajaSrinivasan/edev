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
	"path"
	"path/filepath"
	"strings"
)

// This structure represents the payload.
type FilePayLoad struct {
	Name      string                    // base name of the file
	Size      int64                     // size of the file
	Mode      uint32                    // Mode
	Contents  string                    // file contents converted to base64 notation
	Signature string                    // hash signature of the contents using md5
}

// ToJSON(fn string) (string, error) packs the contents of the specified file
// into a JSON structure suitable for transmission to a server (http)
// The file mode, size are extracted; an md5 signature is computed of the contentts
// to create a FilePayload structure which is convenrted into a JSON string.
// Returns the json string and a status code
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
		Name:      path.Base(fn),
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


// FromJSON(conts string, odir string) (string, error) is the reverse of the ToJSON. It takes
// the json string conts and extracts the file contents and creates a file in the odir directory.
// Before it accepts the data, the md5 signature is verified.
// Returns the output file name (odir + the base name of the file) and a status code
func FromJSON(conts string, odir string) (string, error) {
	var payload FilePayLoad

	err := json.Unmarshal([]byte(conts), &payload)
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	fb, err := base64.StdEncoding.DecodeString(payload.Contents)
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	fsig := md5.New()
	fsig.Write(fb)
	fsigstr := hex.EncodeToString(fsig.Sum(nil))

	if strings.Compare(fsigstr, payload.Signature) != 0 {
		log.Printf("Signature Comparison failed. Got %s Expecting %s\n", fsigstr, payload.Signature)
		return "", errors.New("Signature Comparison failed")
	} else {
		log.Printf("File Signature %s\n", fsigstr)
	}

	ofn := filepath.Join(odir, payload.Name)
	of, err := os.Create(ofn)
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	defer of.Close()

	of.Write(fb)
	return ofn, nil
}
