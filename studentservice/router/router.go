package router

import (
	"myapp/studentservice/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var RegisterRouter = func(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")

	h := handler.NewStudentHandler(db)

	{
		api.POST("/student", h.CreateStudent)
		api.GET("/students/:id", h.GetStudent)
		api.GET("/students/", h.GetAllStudent)
		api.DELETE("/studentsde/:id", h.DeleteStudent)
	}
}
