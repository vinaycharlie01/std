package handler

import (
	"myapp/studentservice/db"
	"myapp/studentservice/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentHandler struct {
	Queries *db.Queries
}

func NewStudentHandler(in *gorm.DB) *StudentHandler {
	return &StudentHandler{db.NewDB(in)}
}

func (s *StudentHandler) CreateStudent(c *gin.Context) {
	var a model.Student
	if err := c.BindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad reql"})
		return
	}
	err := s.Queries.CreateStudeNt(&a)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, a)
}

func (s *StudentHandler) GetStudent(c *gin.Context) {
	id := c.Param("id")

	idx, _ := strconv.Atoi(id)

	res, err := s.Queries.Read(idx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *StudentHandler) GetAllStudent(c *gin.Context) {

	res, err := s.Queries.ReadAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *StudentHandler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	idx, _ := strconv.Atoi(id)

	err := s.Queries.Remove(idx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status:": "deleted"})
}
