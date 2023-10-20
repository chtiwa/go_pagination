package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/chtiwa/go_pagination/initializers"
	"github.com/chtiwa/go_pagination/models"
	"github.com/gin-gonic/gin"
)

type PaginationData struct {
	NextPage     int
	PreviousPage int
	CurrentPage  int
	TotalPages   int
	TwoAfter     int
	TwoBelow     int
}

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
	var people []models.Person
	totalPages := math.Ceil(float64(totalRows) / perPage)
	fmt.Println(totalPages)
	// fmt.Println(totalRows / perPage)
	offset := (page - 1) * 10
	initializers.DB.Limit(10).Offset(offset).Find(&people)

	// render the page
	c.HTML(http.StatusOK, "index.html", gin.H{
		"people": people,
		"pagination": PaginationData{
			NextPage:     page + 1,
			PreviousPage: page - 1,
			CurrentPage:  page,
			TotalPages:   int(totalPages),
			TwoAfter:     page + 2,
			TwoBelow:     page - 2,
		},
	})
}
