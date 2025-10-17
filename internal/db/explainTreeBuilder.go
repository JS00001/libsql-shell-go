package db

import (
	"fmt"
)

type QueryPlanNode struct {
	ID       string
	ParentID string
	NotUsed  string
	Detail   string
	Children []*QueryPlanNode
}

func BuildQueryPlanTree(rows [][]string) (*QueryPlanNode, error) {
	nodeMap := make(map[string]*QueryPlanNode)
	var topLevel []*QueryPlanNode

	for _, row := range rows {
		id := row[0]
		parentID := row[1]
		notUsed := row[2]
		detail := row[3]

		node := &QueryPlanNode{
			ID:       id,
			ParentID: parentID,
			NotUsed:  notUsed,
			Detail:   detail,
		}

		nodeMap[id] = node
		if parentID == "0" {
			topLevel = append(topLevel, node)
		}
	}

	for _, node := range nodeMap {
		if node.ParentID != "0" {
			parent, ok := nodeMap[node.ParentID]
			if ok {
				parent.Children = append(parent.Children, node)
			}
		}
	}

	root := &QueryPlanNode{
		Detail:   "QUERY PLAN",
		Children: topLevel,
	}

	return root, nil
}

func PrintQueryPlan(root *QueryPlanNode) {
	printTree(root, []bool{})
}

func printTree(node *QueryPlanNode, ancestorsLast []bool) {
	pfx := ""

	for range ancestorsLast {
		pfx += "|   "
	}

	if len(ancestorsLast) > 0 {
		pfx += "|- "
	}

	fmt.Println(pfx + node.Detail)

	for i, child := range node.Children {
		isLast := i == len(node.Children)-1
		printTree(child, append(ancestorsLast, isLast))
	}
}
