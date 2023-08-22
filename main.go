package main

import (
	"fmt"
	"log"
	"time"

	"github.com/access-tree/access-tree-gin/tree"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "12345")
		c.Next()
		latency := time.Since(t)
		log.Print(latency)
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()
	tree, err := tree.MakeAccessTree("root")
	if err != nil {
		fmt.Println(err)
	}
	tree.ReadUserFile("./userData.json")
	fmt.Printf("%v", tree)
	r.Use(Logger())

	r.GET("/test", tree.EndpointAccess("read"), func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
