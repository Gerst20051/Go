package main

// [$]> go run script.go

// Complete the implementations for the following DAG operations.
//
// Example DAG:
//
//           1
//          / \
//         v  v
//         2->3->5
//        / \   ^
//       v  v   |
//       4  6->7
//          \
//          v
//          8

type Node struct {
  ID int
}

type Edge struct {
  SourceID int
  TargetID int
}

type Dag struct {
  Nodes []Node
  Edges []Edge
}

func NewDag(nodes []Node, edges []Edge) (*Dag, error) {
  return &Dag{Nodes: nodes, Edges: edges}, nil
}

func (d *Dag) AddNode(node Node, edges ...Edge) error {
  // TODO: implement
  return nil
}

func (d *Dag) FindNode(id int) (Node, error) {
  // TODO: implement
  return Node{}, nil
}
