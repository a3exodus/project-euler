package helpers

import "math"

type Edge struct {
	Vertex string
	Weight int
}

func DijkstrasAlgorithm(G map[string][]Edge, source string) map[string]int {
	// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
	Q := map[string]bool{}
	dist := map[string]int{}
	prev := map[string]string{}
	for v, _ := range G {
		dist[v] = math.MaxInt32
		prev[v] = ""
		Q[v] = true
	}
	dist[source] = 0

	// while Q is not empty
	for len(Q) > 0 {
		// u <- vertex in Q with minimum dist[u]
		u := getMinVertexQ(Q, dist)
		delete(Q, u)

		// for each neighbor v of u still in Q
		edges := G[u]
		for _, edge := range edges {
			v := edge.Vertex
			if Q[v] {
				alt := dist[u] + edge.Weight
				if alt < dist[v] {
					dist[v] = alt
					prev[v] = u
				}
			}
		}
	}
	return dist
}

func getMinVertexQ(Q map[string]bool, dist map[string]int) string {
	minVertex := ""
	minDistance := math.MaxInt32
	for v, exists := range Q {
		if exists {
			distance := dist[v]
			if distance < minDistance {
				minVertex = v
				minDistance = distance
			}
		}
	}
	return minVertex
}
