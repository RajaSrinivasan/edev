package serve

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/RajaSrinivasan/edev/tools"
)

func revokeDevice(c *gin.Context) {
	log.Printf("Revoke Device %s ", c.Param("devspec"))
	v := isValidClient(c)
	if !v {
		c.String(http.StatusForbidden, "Authentication Failure")
		return
	}
	err := tools.Approve(c.Param("devspec"), false)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, "Revoked approval device")

}

func approveDevice(c *gin.Context) {
	log.Printf("Approve Device %s ", c.Param("devspec"))
	v := isValidClient(c)
	if !v {
		c.String(http.StatusForbidden, "Authentication Failure")
		return
	}
	err := tools.Approve(c.Param("devspec"), true)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, "Approved device")
}
