package tree

type Node struct {
	Data     string
	Parent   *Node
	Children map[string]*Node
}
