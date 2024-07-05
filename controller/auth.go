package controller

//import "net/http"
//import "log"
import "github.com/gin-gonic/gin"
import "gin-web-app/model"
//import "strconv"

func Auth(c *gin.Context) {
	var a model.Login
	var access = c.Query("access")     // Notice: string
	var secret = c.Query("secret")     // Notice: string
	a.Access = access
	a.Secret = secret
	_, err := a.Auth()
	if err != nil {
		//log.Println("Error:", res.Desc)
		c.JSON(200, res(500, "unauthorized", nil))
	} else {
		c.JSON(200, res(200, "authorized", nil))
	}
}