package serve

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/RajaSrinivasan/edev/tools"
)

func saveDeviceFile(c *gin.Context) {
	v := isValidClient(c)
	if !v {
		c.JSON(http.StatusForbidden, "Client Not Recognized")
		return
	}

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
