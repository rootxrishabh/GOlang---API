package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rootxrishabh/GOlang---API/pkg/controllers"
)


var CallBookStore = func (r *gin.Engine){
	r.GET("/book", controllers.GetBooks)
	r.POST("/", controllers.CreateBook)
	r.GET("/:bookid", controllers.GetBookById)
	r.PUT("/:bookid", controllers.UpdateBook)
	r.DELETE("/:bookid", controllers.DeleteBook)
}