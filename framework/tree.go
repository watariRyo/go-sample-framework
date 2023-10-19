package framework

import (
	"strings"
)

type TreeNodes struct {
	children []*TreeNodes
	handler  func(ctx *MyContext)
	param    string
	parent   *TreeNodes
}

func Constructor() *TreeNodes {
	return &TreeNodes{
		param:    "",
		children: []*TreeNodes{},
	}
}

func isGeneral(param string) bool {
	return strings.HasPrefix(param, ":")
}

func (this *TreeNodes) Insert(pathname string, handler func(ctx *MyContext)) {
	node := this

	params := strings.Split(pathname, "/")

	for _, param := range params {
		child := node.findChild(param)

		if child == nil {
			child = &TreeNodes{
				param:    param,
				children: []*TreeNodes{},
				parent:   node,
			}

			node.children = append(node.children, child)
		}

		node = child
	}

	node.handler = handler
}

func (this *TreeNodes) findChild(param string) *TreeNodes {
	for _, child := range this.children {
		if child.param == param {
			return child
		}
	}
	return nil
}

func (this *TreeNodes) Search(pathname string) *TreeNodes {
	params := strings.Split(pathname, "/")

	return dfs(this, params)
}

func dfs(node *TreeNodes, params []string) *TreeNodes {
	currentParam := params[0]
	isLastParam := len(params) == 1

	for _, child := range node.children {

		if isLastParam {
			if isGeneral(child.param) {
				return child
			}

			if child.param == currentParam {
				return child
			}

			continue
		}

		if !isGeneral(child.param) && child.param != currentParam {
			continue
		}

		result := dfs(child, params[1:])

		if result != nil {
			return result
		}
	}

	return nil
}

func (this *TreeNodes) ParaseParams(pathname string) map[string]string {
	node := this
	paramArr := strings.Split(pathname, "/")

	paramDicts := make(map[string]string)
	for i := len(paramArr) - 1; i > 0; i-- {
		if isGeneral(node.param) {
			paramDicts[node.param] = paramArr[i]
		}
		node = node.parent
	}

	return paramDicts
}
