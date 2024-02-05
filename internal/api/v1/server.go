package api_v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/kimhieu153255/first-go/internal/config/db/sqlc"
	config_env "github.com/kimhieu153255/first-go/internal/config/env"
	"github.com/kimhieu153255/first-go/pkg/token"
)

type Server struct {
	Store      db.Store
	TokenMaker token.Maker
	Config     config_env.Config
	Router     *gin.Engine
}

func NewServer(store db.Store, config config_env.Config) (*Server, error) {
	maker, err := token.NewJWTMaker(config.SecretToken)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{Store: store, Config: config, TokenMaker: maker}
	router := gin.Default()

	v1Group := router.Group("/v1")

	// Grouping for hello
	v1Group.GET("/health", server.healthCheck)

	helloGroup := v1Group.Group("/hello")
	helloGroup.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World from Gin!",
		})
	})

	// Grouping for user
	userGroup := v1Group.Group("/users")
	userGroup.POST("", server.createUser)
	userGroup.GET("/:id", server.getUserByID)
	userGroup.GET("", server.getListUser)

	// Grouping for Auth
	authGroup := v1Group.Group("/auth")
	authGroup.POST("/login", server.login)
	authGroup.POST("/register", server.register)

	server.Router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
