package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

type NewProfileDataTransferObject struct {
	Name string `json:"name" binding:"required"`
}

func printRequest(req *http.Request) {
	if httpDump, err := httputil.DumpRequest(req, true); err == nil {
		fmt.Println("Printing HTTP Request")
		fmt.Printf("%q", httpDump)
		fmt.Println()
		contents, _ := ioutil.ReadAll(req.Body)
		fmt.Println(string(contents))
		fmt.Println()
		req.Body = ioutil.NopCloser(bytes.NewReader(contents))
	} else {
		fmt.Println("Could not read http content correctly", err)
	}
}

func main() {
	r := gin.Default()

	var prodileDb = newProfileDatabase()

	r.GET("/profile", func(c *gin.Context) {
		printRequest(c.Request)
		c.JSON(200, prodileDb.GetAll())
	})

	r.POST("/profile", func(c *gin.Context) {
		printRequest(c.Request)
		var newProfile NewProfileDataTransferObject
		if err := c.BindJSON(&newProfile); err == nil {
			fmt.Println("profile", newProfile)
			var p = ProfileDatabaseModel{Name: newProfile.Name}
			c.JSON(200, prodileDb.Create(p))
		} else {
			fmt.Println(err)
			panic("failed to parse new profile request")
		}
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
