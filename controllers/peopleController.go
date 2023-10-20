package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/chtiwa/go_pagination/helpers"
	"github.com/chtiwa/go_pagination/initializers"
	"github.com/chtiwa/go_pagination/models"
	"github.com/gin-gonic/gin"
)

func PeopleIndexGET(c *gin.Context) {
	// get page number
	page := 1
	pageString := c.Param("page")
	if pageString != "" {
		page, _ = strconv.Atoi(pageString)
	}
	perPage := 10.0

	// total rows
	var totalRows int64
	initializers.DB.Model(&models.Person{}).Count(&totalRows)

	// get data
	totalPages := math.Ceil(float64(totalRows) / perPage)

	offset := (page - 1) * 10
	var people []models.Person
	initializers.DB.Limit(10).Offset(offset).Find(&people)

	pagination := helpers.GetPaginationData(page, int(totalPages), "/people")

	// render the page
	c.HTML(http.StatusOK, "index.html", gin.H{
		"people":     people,
		"pagination": pagination,
	})
}
