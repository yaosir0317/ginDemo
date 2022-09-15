package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleQuery(context *gin.Context) {
	a := context.DefaultQuery("a", "a")
	context.JSON(http.StatusOK, gin.H{"a": a})
}

func HandleForm(context *gin.Context) {
	a := context.DefaultPostForm("a", "a")
	b := context.PostFormMap("b")
	context.JSON(http.StatusOK, gin.H{
		"a": a,
		"b": b,
	})
}

func HandleJson(context *gin.Context) {
	b, _ := context.GetRawData()
	var data map[string]interface{}
	_ = json.Unmarshal(b, &data)
	context.JSON(http.StatusOK, data)
}
