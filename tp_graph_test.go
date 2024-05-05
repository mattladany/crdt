package crdt

import "testing"

func TestTPGraphInitialization(t *testing.T) {
	tpgraph := NewTwoPhaseGraph[int]("tpgraph")
	if tpgraph.VertexCount() != 0 {
		t.Fatalf("tpgraph should initialize to a size of 0")
	}
	if tpgraph.EdgeCount() != 0 {
		t.Fatalf("tpgraph should initialize to a size of 0")
	}
}

func TestTPGraphAddVertex(t *testing.T) {
	tpgraph := NewTwoPhaseGraph[int]("tpgraph")
	tpgraph.AddVertex(1)
	if !tpgraph.LookupVertex(1) {
		t.Fatalf("tpgraph should have the vertex 1 within it")
	}
	if tpgraph.VertexCount() != 1 {
		t.Fatalf("tpgraph vertices should have a size of 1")
	}
	if tpgraph.EdgeCount() != 0 {
		t.Fatalf("tpgraph edges should have a size of 0")
	}
}

func TestTPGraphAddEdge(t *testing.T) {
	tpgraph := NewTwoPhaseGraph[int]("tpgraph")
	tpgraph.AddVertex(1)
	tpgraph.AddVertex(2)
	tpgraph.AddEdge(1, 2)
	if !tpgraph.LookupEdge(1, 2) {
		t.Fatalf("tpgraph should have had the edge within it")
	}
	if tpgraph.VertexCount() != 2 {
		t.Fatalf("tpgraph vertices should have a size of 2")
	}
	if tpgraph.EdgeCount() != 1 {
		t.Fatalf("tpgraph edges should have a size of 1")
	}
}

func TestTPGraphAddEdgeMissingv1(t *testing.T) {
	tpgraph := NewTwoPhaseGraph[int]("tpgraph")
	tpgraph.AddVertex(2)
	tpgraph.AddEdge(1, 2)
	if tpgraph.LookupEdge(1, 2) {
		t.Fatalf("tpgraph vertices should not have the the edge within it")
	}
	if tpgraph.VertexCount() != 1 {
		t.Fatalf("tpgraph should have a size of 1")
	}
	if tpgraph.EdgeCount() != 0 {
		t.Fatalf("tpgraph edges should have a size of 0")
	}
}

func TestTPGraphAddEdgeMissingv2(t *testing.T) {
	tpgraph := NewTwoPhaseGraph[int]("tpgraph")
	tpgraph.AddVertex(1)
	tpgraph.AddEdge(1, 2)
	if tpgraph.LookupEdge(1, 2) {
		t.Fatalf("tpgraph vertices should not have the the edge within it")
	}
	if tpgraph.VertexCount() != 1 {
		t.Fatalf("tpgraph should have a size of 1")
	}
	if tpgraph.EdgeCount() != 0 {
		t.Fatalf("tpgraph edges should have a size of 0")
	}
}

func TestTPGraphRemoveVertexNoEdge(t *testing.T) {
	tpgraph := NewTwoPhaseGraph[int]("tpgraph")
	tpgraph.AddVertex(1)
	tpgraph.AddVertex(2)
	tpgraph.AddVertex(3)
	tpgraph.AddEdge(1, 2)
	tpgraph.RemoveVertex(3)
	if tpgraph.LookupVertex(3) {
		t.Fatalf("tpgraph vertices should not have the vertex 3 within it")
	}
	if !tpgraph.LookupEdge(1, 2) {
		t.Fatalf("tpgraph vertices should have the edge within it")
	}
	if tpgraph.VertexCount() != 2 {
		t.Fatalf("tpgraph should have a size of 2")
	}
	if tpgraph.EdgeCount() != 1 {
		t.Fatalf("tpgraph edges should have a size of 1")
	}
}

func TestTPGraphRemoveVertexWithv1Edge(t *testing.T) {
	tpgraph := NewTwoPhaseGraph[int]("tpgraph")
	tpgraph.AddVertex(1)
	tpgraph.AddVertex(2)
	tpgraph.AddVertex(3)
	tpgraph.AddEdge(1, 2)
	tpgraph.RemoveVertex(1)
	if tpgraph.LookupVertex(1) {
		t.Fatalf("tpgraph vertices should not have the vertex 1 within it")
	}
	if tpgraph.LookupEdge(1, 2) {
		t.Fatalf("tpgraph vertices should not have the edge within it")
	}
	if tpgraph.VertexCount() != 2 {
		t.Fatalf("tpgraph should have a size of 2")
	}
	if tpgraph.EdgeCount() != 0 {
		t.Fatalf("tpgraph edges should have a size of 0")
	}
}

