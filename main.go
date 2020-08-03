package main

import (
	"fmt"
	
	"github.com/gin-gonic/gin"
) 

func handleOCRConversion(c *gin.Context){
	//To be implemented
}

func main() {
	r := gin.Default()
	
	r.POST("/ocr-conversion", handleOCRConversion)
	//TODO no need to handle static pages as they would be served by React
	//We'll be only building an api back-end for the moment.
	//r.GET("/", handleHomepage???)
	
	r.Run()

}

