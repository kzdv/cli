package login

import (
	"github.com/gin-gonic/gin"
	"github.com/kzdv/cli/internal/assets"
	"github.com/kzdv/cli/pkg/creds"
)

const PORT = 12297

type Server struct {
	Engine *gin.Engine
	Chan   chan creds.Creds
}

var server *Server

func NewServer() *Server {
	gin.SetMode(gin.ReleaseMode)
	eng := gin.New()
	eng.Use(gin.Recovery())

	assets.LoadAssets(eng)

	eng.GET("/authorize", HandleAuthorize)
	eng.GET("/callback", HandleCallback)

	a := make(chan creds.Creds)

	server = &Server{
		Engine: eng,
		Chan:   a,
	}

	return server
}
