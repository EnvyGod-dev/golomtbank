package api

import (
	"fmt"
	db "task/db/sqlc"
	"task/token"
	"task/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %v", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupServer()

	return server, nil
}
func (server *Server) setupServer() {

	router := gin.Default()
	router.POST("/user/create", server.createUser)
	router.POST("/user/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleWare(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccounts)
	authRoutes.GET("/account/:id", server.getAccounts)
	authRoutes.GET("/accounts", server.listAccounts)

	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}
func (server *Server) StartServer(addres string) error {
	return server.router.Run(addres)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
