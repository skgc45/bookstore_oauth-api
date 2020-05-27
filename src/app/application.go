package app

import (
	"github.com/gin-gonic/gin"
	"github.com/skgc45/bookstore_oauth-api/src/http"
	"github.com/skgc45/bookstore_oauth-api/src/repository/db"
	"github.com/skgc45/bookstore_oauth-api/src/repository/rest"
	"github.com/skgc45/bookstore_oauth-api/src/services/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// Entites
	atService := access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository())
	// UseCase
	atHandler := http.NewAccessTokenHandler(atService)
	// Controller
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	// web
	router.Run(":8080")
}
