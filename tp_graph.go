package crdt

// TwoPhaseGraph is a CRDT graph in which edges and vertexes can only be added
// or removed once throughout the lifetime of the graph.
type TwoPhaseGraph[T comparable] struct {
	name     string
	vertices *TwoPhaseSet[T]
	edges    *TwoPhaseSet[edge[T]]
}

// edge is an internal representation of a directional edge between two vertices
// v1 -> v2.
type edge[T comparable] struct {
	v1 T
	v2 T
}

// NewTwoPhaseGraph constructs a TwoPhaseGraph with the provided name.
// It is assumed the name of this specific TwoPhaseGraph uniquely identifies this
// register throughout the cluster.
func NewTwoPhaseGraph[T comparable](name string) *TwoPhaseGraph[T] {
	tpgraph := new(TwoPhaseGraph[T])
	tpgraph.name = name
	tpgraph.vertices = NewTwoPhaseSet[T](name + "-vertices")
	tpgraph.edges = NewTwoPhaseSet[edge[T]](name + "-edges")
	return tpgraph
}

// AddVertex will add the provided vertex to the graph if it has not been added before.
func (graph *TwoPhaseGraph[T]) AddVertex(vertex T) {
	graph.vertices.Add(vertex)
}

// AddEdge will add the edge v1 -> v2 iff:
// 	- v1 and v2 exist in the graph
//  - v1 -> v2 has never been added to the graph before
func (graph *TwoPhaseGraph[T]) AddEdge(v1 T, v2 T) {
	if graph.vertices.Lookup(v1) && graph.vertices.Lookup(v2) {
		e := edge[T]{
			v1: v1,
			v2: v2,
		}
		graph.edges.Add(e)
	}
}

// RemoveVertex removes the provided vertex from the graph if it currently exists
// within it.
func (graph *TwoPhaseGraph[T]) RemoveVertex(vertex T) {
	if graph.vertices.Lookup(vertex) {
		graph.vertices.Remove(vertex)
		graph.edges.RemoveIf(func(e edge[T]) bool {
			return e.v1 == vertex || e.v2 == vertex
		})
	}
}

// RemoveEdge removes the provided edge v1 -> v2 from the graph if it currently
// exists within it.
func (graph *TwoPhaseGraph[T]) RemoveEdge(v1 T, v2 T) {
	graph.edges.Remove(edge[T]{
		v1: v1,
		v2: v2,
	})
}

// LookupVertex reports whether the provided vertex currently exists within the
// graph or not.
func (graph *TwoPhaseGraph[T]) LookupVertex(vertex T) bool {
	return graph.vertices.Lookup(vertex)
}

// LookupEdge reports whether the provided edge v1 -> v2 currently exists within the
// graph or not.
func (graph *TwoPhaseGraph[T]) LookupEdge(v1 T, v2 T) bool {
	return graph.edges.Lookup(edge[T]{
		v1: v1,
		v2: v2,
	})
}

// VertexCount returns the current number of vertices that exist within the graph.
func (graph *TwoPhaseGraph[T]) VertexCount() int {
	return graph.vertices.Size()
}

// EdgeCount returns the current number of edges that exist within the graph.
func (graph *TwoPhaseGraph[T]) EdgeCount() int {
	return graph.edges.Size()
}

// Merge will merge the contents of the TwoPhaseSets vertices and edges from that
// to this graph.
// This is an idempotent operation and is a no-op if graph.name != graph.name.
func (graph *TwoPhaseGraph[T]) Merge(that *TwoPhaseGraph[T]) {
	if graph.name == that.name {
		graph.vertices.Merge(that.vertices)
		graph.edges.Merge(that.edges)
	}
}
