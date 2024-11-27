package main

// [$]> go run tests.go

// https://github.com/smartystreets/goconvey
// https://www.smarty.com/blog/go-testing-part-1-vanillla
// https://www.smarty.com/blog/go-testing-part-2-running-tests
// https://www.smarty.com/blog/go-testing-part-3-convey-behavior
// [$]> go test -v -run=TestDAGAddNode
// go: cannot find main module, but found .git/config in /Users/andrewgerst/code/Github/Gerst20051/Go
//   to create a module there, run:
//   cd ../.. && go mod init

import (
  "testing"
)

func TestDAGAddNode(t *testing.T) {
  t.Run("AddNode adds a node to the DAG", func(t *testing.T) {
    dag, err := NewDag(
      []Node{
        {ID: 1},
        {ID: 2},
        {ID: 3},
        {ID: 4},
        {ID: 5},
        {ID: 6},
        {ID: 7},
        {ID: 8},
      },
      []Edge{
        {SourceID: 1, TargetID: 2},
        {SourceID: 1, TargetID: 3},
        {SourceID: 2, TargetID: 3},
        {SourceID: 2, TargetID: 4},
        {SourceID: 2, TargetID: 6},
        {SourceID: 3, TargetID: 5},
        {SourceID: 6, TargetID: 7},
        {SourceID: 7, TargetID: 5},
        {SourceID: 6, TargetID: 8},
      },
    )
    if err != nil {
      t.Errorf("expected a valid dag, got: %+v", err)
    }
    dag.AddNode(Node{ID: 9}, Edge{SourceID: 8, TargetID: 9})
    node, err := dag.FindNode(9)
    if err != nil {
      t.Errorf("expected no error, got %+v", err)
    }
    if node.ID != 9 {
      t.Errorf("expected node with ID 9, got %+v", node)
    }
  })
}
