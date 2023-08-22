package tree

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

type PathArray struct {
	Paths []string
}

type AccessTree struct {
	Root *Node
}

func (tree *AccessTree) EndpointAccess(level string) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Request.Cookie("username")
		authorized := tree.Find(cookie.Value)
		fmt.Println(authorized)
		c.Set("example", "12345")
		c.Next()
		log.Print("auth fires")
	}
}

func (tree *AccessTree) ReadUserFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}
	var obj PathArray
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}
	for i := 0; i < len(obj.Paths); i++ {
		tree.AddUri(obj.Paths[i])
	}
}

func (tree *AccessTree) add(segments []string) {
	runner := tree.Root
	for len(segments) > 0 {
		if runner.Children[segments[0]] == nil {
			var newNode = tree.AddNode(segments[0], runner)
			runner = &newNode
		} else {
			runner = runner.Children[segments[0]]
		}
		segments = segments[1:]
	}
}

func (tree *AccessTree) AddNode(val string, parent *Node) Node {
	var node = Node{Data: val, Children: make(map[string]*Node)}
	parent.Children[val] = &node
	node.Parent = parent
	return node
}

func (tree *AccessTree) AddUri(path string) {
	splitPath := strings.Split(path, "/")
	out := splitPath[1:]
	tree.add(out)
}

func (tree *AccessTree) Find(name string) string {
	return name
}

func (tree *AccessTree) ListUsers() []string {
	return []string{"ok"}
}

func (tree AccessTree) RemoveUser() {

}

func MakeAccessTree(root_name string) (*AccessTree, error) {
	root := Node{Data: root_name, Parent: nil, Children: make(map[string]*Node)}
	return &AccessTree{Root: &root}, nil
}
