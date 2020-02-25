package serve

import (
	"bufio"
	"crypto/tls"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/RajaSrinivasan/edev/server/login"
)

var tempid uuid.UUID

func init() {
	tempid, _ = uuid.Parse("301245be-4e02-4036-bec4-ec20edbdaadd")
}

func serveClient(client net.Conn) {
	defer client.Close()
	log.Println("Serving Client")
	rdr := bufio.NewReader(client)
	for {
		msg, err := rdr.ReadString('\n')
		if err != nil {
			log.Printf("ClientSock Error: %s", err)
			return
		}
		log.Printf("Received %s", msg)
		client.Write([]byte(msg))
	}
}
func ProvideServiceRaw(certfn, pvtkeyfn, hostnport string) {
	log.Printf("Providing Service")
	servercert, err := tls.LoadX509KeyPair(certfn, pvtkeyfn)
	if err != nil {
		log.Fatal(err)
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{servercert},
	}

	listener, err := tls.Listen("tcp", hostnport, config)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		client, err := listener.Accept()
		if err != nil {
			log.Printf("%s", err)
			continue
		}
		go serveClient(client)
	}
}

func getTop(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"title": "Home Page",
		},
	)
	//c.String(http.StatusOK, "Hello from TOPR LLC.")
}

func showDevice(c *gin.Context) {
	c.String(http.StatusOK, "Show Device")
}

func showAllDevices(c *gin.Context) {
	c.String(http.StatusOK, "Show All Devices")
}

func loginAdminAPI(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")
	log.Printf("Login: %s %s", username, password)

	status := login.Verify(username, password, tempid)
	if status {
		c.String(http.StatusOK, "Login "+username+" succeeded ")
	} else {
		c.String(http.StatusForbidden, "Login failed "+username+" "+password)
	}
}

func ProvideService(certfn, pvtkeyfn, hostnport string, htmlpath string) {
	log.Printf("Providing Service HTTPS")

	r := gin.Default()
	r.LoadHTMLGlob(htmlpath + "/*")
	r.GET("/", getTop)
	// Device functions
	devroutes := r.Group("/d")
	devroutes.POST("/reg/:deviceid/:password/:uuid/:softrev", registerDevice)
	devroutes.GET("/show", showAllDevices)
	devroutes.GET("/show/:deviceid", showDevice)

	// Admin functions
	adminroutes := r.Group("/a")
	adminroutes.GET("/login/:username/:password", loginAdminAPI)
	adminroutes.POST("/reg/:deviceid/:password/:uuid/:softrev", registerAdmin)

	r.RunTLS(hostnport, certfn, pvtkeyfn)

}
