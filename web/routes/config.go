package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/loganetherton/pm-go/config"
	"strings"
)

func setupRouter() *gin.Engine {
	setGinMode()
	r := gin.Default()
	setTrustedProxies(r)
	return r
}

func setGinMode() {
	if config.IsDev {
		gin.SetMode(gin.DebugMode)
	} else if config.IsTest {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func setTrustedProxies(r *gin.Engine) {
	var trusted string
	trusted = config.TrustedProxies
	if trusted == "" {
		trusted = "127.0.0.1"
	}
	proxies := strings.Split(trusted, ",")
	if err := r.SetTrustedProxies(proxies); err != nil {
		panic(err)
	}
}
