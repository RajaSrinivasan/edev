package serve

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/RajaSrinivasan/edev/server/device"
	"gitlab.com/RajaSrinivasan/edev/server/login"
)

func isValidClient(c *gin.Context) bool {
	d := device.Find(c.Param("deviceid"))
	if d == nil {
		return false
	}
	uid, _ := uuid.Parse(d.UniqueID)
	return login.Verify(c.Param("deviceid"), c.Param("password"), uid, time.Now().UTC())
}
func showDevices(c *gin.Context) {
	v := isValidClient(c)
	if !v {
		c.JSON(http.StatusForbidden, "Client Not Recognized")
		return
	}

	d := device.Find(c.Param("devspec"))
	if d != nil {
		c.JSON(http.StatusOK, d)
	} else {
		c.JSON(http.StatusNotFound, nil)
	}

}

func showAllDevices(c *gin.Context) {
	v := isValidClient(c)
	if !v {
		c.JSON(http.StatusForbidden, "Client Not Recognized")
		return
	}
	c.JSON(http.StatusOK, device.KnownDevices)
}
