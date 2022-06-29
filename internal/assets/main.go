package assets

import (
	"embed"
	"html/template"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var f embed.FS

func LoadAssets(router *gin.Engine) {
	templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl"))
	router.SetHTMLTemplate(templ)
}
