package grapher

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

// Generate creates graph from passed data
func Generate(data []float64) {
	graph := asciigraph.Plot(data, asciigraph.Height(40), asciigraph.Width(100))
	fmt.Println(graph)
}
