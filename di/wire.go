//go:build wireinject
// +build wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/samriddhac/acme-flight/api/controller"
	"github.com/samriddhac/acme-flight/api/router"
	"github.com/samriddhac/acme-flight/config"
)

func InitializeRouter(cfg *config.Config, ctx *gin.Engine) router.Router {
	wire.Build(
		router.NewRouter, 
		controller.NewAcmeFlightController,
	)
	return router.Router{}
}
