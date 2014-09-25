import (
    "gonerics.io/d/set/T.git"
)

type Graph map[T]set.Set

func New() Graph {
    return make(Graph)
}

func (g Graph) Connect(a, b T) {
    _, ok := g[a]

    if !ok {
        g[a] = set.New()
    }

    g[a].Add(b)
}

func (g Graph) Neighbours(a T) (ret []T) {
    if g[a] != nil {
        for key := range g[a] {
            ret = append(ret, key)
        }
    }

    return
}

func (g Graph) connected(a, b T, visited set.Set) bool {
    if visited.Contains(a) {
        return false
    } else {
        visited.Add(a)
    }

    for _, neighbour := range g.Neighbours(a) {
        if neighbour == b || g.connected(neighbour, b, visited) {
            return true
        }
    }
    return false
}

func (g Graph) Connected(a, b T) bool {
    visited := set.New()
    return g.connected(a, b, visited)
}