func TestTPGraphRemoveVertexWithv2Edge(t *testing.T) {
	tpgraph := NewTwoPhaseGraph[int]("tpgraph")
	tpgraph.AddVertex(1)
	tpgraph.AddVertex(2)
	tpgraph.AddVertex(3)
	tpgraph.AddEdge(1, 2)
	tpgraph.RemoveVertex(2)
	if tpgraph.LookupVertex(2) {
		t.Fatalf("tpgraph vertices should not have the vertex 2 within it")
	}
	if tpgraph.LookupEdge(1, 2) {
		t.Fatalf("tpgraph vertices should not have the edge within it")
	}
	if tpgraph.VertexCount() != 2 {
		t.Fatalf("tpgraph should have a size of 2")
	}
	if tpgraph.EdgeCount() != 0 {
		t.Fatalf("tpgraph edges should have a size of 0")
	}
}

func TestTPGraphRemoveEdge(t *testing.T) {
	tpgraph := NewTwoPhaseGraph[int]("tpgraph")
	tpgraph.AddVertex(1)
	tpgraph.AddVertex(2)
	tpgraph.AddVertex(3)
	tpgraph.AddEdge(1, 2)
	tpgraph.RemoveEdge(1, 2)
	if tpgraph.LookupEdge(1, 2) {
		t.Fatalf("tpgraph edges should not have the edge within it")
	}
	if tpgraph.VertexCount() != 3 {
		t.Fatalf("tpgraph should have a size of 3")
	}
	if tpgraph.EdgeCount() != 0 {
		t.Fatalf("tpgraph edges should have a size of 0")
	}
}

func TestTPGraphRemoveInvalidEdge(t *testing.T) {
	tpgraph := NewTwoPhaseGraph[int]("tpgraph")
	tpgraph.AddVertex(1)
	tpgraph.AddVertex(2)
	tpgraph.AddVertex(3)
	tpgraph.AddEdge(1, 2)
	tpgraph.RemoveEdge(2, 3)
	if !tpgraph.LookupEdge(1, 2) {
		t.Fatalf("tpgraph edges should have the edge within it")
	}
	if tpgraph.VertexCount() != 3 {
		t.Fatalf("tpgraph should have a size of 3")
	}
	if tpgraph.EdgeCount() != 1 {
		t.Fatalf("tpgraph edges should have a size of 1")
	}
}

func TestTPGraphRemoveInvertedEdge(t *testing.T) {
	tpgraph := NewTwoPhaseGraph[int]("tpgraph")
	tpgraph.AddVertex(1)
	tpgraph.AddVertex(2)
	tpgraph.AddVertex(3)
	tpgraph.AddEdge(1, 2)
	tpgraph.RemoveEdge(2, 1)
	if !tpgraph.LookupEdge(1, 2) {
		t.Fatalf("tpgraph edges should have the edge within it")
	}
	if tpgraph.VertexCount() != 3 {
		t.Fatalf("tpgraph should have a size of 3")
	}
	if tpgraph.EdgeCount() != 1 {
		t.Fatalf("tpgraph edges should have a size of 1")
	}
}

func TestTPGraphMerge(t *testing.T) {
	tpgraph := NewTwoPhaseGraph[int]("tpgraph")
	tpgraph.AddVertex(1)
	tpgraph.AddVertex(2)
	tpgraph.AddVertex(3)
	tpgraph.AddEdge(1, 2)

	tpgraph2 := NewTwoPhaseGraph[int]("tpgraph")
	tpgraph2.AddVertex(4)
	tpgraph2.AddVertex(5)
	tpgraph2.AddEdge(5, 4)

	tpgraph.Merge(tpgraph2)

	if !tpgraph.LookupVertex(1) {
		t.Fatalf("tpgraph vertices should not have the vertex 1 within it")
	}
	if !tpgraph.LookupVertex(2) {
		t.Fatalf("tpgraph vertices should not have the vertex 2 within it")
	}
	if !tpgraph.LookupVertex(3) {
		t.Fatalf("tpgraph vertices should not have the vertex 3 within it")
	}
	if !tpgraph.LookupVertex(4) {
		t.Fatalf("tpgraph vertices should not have the vertex 4 within it")
	}
	if !tpgraph.LookupVertex(5) {
		t.Fatalf("tpgraph vertices should not have the vertex 5 within it")
	}
	if !tpgraph.LookupEdge(5, 4) {
		t.Fatalf("tpgraph vertices should not have the edge within it")
	}
	if !tpgraph.LookupEdge(1, 2) {
		t.Fatalf("tpgraph vertices should not have the edge within it")
	}
	if tpgraph.VertexCount() != 5 {
		t.Fatalf("tpgraph should have a size of 5")
	}
	if tpgraph.EdgeCount() != 2 {
		t.Fatalf("tpgraph edges should have a size of 2")
	}
}

