package api

import (
	"net/http"
	"report-lkl-morning/repository"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repo repository.MongoRepository
}

func (h Handler) GetAllReport(c *gin.Context) {
	InfoGet, err := h.Repo.GetAll()
	c.JSON(http.StatusOK, gin.H{"Report": InfoGet})
	if err != nil {
		panic(err)
	}
}

type InfoReport struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Detail    string `json:"detail"`
}

func (h Handler) Postinfo(c *gin.Context) {
	var info InfoReport
	err := c.ShouldBindJSON(&info)
	c.JSON(http.StatusOK, gin.H{"เพิ่มรายละเอียดของ": info.FirstName})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err2 := h.Repo.InsertInfo(info.ID, info.FirstName, info.LastName, info.Detail)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h Handler) Delinfo(c *gin.Context) {
	query := c.Query("id")
	info, err := h.Repo.FindInfo(query)
	c.JSON(http.StatusOK, gin.H{"ข้อมูลที่ลบ": info})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	count, err2 := h.Repo.Deleteinfo(query)
	c.JSON(http.StatusOK, gin.H{"จำนวนในการลบ ": count})
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
	}
}

func (h Handler) GetReport(c *gin.Context) {
	info, err := h.Repo.FindInfo(c.Query("id"))
	c.JSON(http.StatusOK, gin.H{"ข้อมูลที่ค้นหา": info})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
