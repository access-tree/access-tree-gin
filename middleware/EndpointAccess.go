package middleware

import (
	"strings"

	"github.com/access-tree/access-tree-gin/tree"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func EndpointAccess(tree *tree.AccessTree, level string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var path []string
		user, err := c.Request.Cookie("username")
		if err != nil {
			c.AbortWithStatus(403)
		}
		path = strings.Split(c.Request.URL.Path, "/")
		path = path[1:]
		path = append(path, c.Request.Method)
		path = append([]string{user.Value}, path...)
		paramsArray := maps.Values(c.Params)
		newPath := []string{}
		if len(paramsArray) > 0 {
			for _, item := range path {
				itemInParams := contains(paramsArray, item)
				if !itemInParams {
					newPath = append(newPath, item)
				}
			}
		} else {
			newPath = path
		}
		accessPermission := tree.Find(newPath)
		read := (accessPermission > 3)
		readWrite := (accessPermission > 4)
		if level == "read" {
			if read || readWrite {
				c.Next()
			} else {
				c.AbortWithStatus(403)
			}
		} else if level == "write" {
			if readWrite {
				c.Next()
			} else {
				c.AbortWithStatus(403)
			}
		}
	}
}
