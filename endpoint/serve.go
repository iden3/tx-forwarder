package endpoint

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-contrib/cors"
	"github.com/iden3/gas-station/config"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func serveServiceApi() *http.Server {
	// start serviceapi
	api := gin.Default()
	api.Use(cors.Default())

	serviceapi := api.Group("/api/unstable")
	serviceapi.GET("/", handleGetInfo)

	serviceapisrv := &http.Server{Addr: config.C.Server.ServiceApi, Handler: api}
	go func() {
		log.Info("API server at ", config.C.Server.ServiceApi)
		if err := serviceapisrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("listen: %s\n", err)
		}
	}()
	return serviceapisrv
}

func Serve() {

	stopch := make(chan interface{})

	// catch ^C to send the stop signal
	ossig := make(chan os.Signal, 1)
	signal.Notify(ossig, os.Interrupt)
	go func() {
		for sig := range ossig {
			if sig == os.Interrupt {
				stopch <- nil
			}
		}
	}()

	// start servers
	serviceapisrv := serveServiceApi()

	// wait until shutdown signal
	<-stopch
	log.Info("Shutdown Server ...")

	if err := serviceapisrv.Shutdown(context.Background()); err != nil {
		log.Error("ServiceApi Shutdown:", err)
	} else {
		log.Info("ServiceApi stopped")
	}
}
