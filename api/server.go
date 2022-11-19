package api

import (
	db "github.com/Sciipia/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server HTTP requests fot our banking service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer created a new HTTP Server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts/:id", server.GetAccount)
	router.GET("/accounts", server.ListAccount)
	router.PUT("/accounts", server.UpdateAccount)
	router.DELETE("/accounts/:id", server.DeleteAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
