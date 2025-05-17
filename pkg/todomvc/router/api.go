package router

import (
	"html/template"
	"io/fs"
	"net/http"
	"todomvc/pkg/todomvc/router/handler/todo"
	"todomvc/public"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(cors.Default())
	htmlFS, _ := fs.Sub(public.Public, "dist")
	jsFS, _ := fs.Sub(public.Public, "dist/js")
	node_modulesFS, _ := fs.Sub(public.Public, "dist/node_modules")
	router.StaticFS("/js", http.FS(jsFS))
	router.StaticFS("/node_modules", http.FS(node_modulesFS))
	templ := template.Must(template.New("").ParseFS(htmlFS, "index.html"))
	router.SetHTMLTemplate(templ)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	todoRoute := router.Group("/todo")
	todoRoute.GET("", todo.List)
	todoRoute.GET(":id", todo.Get)
	todoRoute.DELETE(":id", todo.Delete)
	todoRoute.POST("", todo.Create)
	todoRoute.PUT(":id", todo.Update)
	todoRoute.POST(":id/done", todo.Toggle)

	return router
}
