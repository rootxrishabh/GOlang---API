package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rootxrishabh/GOlang-API/pkg/controllers"
)


var callBookStore = func (r *gin.Default()){
	r.GET("/", controllers.GetAllBooks)
	r.POST("/", controllers.CreateBook)
	r.GET("/{bookid}", controllers.GetBookById)
	r.PUT("/{bookid}", controllers.UpdateBook)
	r.DELETE("/{bookid}", controllers.DeleteBookById)
}

