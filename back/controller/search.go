package controller

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"gopkg.in/gin-gonic/gin.v1"
)

type Searcher struct {
	Language string `json:"language"`
	Code     string `json:"code"`
	Output   string `json:"output"`
}

type Page struct {
	Title string
	URL   string
}

func (s *Searcher) Search(c *gin.Context) {
	var request Searcher
	c.BindJSON(&request)
	request.Language = `"""` + request.Language + `"""`
	request.Code = `"""` + request.Code + `"""`
	request.Output = `"""` + request.Output + `"""`
	output, err := exec.Command("python", "search.py", request.Language, request.Code, request.Output).Output()
	if err != nil {
		fmt.Println("ERROR!!!: ", err)
	}
	var response Page
	response.Title, response.URL = strings.Split(string(output), "\n")[0], strings.Split(string(output), "\n")[1]
	c.JSON(http.StatusOK, gin.H{
		"title": response.Title,
		"url":   response.URL,
	})
}
