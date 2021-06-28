package zdf

type node struct {
	pattern   string
	part      string
	chirldren []*node
	isWild    bool
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.chirldren {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	var nodes []*node
	for _, child := range n.chirldren {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.chirldren = append(n.chirldren, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || (n.part != "" && n.part[0] == '*') {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	childrens := n.matchChildren(part)
	for _, child := range childrens {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
