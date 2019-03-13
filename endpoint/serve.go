package endpoint

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-contrib/cors"
	"github.com/iden3/tx-forwarder/config"
	"github.com/iden3/tx-forwarder/eth"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

var ethsrv *eth.EthService

func init() {
	gin.SetMode(gin.DebugMode)
}

func serveServiceApi() *http.Server {
	// start serviceapi
	api := gin.Default()
	api.Use(cors.Default())

	serviceapi := api.Group("/api/unstable")
	serviceapi.GET("/", handleGetInfo)
	serviceapi.POST("/tx", handlePostTx)

	serviceapisrv := &http.Server{Addr: config.C.Server.ServiceApi, Handler: api}
	go func() {
		log.Info("API server at ", config.C.Server.ServiceApi)
		if err := serviceapisrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("listen: %s\n", err)
		}
	}()
	return serviceapisrv
}

func serveAdminApi(stopch chan interface{}) *http.Server {
	api := gin.Default()
	api.Use(cors.Default())
	adminapi := api.Group("/api/unstable")

	adminapi.POST("/stop", func(c *gin.Context) {
		// yeah, use curl -X POST http://<adminserver>/stop
		c.String(http.StatusOK, "got it, shutdowning server")
		stopch <- nil
	})

	adminapi.GET("/info", handleInfo)

	adminapisrv := &http.Server{Addr: config.C.Server.AdminApi, Handler: api}
	go func() {
		log.Info("API server at ", config.C.Server.AdminApi)
		if err := adminapisrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("listen: %s\n", err)
		}
	}()
	return adminapisrv
}

func Serve(ethservice *eth.EthService) {
	ethsrv = ethservice

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
	adminapisrv := serveAdminApi(stopch)

	// wait until shutdown signal
	<-stopch
	log.Info("Shutdown Server ...")

	if err := serviceapisrv.Shutdown(context.Background()); err != nil {
		log.Error("ServiceApi Shutdown:", err)
	} else {
		log.Info("ServiceApi stopped")
	}
	if err := adminapisrv.Shutdown(context.Background()); err != nil {
		log.Error("AdminApi Shutdown:", err)
	} else {
		log.Info("AdminApi stopped")
	}
}
