package serve

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/RajaSrinivasan/edev/tools"
	"log"
	"net/http"
	"path/filepath"
)

func saveDeviceFile(c *gin.Context) {
	v := isValidClient(c)
	if !v {
		c.JSON(http.StatusForbidden, "Client Not Recognized")
		return
	}
	ofdir := filepath.Join()
	fj := c.Param("data")
	ofn, err := tools.FromJSON(fj, "/tmp")
	if err != nil {
		log.Printf("Error Receiving file %s from %s\n", ofn, c.Param("device"))
		c.JSON(http.StatusPreconditionFailed, "File Data corrupt")
		return
	}
	log.Printf("Received %s from %s", ofn, c.Param("device"))
	c.JSON(http.StatusOK, ofn)
}
