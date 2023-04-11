package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @route: http://localhost:3000/
// @Method: GET
// @description: Homepage description about api

func HomePage(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the go-api challenge by D3V, Please read the README.md file for instruction and use postman tool to test this api",
	})
}
