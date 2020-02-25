package serve

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/RajaSrinivasan/edev/server/device"
	"gitlab.com/RajaSrinivasan/edev/server/login"
)

func validateNewUser(c *gin.Context) bool {
	uid, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		log.Printf("Invalid UUID %s (err)", uid, err)
		return false
	}

	pw := login.Generate(c.Param("deviceid"), uid, time.Now().UTC())
	if strings.Compare(pw, c.Param("password")) != 0 {
		log.Printf("Password Incorrect. Got %s", pw)
		return false
	}
	return true
}

func register(c *gin.Context, admin bool) {
	log.Printf("Device %s Password %s ", c.Param("deviceid"), c.Param("password"))

	v := validateNewUser(c)
	if !v {
		c.String(http.StatusForbidden, "Authentication Failure")
		return
	}

	dev := device.Device{
		Name:       c.Param("deviceid"),
		UniqueID:   c.Param("uuid"),
		Version:    c.Param("softrev"),
		Approved:   false,
		Publisher:  admin,
		Registered: time.Now().UTC()}
	err := dev.Register()
	if err != nil {
		log.Printf("%s", err)
		c.String(http.StatusBadRequest, "Failed to register device")
		return
	}

	err = device.Save()
	if err != nil {
		log.Printf("%s", err)
		c.String(http.StatusBadRequest, "Failed to register device")
		return
	}
	c.String(http.StatusOK, "Registratered device")
}

func registerDevice(c *gin.Context) {
	register(c, false)
}

func registerAdmin(c *gin.Context) {
	register(c, true)
}