func TestTPGraphMergeIdempotent(t *testing.T) {
	tpgraph := NewTwoPhaseGraph[int]("tpgraph")
	tpgraph.AddVertex(1)
	tpgraph.AddVertex(2)
	tpgraph.AddVertex(3)
	tpgraph.AddEdge(1, 2)

	tpgraph2 := NewTwoPhaseGraph[int]("tpgraph")
	tpgraph2.AddVertex(4)
	tpgraph2.AddVertex(5)
	tpgraph2.AddEdge(5, 4)

	tpgraph.Merge(tpgraph2)

	if !tpgraph.LookupVertex(1) {
		t.Fatalf("tpgraph vertices should not have the vertex 1 within it")
	}
	if !tpgraph.LookupVertex(2) {
		t.Fatalf("tpgraph vertices should not have the vertex 2 within it")
	}
	if !tpgraph.LookupVertex(3) {
		t.Fatalf("tpgraph vertices should not have the vertex 3 within it")
	}
	if !tpgraph.LookupVertex(4) {
		t.Fatalf("tpgraph vertices should not have the vertex 4 within it")
	}
	if !tpgraph.LookupVertex(5) {
		t.Fatalf("tpgraph vertices should not have the vertex 5 within it")
	}
	if !tpgraph.LookupEdge(5, 4) {
		t.Fatalf("tpgraph vertices should not have the edge within it")
	}
	if !tpgraph.LookupEdge(1, 2) {
		t.Fatalf("tpgraph vertices should not have the edge within it")
	}
	if tpgraph.VertexCount() != 5 {
		t.Fatalf("tpgraph should have a size of 5")
	}
	if tpgraph.EdgeCount() != 2 {
		t.Fatalf("tpgraph edges should have a size of 2")
	}

	tpgraph.Merge(tpgraph2)
	tpgraph.Merge(tpgraph2)

	if !tpgraph.LookupVertex(1) {
		t.Fatalf("tpgraph vertices should have had the vertex 1 within it")
	}
	if !tpgraph.LookupVertex(2) {
		t.Fatalf("tpgraph vertices should have had the vertex 2 within it")
	}
	if !tpgraph.LookupVertex(3) {
		t.Fatalf("tpgraph vertices should have had the vertex 3 within it")
	}
	if !tpgraph.LookupVertex(4) {
		t.Fatalf("tpgraph vertices should have had the vertex 4 within it")
	}
	if !tpgraph.LookupVertex(5) {
		t.Fatalf("tpgraph vertices should have had the vertex 5 within it")
	}
	if !tpgraph.LookupEdge(5, 4) {
		t.Fatalf("tpgraph vertices should have had the edge 5 -> 4 within it")
	}
	if !tpgraph.LookupEdge(1, 2) {
		t.Fatalf("tpgraph vertices should have had the edge 1 -> 2 within it")
	}
	if tpgraph.VertexCount() != 5 {
		t.Fatalf("tpgraph should have had a size of 5")
	}
	if tpgraph.EdgeCount() != 2 {
		t.Fatalf("tpgraph edges should have had a size of 2")
	}
}

func TestTPGraphMergeNameMismatch(t *testing.T) {
	tpgraph := NewTwoPhaseGraph[int]("tpgraph")
	tpgraph.AddVertex(1)
	tpgraph.AddVertex(2)
	tpgraph.AddVertex(3)
	tpgraph.AddEdge(1, 2)

	tpgraph2 := NewTwoPhaseGraph[int]("tpgraph2")
	tpgraph2.AddVertex(4)
	tpgraph2.AddVertex(5)
	tpgraph2.AddEdge(5, 4)

	tpgraph.Merge(tpgraph2)

	if !tpgraph.LookupVertex(1) {
		t.Fatalf("tpgraph vertices should have the vertex 1 within it")
	}
	if !tpgraph.LookupVertex(2) {
		t.Fatalf("tpgraph vertices should have the vertex 2 within it")
	}
	if !tpgraph.LookupVertex(3) {
		t.Fatalf("tpgraph vertices should have the vertex 3 within it")
	}
	if tpgraph.LookupVertex(4) {
		t.Fatalf("tpgraph vertices should not have the vertex 4 within it")
	}
	if tpgraph.LookupVertex(5) {
		t.Fatalf("tpgraph vertices should not have the vertex 5 within it")
	}
	if tpgraph.LookupEdge(5, 4) {
		t.Fatalf("tpgraph vertices should not have the edge within it")
	}
	if !tpgraph.LookupEdge(1, 2) {
		t.Fatalf("tpgraph vertices should have the edge within it")
	}
	if tpgraph.VertexCount() != 3 {
		t.Fatalf("tpgraph should have a size of 3")
	}
	if tpgraph.EdgeCount() != 1 {
		t.Fatalf("tpgraph edges should have a size of 1")
	}
}
