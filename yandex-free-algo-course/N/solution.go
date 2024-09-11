package main

// Comment it before submitting
/*
type Node struct {
	val        int
	neighbours []*Node
}

func (n *Node) setVal(val int) {
	n.val = val
}

func (n *Node) getVal() int {
	return n.val
}

func (n *Node) setNeighbours(neighbours []*Node) {
	n.neighbours = neighbours
}

func (n *Node) getNeighbours() []*Node {
	return n.neighbours
}

func (n *Node) addNeighbour(neighbour *Node) {
	n.neighbours = append(n.neighbours, neighbour)
}

func newNode(val int) *Node {
	n := &Node{}
	n.setNeighbours(make([]*Node, 0))
	n.setVal(val)
	return n
}
*/
// Comment until this line

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return cloneNode(node, visited)
}

func cloneNode(node *Node, visited map[*Node]*Node) *Node {

	// Anti Loop
	if nd, ok := visited[node]; ok {
		return nd
	}

	if len(node.neighbours) == 0 {
		res := newNode(node.getVal())
		visited[node] = res
		return res
	}

	res := newNode(node.getVal())
	visited[node] = res
	for _, neighbour := range node.getNeighbours() {
		res.addNeighbour(cloneNode(neighbour, visited))
	}

	return res
}

/*
func printGraph(node *Node, lvl int, visited map[*Node]bool) {
	prefix := ""
	for range lvl {
		prefix += " "
	}

	if visited[node] {
		fmt.Printf("%sLoop!!! Node: %d\n", prefix, node.val)
		return
	}
	visited[node] = true

	fmt.Printf("%sNode: %d\n", prefix, node.val)
	for _, neighbour := range node.neighbours {
		printGraph(neighbour, lvl+1, visited)
	}
}

func main() {

	nb111 := newNode(111)
	nb11 := newNode(11)
	nb11.setNeighbours([]*Node{nb111})
	nb12 := newNode(12)
	head := newNode(1)
	head.setNeighbours([]*Node{nb11, nb12})
	nb111.addNeighbour(head)

	newGraph := cloneGraph(head)

	printGraph(newGraph, 0, map[*Node]bool{})
}
*/
