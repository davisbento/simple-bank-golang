package api

import (
	db "github.com/davisbento/simple-bank-golang/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for the API.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// New server creates a new server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts/", server.listAllAccounts)

	server.router = router
	return server
}

// Start run the HTTP server on specified address
func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func errorResponseString(err string) gin.H {
	return gin.H{
		"error": err,
	}
}
